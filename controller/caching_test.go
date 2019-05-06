package controller

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

//INTEGRATION TEST FOR CACHING
func Test_Caching(t *testing.T) {

	var err error

	dbRoot = &DatabaseRoot{}

	testUser := &User{
		Login:         "m-lukas",
		Name:          "Lukas Müller",
		Location:      "Berlin",
		Repositories:  100,
		ActivityScore: 50.001,
	}

	err = dbRoot.initMongoClient()
	require.Nil(t, err)

	mongoClient := dbRoot.MongoClient

	require.Equal(t, mongoClient.Config.Enviroment, ENV_TEST) //check for right db config

	collectionName := "users_test"
	collection := mongoClient.Database.Collection(collectionName)

	err = collection.Drop(context.Background())
	require.Nil(t, err, "droping of collection failed")

	t.Run("user couldn't be cached", func(t *testing.T) {
		err = CacheUser(testUser, collectionName)
		require.Nil(t, err, "failed to cache user")

		err = CacheUser(testUser, collectionName)
		require.Nil(t, err, "failed to cache user")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		count, err := collection.CountDocuments(ctx, bson.M{"login": "m-lukas"})
		assert.Nil(t, err, "mongo internal: failed to count documents")
		assert.Equal(t, int64(1), count, "didn't update existing user")

		var dbUser User

		err = collection.FindOne(ctx, bson.M{}).Decode(&dbUser)
		assert.Nil(t, err, "failed to get user from database")
		assert.Equal(t, testUser.Login, dbUser.Login, "user wasn't saved properly")
	})

	err = collection.Drop(context.Background())
	assert.Nil(t, err, "droping of collection failed")

	dbRoot = nil
}
