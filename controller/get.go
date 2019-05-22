package controller

import (
	"errors"
	"fmt"

	"github.com/m-lukas/github-analyser/logger"
	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/db"
)

const timoutSeconds = 15

//GetUser fetches the user data from github and caches it or retrieves from cache if possible
func GetUser(userName string) (*db.User, error) {

	//define collections manually since there is no access to client
	collectionName := "users"
	if util.IsTesting() {
		collectionName = "test_getuser"
	}

	//fetch user data
	user, err := fetchUser(userName)
	if err != nil {

		//try to get user from cache if fetching fails
		logger.ErrorNoMail(fmt.Sprintf("Failed to fetch data of user: %s, error %s", userName, err.Error()))
		logger.Warn(fmt.Sprintf("Trying to get user data from cache for: %s", userName))
		user, err = GetUserFromCache(userName, collectionName)
		if err != nil {
			//return err if user can't be retrieved from cache
			return nil, err
		}
	}

	//if user doesn't exist, there is no error but "Login" field is empty
	if user.Login == "" {
		return nil, errors.New("User does not exist!")
	}

	config, err := db.GetScoreConfig()
	if err != nil {
		return nil, err
	}

	//update score on user obj
	SetScore(user, config)

	//cache user in async process
	go func(user *db.User) {
		err = CacheUser(user, collectionName)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to cache user: %s, error: %s", user.Login, err.Error()))
		}
	}(user)

	return user, nil
}
