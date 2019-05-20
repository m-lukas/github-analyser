package controller

import (
	"github.com/m-lukas/github-analyser/db"
)

func SearchUser(query string) ([]*db.ElasticUser, error) {

	elasticClient, err := db.GetElastic()
	if err != nil {
		return nil, err
	}
	elasticIndex := elasticClient.Config.DefaultIndex

	rawList, err := elasticClient.Search(query, elasticIndex, "login", "email", "name", "bio")
	if err != nil {
		return nil, err
	}

	results, err := db.ConvertUsers(rawList)
	if err != nil {
		return nil, err
	}

	return results, nil
}
