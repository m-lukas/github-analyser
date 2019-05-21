package controller

import (
	"errors"
	"fmt"

	"github.com/m-lukas/github-analyser/logger"
	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/db"
)

const timoutSeconds = 15

func GetUser(userName string) (*db.User, error) {

	collectionName := "users"

	if util.IsTesting() {
		collectionName = "test_getuser"
	}

	user, err := fetchUser(userName)
	if err != nil {

		logger.ErrorNoMail(fmt.Sprintf("Failed to fetch data of user: %s, error %s", userName, err.Error()))
		logger.Warn(fmt.Sprintf("Trying to get user data from cache for: %s", userName))
		user, err = GetUserFromCache(userName, collectionName)
		if err != nil {
			return nil, err
		}
	}

	if user.Login == "" {
		return nil, errors.New("User does not exist!")
	}

	config, err := db.GetScoreConfig()
	if err != nil {
		return nil, err
	}

	SetScore(user, config)

	go func(user *db.User) {
		err = CacheUser(user, collectionName)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to cache user: %s, error: %s", user.Login, err.Error()))
		}
	}(user)

	return user, nil
}
