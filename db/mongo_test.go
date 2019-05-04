package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_Mongo(t *testing.T) {

	var err error

	root := &DatabaseRoot{}
	err = root.initMongoClient()
	require.Nil(t, err, "failed to initialize mongo client")
	mongoClient := root.MongoClient
	require.NotNil(t, mongoClient, "failed to initialize mongo client")

	collectionName := "users_test"
	err = mongoClient.Database.Collection(collectionName).Drop(context.Background())
	require.Nil(t, err, "droping of collection failed")

	testSlice := []*User{
		{
			Login:         "m-lukas",
			Name:          "Lukas Müller",
			Email:         "lukas@test.com",
			ActivityScore: 50.0,
			Repositories:  400,
		},
		{
			Login:         "Urhengulas",
			Name:          "Johann Hemmann",
			Email:         "johann@test.com",
			ActivityScore: 23.999999,
			Repositories:  67,
		},
		{
			Login:         "sindresorhus",
			Name:          "sindresorhus",
			Email:         "sth@sth.com",
			ActivityScore: 1000000.0,
			Repositories:  1000000,
		},
	}

	t.Run("database functionality test", func(t *testing.T) {
		for _, user := range testSlice {
			err := mongoClient.Insert(user, collectionName)
			require.Nil(t, err, "insert failed")
		}

		testUser := "m-lukas"
		update := &User{
			Login:         testUser,
			Name:          "Lukas Müller",
			Email:         "lukas@test.com",
			ActivityScore: 50.0,
			Repositories:  200,
		}

		filter := bson.D{{"login", testSlice[0].Login}}
		err := mongoClient.UpdateAll(filter, update, collectionName)
		require.Nil(t, err, "update failed")

		retrivedUser, err := mongoClient.FindUser(testUser, collectionName)
		require.Nil(t, err, "user not found")
		require.Equal(t, update.Repositories, retrivedUser.Repositories, "update hasn't update document")

		userSlice, err := mongoClient.FindAllUsers(collectionName)
		require.Nil(t, err, "find all failed")
		require.Equal(t, len(testSlice), len(userSlice))
	})

	err = mongoClient.Database.Collection(collectionName).Drop(context.Background())
	require.Nil(t, err, "droping of collection failed")
}
