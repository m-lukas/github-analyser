package metrix

import (
	"errors"
	"fmt"
	"time"

	"github.com/m-lukas/github-analyser/logger"
	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"
)

const (
	cooldown  int = 15
	blockSize int = 10
)

type UserResponse struct {
	Login string
	User  *db.User
	Error error
}

func populateData(filepaths []string) ([]*db.User, error) {

	inputArray, err := util.ReadInputFiles(filepaths)
	if err != nil {
		return nil, err
	}

	userArray := queryUserData(inputArray)
	if len(userArray) == 0 && len(inputArray) != 0 {
		return nil, errors.New("Empty user data array!")
	}

	return userArray, nil
}

func queryUserData(inputArray []string) []*db.User {

	logger.Info(fmt.Sprintf("%s --- STARTING TO QUERY USERS ---", prefix))
	var queriedUsers []*db.User

	var numberOfResponses = 0
	var queryLength = len(inputArray)
	var startTime = time.Now()

	for {

		var block []string

		inputArray, block = util.PopN(inputArray, blockSize)
		if len(block) == 0 {
			break
		}

		channel := make(chan UserResponse)

		for _, login := range block {

			go func(userLogin string) {

				user, err := controller.GetUser(userLogin)
				if err != nil {
					channel <- UserResponse{Login: userLogin, User: nil, Error: err}
				} else {
					channel <- UserResponse{Login: userLogin, User: user, Error: nil}
				}

			}(login)

		}

		numberOfResponse := 0

		for {
			select {
			case resp := <-channel:

				numberOfResponse++
				numberOfResponses++
				if resp.Error != nil {
					logger.ErrorNoMail(fmt.Sprintf("%s %d/%d Failed to get data of user: %s", prefix, numberOfResponses, queryLength, resp.Login))
					logger.ErrorNoMail(resp.Error.Error())
				} else {
					querySuccessMessage(resp.Login, startTime, numberOfResponses, queryLength)
					queriedUsers = append(queriedUsers, resp.User)
				}

			case <-time.After(50 * time.Millisecond):
				break

			}

			if numberOfResponse == blockSize {
				logger.Info(fmt.Sprintf("%s --- COOLDOWN: %ds ---", prefix, cooldown))
				logger.Info(fmt.Sprintf("%s --- TIME: %s ---", prefix, util.FormatDuration(time.Since(startTime))))
				break
			}

		}

		time.Sleep(time.Duration(cooldown) * time.Second)

	}

	logger.Info(fmt.Sprintf("%s --- FINISHED QUERYING USERS ---", prefix))

	return queriedUsers

}

func querySuccessMessage(login string, startTime time.Time, n int, overall int) {
	logger.Info(fmt.Sprintf("%s %d/%d User: %s, Time: %s", prefix, n, overall, login, util.FormatDuration(time.Since(startTime))))
}
