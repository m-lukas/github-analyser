package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/globalsign/mgo/bson"
)

func CacheUser(user *User) error {

	collection, err := Get("users")
	if err != nil {
		return err
	}

	dbUser, err := FindUser(user.Login, collection)
	if err != nil {
		log.Println(err)
	}
	if dbUser != nil {

		log.Println("db document exists")

	}

	err = Insert(user, collection)
	if err != nil {
		return err
	}

	return nil

}

func Insert(document interface{}, collection *mongo.Collection) error {

	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return nil

}

func FindUser(login string, collection *mongo.Collection) (*User, error) {

	collection, err := Get("users")
	if err != nil {
		return nil, err
	}

	result := collection.FindOne(context.Background(), bson.M{"login": login})
	if result.Err() != nil {
		return nil, result.Err()
	}

	user := &User{}
	err = result.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}
