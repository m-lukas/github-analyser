package controller

import (
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
)

//CacheUser saves all fields of the user into the given collection
func CacheUser(user *db.User, collectionName string) error {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return err
	}
	elasticClient, err := db.GetElastic()
	if err != nil {
		return err
	}
	elasticIndex := elasticClient.Config.DefaultIndex

	//check if user with login exists in collection
	mongoUser, _ := mongoClient.FindUser(user.Login, collectionName)
	//not found throws unspecific error as well as connection issues

	elasticResultList, err := elasticClient.Search(user.Login, elasticIndex, "login")
	if err != nil {
		return err
	}

	elasticUser := &db.ElasticUser{
		Login: user.Login,
		Email: user.Email,
		Name:  user.Name,
		Bio:   user.Bio,
	}

	if mongoUser == nil && len(elasticResultList) == 0 {

		elasticId, err := elasticClient.Insert(elasticUser, elasticIndex)
		if err != nil {
			return err
		}

		user.ElasticID = elasticId

		//insert new user into collection if not existing
		err = mongoClient.Insert(user, collectionName)
		if err != nil {
			return err
		}

	}

	if dbUser != nil {
		//updata user if existing
		filter := bson.D{{"login", user.Login}}
		err = mongoClient.UpdateAll(filter, user, collectionName)
		if err != nil {
			return err
		}

	} else {

	}

	return nil
}

//GetUserFromCache retrieves a user from the given collection by login
func GetUserFromCache(login string, collectionName string) (*db.User, error) {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return nil, err
	}

	dbUser, err := mongoClient.FindUser(login, collectionName)
	if err != nil {
		return nil, err
	}

	return dbUser, nil

}
