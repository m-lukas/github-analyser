package app

import (
	"context"
	"testing"

	"github.com/m-lukas/github-analyser/db"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

//setupMongoTest initializes the mongo client and provides a cleared collection
func setupMongoTest(t *testing.T, root *db.DatabaseRoot, collectionName string, ctx context.Context) (*db.MongoClient, *mongo.Collection) {

	err := root.InitMongoClient()
	require.Nil(t, err)

	mongoClient := root.MongoClient

	require.Equal(t, mongoClient.Config.Enviroment, db.ENV_TEST) //check for right db config
	collection := mongoClient.Database.Collection(collectionName)

	clearMongoTestCollection(t, collection, ctx)

	return mongoClient, collection
}

//removes all documents from the specific collection
func clearMongoTestCollection(t *testing.T, collection *mongo.Collection, ctx context.Context) {
	err := collection.Drop(ctx)
	require.Nil(t, err, "droping of collection failed")
}

//setupElasticTest initializes the elastic client and with a cleared test index
func setupElasticTest(t *testing.T, root *db.DatabaseRoot, ctx context.Context) *db.ElasticClient {

	err := root.InitElasticClient()
	require.Nil(t, err)

	elasticClient := root.ElasticClient
	require.Equal(t, elasticClient.Config.Enviroment, db.ENV_TEST) //check for right db config

	index := elasticClient.Config.DefaultIndex

	_, err = elasticClient.Client.DeleteIndex(index).Do(ctx)
	require.Nil(t, err, "index deletion failed")

	err = elasticClient.CheckIndexes()
	require.Nil(t, err, "index creation failed")

	return elasticClient
}
