package controller

import (
	"errors"
	"log"

	"github.com/m-lukas/github-analyser/db"
)

const timoutSeconds = 15

func GetUser(userName string) (*db.User, error) {

	user, err := fetchUser(userName)
	if err != nil {
		return nil, err
	}

	if user.Login == "" {
		return nil, errors.New("User does not exist!")
	}

	config, err := db.GetScoreConfig()
	if err != nil {
		return nil, err
	}

	user.Scores = CalcScores(user, config)

	user.ActivityScore = CalcActivityScore(user.Scores, config)
	user.PopularityScore = CalcPopularityScore(user.Scores, config)

	go func(user *db.User) {
		err = CacheUser(user, "users")
		if err != nil {
			log.Println(err)
		}
	}(user)

	return user, nil

}
