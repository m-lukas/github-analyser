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

	//check if user with login exists in collection
	dbUser, err := mongoClient.FindUser(user.Login, collectionName)
	if dbUser != nil {
		//updata user if existing
		filter := bson.D{{"login", user.Login}}
		err = mongoClient.UpdateAll(filter, user, collectionName)
		if err != nil {
			return err
		}

	} else {
		//insert new user into collection if not existing
		err = mongoClient.Insert(user, collectionName)
		if err != nil {
			return err
		}
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
