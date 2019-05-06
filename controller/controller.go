package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/m-lukas/github-analyser/db"
)

const timoutSeconds = 15

func GetUser(userName string) (*db.User, error) {

	user, err := fetchUser(userName)
	if err != nil {

		fmt.Println("Trying to retrieve user from cache.")
		user, err = GetUserFromCache(userName)
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
		err = CacheUser(user, "users")
		if err != nil {
			log.Println(err)
		}
	}(user)

	return user, nil
}
