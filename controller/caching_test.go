package controller

import (
	"context"
	"testing"
	"time"

	"github.com/m-lukas/github-analyser/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

//INTEGRATION TEST FOR CACHING
func Test_Caching(t *testing.T) {

	var err error

	db.TestRoot = &db.DatabaseRoot{}

	testUser := &db.User{
		Login:         "m-lukas",
		Name:          "Lukas Müller",
		Location:      "Berlin",
		Repositories:  100,
		ActivityScore: 50.001,
	}

	err = db.TestRoot.InitMongoClient()
	require.Nil(t, err)

	mongoClient := db.TestRoot.MongoClient

	require.Equal(t, mongoClient.Config.Enviroment, db.ENV_TEST) //check for right db config

	collectionName := "test_caching"
	collection := mongoClient.Database.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.Drop(ctx)
	require.Nil(t, err, "droping of collection failed")

	t.Run("user couldn't be cached", func(t *testing.T) {
		err = CacheUser(testUser, collectionName)
		require.Nil(t, err, "failed to cache user")

		err = CacheUser(testUser, collectionName)
		require.Nil(t, err, "failed to cache user")

		count, err := collection.CountDocuments(ctx, bson.M{"login": testUser.Login})
		assert.Nil(t, err, "mongo internal: failed to count documents")
		assert.Equal(t, int64(1), count, "didn't update existing user")

		var dbUser db.User

		err = collection.FindOne(ctx, bson.M{}).Decode(&dbUser)
		assert.Nil(t, err, "failed to get user from database")
		assert.Equal(t, testUser.Login, dbUser.Login, "user wasn't saved properly")
	})

	count, err := collection.CountDocuments(ctx, bson.M{"login": testUser.Login})
	assert.Nil(t, err, "mongo internal: failed to count documents")
	require.Equal(t, int64(1), count, "something went wrong while caching")

	t.Run("user couldn't retrive user from cache", func(t *testing.T) {

		user, err := GetUserFromCache(testUser.Login, collectionName)
		assert.Nil(t, err, "failed to get user from cache")
		assert.Equal(t, user.Login, testUser.Login)
		assert.Equal(t, user.Name, testUser.Name)
		assert.Equal(t, user.Location, testUser.Location)
		assert.Equal(t, user.Repositories, testUser.Repositories)
		assert.Equal(t, user.ActivityScore, testUser.ActivityScore)

	})

	err = collection.Drop(ctx)
	assert.Nil(t, err, "droping of collection failed")

	db.TestRoot = nil
}
