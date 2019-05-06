package controller

import (
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
)

func CacheUser(user *db.User, collectionName string) error {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return err
	}

	dbUser, err := mongoClient.FindUser(user.Login, collectionName)
	if dbUser != nil {
		filter := bson.D{{"login", user.Login}}
		err = mongoClient.UpdateAll(filter, user, collectionName)
		if err != nil {
			return err
		}

	} else {
		err = mongoClient.Insert(user, collectionName)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetUserFromCache(login string) (*db.User, error) {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return nil, err
	}
	collectionName := "users"

	dbUser, err := mongoClient.FindUser(login, collectionName)
	if err != nil {
		return nil, err
	}

	return dbUser, nil

}
