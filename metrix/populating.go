package metrix

import (
	"errors"
	"fmt"
	"time"

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

	fmt.Printf("%s --- STARTING TO QUERY USERS ---\n", prefix)
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
					fmt.Printf("%s %d/%d Failed to get data of user: %s\n", prefix, numberOfResponses, queryLength, resp.Login)
					fmt.Printf("%s %d/%d Trying to get user data from cache for: %s\n", prefix, numberOfResponses, queryLength, resp.Login)
					dbData, err := controller.GetUserFromCache(resp.Login)
					if err != nil {
						fmt.Printf("%s %d/%d Failed to get data of user: %s\n", prefix, numberOfResponses, queryLength, resp.Login)
					} else {
						querySuccessMessage(resp.Login, startTime, numberOfResponses, queryLength)
						queriedUsers = append(queriedUsers, dbData)
					}
				} else {
					querySuccessMessage(resp.Login, startTime, numberOfResponses, queryLength)
					queriedUsers = append(queriedUsers, resp.User)
				}

			case <-time.After(50 * time.Millisecond):
				break

			}

			if numberOfResponse == blockSize {
				fmt.Printf("%s --- COOLDOWN: %ds ---\n", prefix, cooldown)
				fmt.Printf("%s --- TIME: %s ---\n", prefix, util.FormatDuration(time.Since(startTime)))
				break
			}

		}

		time.Sleep(time.Duration(cooldown) * time.Second)

	}

	fmt.Printf("%s --- FINISHED QUERYING USERS ---\n", prefix)

	return queriedUsers

}

func querySuccessMessage(login string, startTime time.Time, n int, overall int) {
	fmt.Printf("%s %d/%d User: %s, Time: %s\n", prefix, n, overall, login, util.FormatDuration(time.Since(startTime)))
}
