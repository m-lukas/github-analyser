package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (client *MongoClient) UpdateAll(filter []primitive.E, document interface{}, collectionName string) error {

	collection := client.Database.Collection(collectionName)

	if len(filter) == 0 {
		return errors.New("filter must not be empty!")
	}

	_, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": document})
	if err != nil {
		return err
	}

	return nil

}

func (client *MongoClient) Insert(document interface{}, collectionName string) error {

	collection := client.Database.Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return nil

}

func (client *MongoClient) FindAllUsers(collectionName string) ([]*User, error) {

	var userArray []*User
	collection := client.Database.Collection(collectionName)
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

func (client *MongoClient) FindUser(login string, collectionName string) (*User, error) {

	collection := client.Database.Collection(collectionName)

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
