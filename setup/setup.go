package setup

import (
	"errors"
	"fmt"
	"time"

	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/graphql"
)

const prefix = "SETUP |"

func SetupInputFile(rootUser string) error {

	inputArray, err := apiCallWrapper(rootUser)
	if err != nil {
		return err
	}

	fmt.Printf("%s Successfully retrived data from GitHub!\n", prefix)

	err = util.WriteFile(fmt.Sprintf("./metrix/input/%s.txt", rootUser), inputArray)
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
	var cooldown = 5

	inputArray, err = apiCall(rootUser)

	for err != nil {

		tryCount++
		time.Sleep(time.Duration(cooldown) * time.Second)
		inputArray, err = apiCall(rootUser)

		if err != nil {
			fmt.Printf("%s Failed to fetch data! COOLDOWN: %d seconds\n", prefix, cooldown)
		}

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
