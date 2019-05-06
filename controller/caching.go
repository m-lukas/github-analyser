package controller

import (
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
)

func CacheUser(user *db.User, collectionName string) error {

	client, err := db.GetMongo()
	if err != nil {
		return err
	}

	dbUser, err := client.FindUser(user.Login, collectionName)
	if dbUser != nil {
		filter := bson.D{{"login", user.Login}}
		err = client.UpdateAll(filter, user, collectionName)
		if err != nil {
			return err
		}

	} else {
		err = client.Insert(user, collectionName)
		if err != nil {
			return err
		}
	}

	return nil
}
