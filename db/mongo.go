package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	Returns default mongo database.
*/

func UpdateAll(filter []primitive.E, document interface{}, collection *mongo.Collection) error {

	if len(filter) == 0 {
		return errors.New("filter must not be empty!")
	}

	_, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": document})
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

func FindAllUsers(collection *mongo.Collection) ([]*User, error) {

	var userArray []*User
	ctx := context.Background()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		user := &User{}
		err = cursor.Decode(user)
		if err != nil {
			return nil, err
		}

		if user.Login != "" {

			userArray = append(userArray, user)

		}
	}

	return userArray, nil

}

func FindUser(login string, collection *mongo.Collection) (*User, error) {

	result := collection.FindOne(context.Background(), bson.M{"login": login})
	if result.Err() != nil {
		return nil, result.Err()
	}

	user := &User{}
	err := result.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}
