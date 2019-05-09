package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/m-lukas/github-analyser/util"

	"github.com/m-lukas/github-analyser/db"
)

const timoutSeconds = 15

func GetUser(userName string) (*db.User, error) {

	collectionName := "users"

	if util.IsTesting() {
		collectionName = "test_getuser"
		fmt.Println("TRUE")
	}

	user, err := fetchUser(userName)
	if err != nil {

		log.Println(err)
		fmt.Println("Trying to retrieve user from cache.")
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
			log.Println(err)
		}
	}(user)

	return user, nil
}
