package db

import (
	"go.mongodb.org/mongo-driver/bson"
)

func CacheUser(user *User, collectionName string) error {

	client, err := GetMongo()
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
