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
	channel := make(chan UserResponse)
	startTime := time.Now()

	timeDelay := 0

	for index, login := range inputArray {

		go func(userLogin string, delay int) {

			time.Sleep(time.Duration(delay) * time.Second)

			user, err := controller.GetUser(userLogin)
			if err != nil {
				channel <- UserResponse{Login: userLogin, User: nil, Error: err}
			} else {
				channel <- UserResponse{Login: userLogin, User: user, Error: nil}
			}

		}(login, timeDelay)

		timeDelay += 1

		if index == 100 {
			break
		}

	}

	var queryLength = len(inputArray)
	var numberOfResponse = 0

	for {
		select {
		case resp := <-channel:

			numberOfResponse++
			if resp.Error != nil {
				fmt.Println(resp.Error)
				fmt.Printf("%s %d/%d Trying to get user data from cache for: %s\n", prefix, numberOfResponse, queryLength, resp.Login)
				dbData, err := getUserFromCache(resp.Login)
				if err != nil {
					fmt.Printf("%s %d/%d Failed to get data of user: %s\n", prefix, numberOfResponse, queryLength, resp.Login)
				} else {
					querySuccessMessage(resp.Login, startTime, numberOfResponse, queryLength)
					queriedUsers = append(queriedUsers, dbData)
				}
			} else {
				querySuccessMessage(resp.Login, startTime, numberOfResponse, queryLength)
				queriedUsers = append(queriedUsers, resp.User)
			}

		case <-time.After(50 * time.Millisecond):
			break

		}

		if numberOfResponse == len(inputArray) {
			fmt.Printf("%s --- FINISHED QUERYING USERS ---\n", prefix)
			break
		}

	}

	return queriedUsers

}

func querySuccessMessage(login string, startTime time.Time, n int, overall int) {
	timeDiff := time.Since(startTime)
	timeDiffString := fmt.Sprintf("%fs", timeDiff.Seconds())
	fmt.Printf("%s %d/%d User: %s, Time: %s\n", prefix, n, overall, login, timeDiffString)
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
