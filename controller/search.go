package controller

import (
	"github.com/m-lukas/github-analyser/db"
)

//SearchUser searches multiple fields in the default elastic index for the given search query
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

	//unmarshal json response into ElasticUser list
	results, err := db.ConvertUsers(rawList)
	if err != nil {
		return nil, err
	}

	return results, nil
}
