package controller

import (
	"context"
	"testing"

	"github.com/m-lukas/github-analyser/db"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupMongoTest(t *testing.T, root *db.DatabaseRoot, collectionName string, ctx context.Context) (*db.MongoClient, *mongo.Collection) {

	err := root.InitMongoClient()
	require.Nil(t, err)

	mongoClient := root.MongoClient

	require.Equal(t, mongoClient.Config.Enviroment, db.ENV_TEST) //check for right db config
	collection := mongoClient.Database.Collection(collectionName)

	clearMongoTestCollection(t, collection, ctx)

	return mongoClient, collection
}

func clearMongoTestCollection(t *testing.T, collection *mongo.Collection, ctx context.Context) {
	err := collection.Drop(ctx)
	require.Nil(t, err, "droping of collection failed")
}
