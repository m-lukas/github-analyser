package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type aggregationResponseObject struct {
	Data []*aggregationResponse
}

type aggregationResponse struct {
	ID            primitive.ObjectID `bson:"_id"`
	Login         string             `bson:"login"`
	Name          string             `bson:"name"`
	ActivityScore float64            `bson:"activity_score"`
	NewField      string             `bson:"newField"`
}

func Test_Aggregation(t *testing.T) {

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//initialize mongo database and test collection
	db.TestRoot = &db.DatabaseRoot{}
	collectionName := "test_caching"
	mongoClient, collection := setupMongoTest(t, db.TestRoot, collectionName, ctx)

	testUser := &db.User{
		Login:         "m-lukas",
		Name:          "Lukas Müller",
		Location:      "Berlin",
		Repositories:  100,
		ActivityScore: 50.001,
	}

	//inser testUser into test collection
	err = mongoClient.Insert(testUser, collectionName)
	require.Nil(t, err, "mongo internal: insert failed")

	//check if insert of testdata was successful
	documents, err := mongoClient.FindAllUsers(collectionName)
	require.Nil(t, err, "mongo internal: find all failed")
	require.Equal(t, 1, len(documents), "mongo internal: insert failed")

	t.Run("Add(): stages aren't added", func(t *testing.T) {

		//create pipeline and add two stages
		pipeline := &db.Pipeline{}
		pipeline.Add(bson.D{{}})
		pipeline.Add(bson.D{{}})

		//check number of stages
		assert.Equal(t, 2, len(pipeline.Stages))

	})

	t.Run("Run(): aggregation execution failed", func(t *testing.T) {

		var result aggregationResponseObject

		//create pipeline
		pipeline := &db.Pipeline{}
		pipeline.Add(bson.D{{"$match", bson.D{{"login", "m-lukas"}}}})
		pipeline.Add(bson.D{{"$project", bson.D{{"login", 1}, {"name", 1}, {"activity_score", 1}, {"newField", "newField"}}}})

		//check num of stages
		require.Equal(t, 2, len(pipeline.Stages))

		//exec pipeline -> decode into result struct
		err = pipeline.Run(&result, collectionName)
		assert.Nil(t, err)

		//compare output with input
		data := result.Data
		assert.Equal(t, testUser.Login, data[0].Login)
		assert.Equal(t, testUser.Name, data[0].Name)
		assert.Equal(t, testUser.ActivityScore, data[0].ActivityScore)
		assert.Equal(t, "newField", data[0].NewField)

	})

	clearMongoTestCollection(t, collection, ctx)

}

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
