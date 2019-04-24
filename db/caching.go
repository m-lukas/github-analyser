package db

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CacheUser(user *User) error {

	db, err := Get(DB_MONGO)
	if err != nil {
		return err
	}

	mongo, ok := db.(*mongo.Database)
	if ok != true {
		return errors.New("type conversion of database failed!")
	}

	collection := mongo.Collection("users")

	dbUser, err := FindUser(user.Login, collection)
	if dbUser != nil {
		filter := bson.D{{"login", user.Login}}
		err = UpdateAll(filter, user, collection)
		if err != nil {
			return err
		}

	} else {
		err = Insert(user, collection)
		if err != nil {
			return err
		}
	}

	return nil

}
