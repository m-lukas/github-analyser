package setup

import (
	"errors"
	"time"

	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/graphql"
)

func SetupInputFile(rootUser string) error {

	inputArray, err := apiCallWrapper(rootUser)
	if err != nil {
		return err
	}

	err = util.WriteFile("./metrix/input/users.txt", inputArray)
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
			return nil, errors.New("Exceded number of tries!")
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
