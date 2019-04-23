package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	Returns default mongo database.
*/
func (root *DatabaseRoot) GetMongo(collectionName string) (*mongo.Collection, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := root.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		err := root.initMongoClient()
		if err != nil {
			return nil, err
		}
	}

	client := root.MongoClient
	config := root.Config
	db := client.Database(config.MongoDatabaseName)
	return db.Collection(collectionName), nil
}

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
