package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	db.TestRoot = &db.DatabaseRoot{}

	err = db.TestRoot.InitMongoClient()
	require.Nil(t, err)

	mongoClient := db.TestRoot.MongoClient

	require.Equal(t, mongoClient.Config.Enviroment, db.ENV_TEST) //check for right db config

	collectionName := "test_aggregation"
	collection := mongoClient.Database.Collection(collectionName)

	err = collection.Drop(context.Background())
	require.Nil(t, err, "droping of collection failed")

	testUser := &db.User{
		Login:         "m-lukas",
		Name:          "Lukas Müller",
		Location:      "Berlin",
		Repositories:  100,
		ActivityScore: 50.001,
	}

	err = mongoClient.Insert(testUser, collectionName)
	require.Nil(t, err, "setup: insert failed")

	documents, err := mongoClient.FindAllUsers(collectionName)
	require.Nil(t, err, "setup: find all failed")
	require.Equal(t, 1, len(documents), "setup: insert failed")

	t.Run("Add(): stages aren't added", func(t *testing.T) {

		pipeline := &db.Pipeline{}
		pipeline.Add(bson.D{{}})
		pipeline.Add(bson.D{{}})

		assert.Equal(t, 2, len(pipeline.Stages))

	})

	t.Run("Run(): aggregation execution failed", func(t *testing.T) {

		var result aggregationResponseObject

		pipeline := &db.Pipeline{}
		pipeline.Add(bson.D{{"$match", bson.D{{"login", "m-lukas"}}}})
		pipeline.Add(bson.D{{"$project", bson.D{{"login", 1}, {"name", 1}, {"activity_score", 1}, {"newField", "newField"}}}})

		require.Equal(t, 2, len(pipeline.Stages))

		err = pipeline.Run(&result, collectionName)
		assert.Nil(t, err)

		data := result.Data
		assert.Equal(t, testUser.Login, data[0].Login)
		assert.Equal(t, testUser.Name, data[0].Name)
		assert.Equal(t, testUser.ActivityScore, data[0].ActivityScore)
		assert.Equal(t, "newField", data[0].NewField)

	})

	err = collection.Drop(context.Background())
	require.Nil(t, err, "droping of collection failed")

}
