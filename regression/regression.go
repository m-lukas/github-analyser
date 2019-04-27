package regression

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"

	"github.com/m-lukas/github-analyser/util"
)

const prefix = "REGRESSION:"

type UserResponse struct {
	Login string
	User  *db.User
	Error error
}

func PopulateDatabase() error {

	inputArray, err := readInput([]string{"./regression/input/users.txt"})
	if err != nil {
		return err
	}

	userArray := queryUserData(inputArray)

	fmt.Println(len(userArray))

	return nil
}

func readInput(filepathes []string) ([]string, error) {

	var inputArray []string
	var successfulFileCount = 0

	for index, filepath := range filepathes {

		fmt.Printf("%s Started scanning of file: %s\n", prefix, filepath)
		fmt.Printf("%s Scanning file: %d/%d\n", prefix, index+1, len(filepathes))

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

	fmt.Printf("%s STARTING TO QUERY USERS.\n", prefix)

	var queriedUsers []*db.User
	channel := make(chan UserResponse)
	startTime := time.Now()

	for _, login := range inputArray {

		go func(userLogin string) {

			user, err := controller.GetUser(userLogin)
			if err != nil {
				channel <- UserResponse{Login: userLogin, User: nil, Error: err}
			} else {
				channel <- UserResponse{Login: userLogin, User: user, Error: nil}
			}

		}(login)

	}

	var numberOfResponses = 0

	for {
		select {
		case resp := <-channel:

			numberOfResponses++
			if resp.Error != nil {
				fmt.Printf("%s Trying to get user data from cache for: %s\n", prefix, resp.Login)
				dbData, err := getUserFromCache(resp.Login)
				if err != nil {
					fmt.Printf("%s Failed to get data of user: %s\n", prefix, resp.Login)
				} else {
					successMessage(resp.Login, startTime)
					queriedUsers = append(queriedUsers, dbData)
				}
			} else {
				successMessage(resp.Login, startTime)
				queriedUsers = append(queriedUsers, resp.User)
			}

		case <-time.After(50 * time.Millisecond):
			break

		}

		if numberOfResponses == len(inputArray) {
			fmt.Printf("%s FINISHED QUERYING USERS\n", prefix)
			break
		}

	}

	return queriedUsers

}

func successMessage(login string, startTime time.Time) {
	timeDiff := time.Since(startTime)
	timeDiffString := fmt.Sprintf("%fs", timeDiff.Seconds())
	fmt.Printf("%s User: %s, Time: %s\n", prefix, login, timeDiffString)
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
