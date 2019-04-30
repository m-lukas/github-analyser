package setup

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/m-lukas/github-analyser/graphql"
)

const prefix = "SETUP |"

func SetupInputFile(rootUser string) error {

	inputArray, err := apiCallWrapper(rootUser)
	if err != nil {
		return err
	}

	err = writeFile("./metrix/input/users.txt", inputArray)
	if err != nil {
		return err
	}

	return nil

}

func apiCallWrapper(rootUser string) ([]string, error) {

	var err error
	var inputArray []string
	var tryCount = 0
	var maxTries = 5

	inputArray, err = apiCall(rootUser)

	for err != nil {

		tryCount++
		time.Sleep(5 * time.Second)
		inputArray, err = apiCall(rootUser)

		if tryCount == maxTries {
			return nil, errors.New("Exceeded number of tries!")
		}

	}

	return inputArray, nil

}

func apiCall(rootUser string) ([]string, error) {

	inputArray, err := graphql.GetPopulatingData(rootUser)
	if err != nil {
		return nil, err
	}

	return inputArray, nil

}

func writeFile(filepath string, input []string) error {

	file, err := os.Create(filepath)
	if err != nil {
		file.Close()
		return err
	}

	for _, login := range input {

		_, err = fmt.Fprintln(file, login)
		if err != nil {
			return err
		}

	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil

}
