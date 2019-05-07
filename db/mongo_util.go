package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateAll overrides all fields of document
func (client *MongoClient) UpdateAll(filter []primitive.E, document interface{}, collectionName string) error {

	collection := client.Database.Collection(collectionName)

	//check filter length to avoid entire collection override
	if len(filter) == 0 {
		return errors.New("filter must not be empty!")
	}

	//update document in collection
	_, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": document})
	if err != nil {
		return err
	}

	return nil

}

//Insert inserts a document into the collection, without duplicate check
func (client *MongoClient) Insert(document interface{}, collectionName string) error {

	collection := client.Database.Collection(collectionName)

	//insert document into collection
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return nil
}

//FindAllUsers retrieves all documents in the collection and decodes them into User structs
func (client *MongoClient) FindAllUsers(collectionName string) ([]*User, error) {

	//result slice
	var userArray []*User
	collection := client.Database.Collection(collectionName)
	ctx := context.Background()

	//find all documents -> empty filter
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	//loop through cursor
	for cursor.Next(ctx) {

		//decode into &User{}
		user := &User{}
		err = cursor.Decode(user)
		if err != nil {
			return nil, err
		}

		//exclude fault documents
		if user.Login != "" {
			userArray = append(userArray, user)
		}
	}

	return userArray, nil
}

//FindUser finds a user by login in a given collection
func (client *MongoClient) FindUser(login string, collectionName string) (*User, error) {

	collection := client.Database.Collection(collectionName)

	//set filter on login, find first document
	result := collection.FindOne(context.Background(), bson.M{"login": login})
	if result.Err() != nil {
		return nil, result.Err()
	}

	//decode into &User{} struct
	user := &User{}
	err := result.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}
