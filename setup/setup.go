package setup

import (
	"errors"
	"fmt"
	"time"

	"github.com/m-lukas/github-analyser/logger"
	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/graphql"
)

const prefix = "SETUP |"

func SetupInputFile(rootUser string) error {

	inputArray, err := apiCallWrapper(rootUser)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%s Successfully retrived data from GitHub!", prefix))

	err = util.WriteFile(fmt.Sprintf("./metrix/input/%s.txt", rootUser), inputArray)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%s Successfully wrote input files!", prefix))

	return nil

}

func apiCallWrapper(rootUser string) ([]string, error) {

	var err error
	var inputArray []string
	var tryCount = 0
	var maxTries = 5
	var cooldown = 5

	inputArray, err = apiCall(rootUser)

	for {

		tryCount++
		inputArray, err = apiCall(rootUser)

		if tryCount == maxTries {
			return nil, errors.New("Exceeded number of tries in setup!")
		}

		if err == nil {
			break
		}

		logger.Warn(fmt.Sprintf("%s Failed to fetch data! COOLDOWN: %d seconds [%d/%d]", prefix, cooldown, tryCount, maxTries))
		time.Sleep(time.Duration(cooldown) * time.Second)

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
