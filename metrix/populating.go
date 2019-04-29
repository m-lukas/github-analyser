package metrix

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/graphql"
	"github.com/m-lukas/github-analyser/util"
)

const (
	cooldown  int = 12
	blockSize int = 10
)

type UserResponse struct {
	Login string
	User  *db.User
	Error error
}

func populateData(filepaths []string) ([]*db.User, error) {

	/*
		inputArray, err := readInput(filepaths)
		if err != nil {
			return nil, err
		}
	*/

	inputArray, err := graphql.GetPopulatingData("sindresorhus")
	if err != nil {
		return nil, err
	}

	userArray := queryUserData(inputArray)
	if len(userArray) == 0 && len(inputArray) != 0 {
		return nil, errors.New("Empty user data array!")
	}

	return userArray, nil
}

func readInput(filepathes []string) ([]string, error) {

	var inputArray []string
	var successfulFileCount = 0

	for index, filepath := range filepathes {

		fmt.Printf("%s Started scanning of file: %s\n", prefix, filepath)
		fmt.Printf("%s Scanning file: %d/%d\n", prefix, index+1, len(filepathes))

		if !hasFileFormat(filepath, "txt") {
			fmt.Printf("%s Wrong file format: %s\n", prefix, filepath)
			continue
		}

		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("%s Could not open file: %s\n", prefix, filepath)
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if line != "" && !util.Contains(inputArray, line) {
				inputArray = append(inputArray, line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("%s Error while scanning file: %s\n", prefix, filepath)
			return nil, err
		}

		successfulFileCount++
		fmt.Printf("%s Finished file: %s\n", prefix, filepath)
	}

	fmt.Printf("%s Successful: %d/%d\n", prefix, successfulFileCount, len(filepathes))

	return inputArray, nil

}

func queryUserData(inputArray []string) []*db.User {

	fmt.Printf("%s --- STARTING TO QUERY USERS ---\n", prefix)
	var queriedUsers []*db.User

	var numberOfResponses = 0
	var queryLength = len(inputArray)
	var startTime = time.Now()

	for {

		if numberOfResponses == 100 {
			return queriedUsers
		}

		var block []string

		inputArray, block = popN(inputArray, blockSize)
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
					/*
						fmt.Printf("%s %d/%d Trying to get user data from cache for: %s\n", prefix, numberOfResponses, queryLength, resp.Login)
						dbData, err := getUserFromCache(resp.Login)
						if err != nil {
							fmt.Printf("%s %d/%d Failed to get data of user: %s\n", prefix, numberOfResponses, queryLength, resp.Login)
						} else {
							querySuccessMessage(resp.Login, startTime, numberOfResponses, queryLength)
							queriedUsers = append(queriedUsers, dbData)
						}
					*/
				} else {
					querySuccessMessage(resp.Login, startTime, numberOfResponses, queryLength)
					queriedUsers = append(queriedUsers, resp.User)
				}

			case <-time.After(50 * time.Millisecond):
				break

			}

			if numberOfResponse == blockSize {
				fmt.Printf("%s --- COOLDOWN: %ds ---\n", prefix, cooldown)
				fmt.Printf("%s --- TIME: %s ---\n", prefix, formatDuration(time.Since(startTime)))
				break
			}

		}

		time.Sleep(time.Duration(cooldown) * time.Second)

	}

	fmt.Printf("%s --- FINISHED QUERYING USERS ---\n", prefix)

	return queriedUsers

}

func querySuccessMessage(login string, startTime time.Time, n int, overall int) {
	fmt.Printf("%s %d/%d User: %s, Time: %s\n", prefix, n, overall, login, formatDuration(time.Since(startTime)))
}

func getUserFromCache(login string) (*db.User, error) {

	mongo, err := db.GetMongo()
	if err != nil {
		return nil, err
	}
	collection := mongo.Collection("users")

	dbUser, err := db.FindUser(login, collection)
	if err != nil {
		return nil, err
	}

	return dbUser, nil

}
