package db

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Elastic(t *testing.T) {

	var err error

	root := &DatabaseRoot{}
	var elasticClient *ElasticClient

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("InitElasticClient(): elastic initialization failed", func(t *testing.T) {
		err = root.InitElasticClient()
		fmt.Println(err)
		require.Nil(t, err, "failed to initialize client")

		elasticClient = root.ElasticClient
		require.NotNil(t, elasticClient, "failed to initialize client")

		client := elasticClient.Client
		config := elasticClient.Config

		//ping client
		_, _, err = client.Ping(config.ElasticURI).Do(ctx)
		require.Nil(t, err, "initialized mongo database not reachable")
	})

	//check config for futher operations on the database
	require.Equal(t, elasticClient.Config.Enviroment, ENV_TEST) //check for right db config

	index := "user_test_index"
	client := elasticClient.Client

	exists, err := client.IndexExists(index).Do(ctx)
	require.Nil(t, err, "index check failed")

	//drop test collection
	if exists {
		_, err = client.DeleteIndex(index).Do(ctx)
		require.Nil(t, err, "index deletion failed")
	}

	err = elasticClient.checkIndexes()
	require.Nil(t, err, "index creation failed")

	testSlice := []*ElasticUser{
		{
			Login: "m-lukas",
			Email: "lukas@test.com",
			Name:  "Lukas Müller",
			Bio:   "This is written in code",
		},
		{
			Login: "Urhengulas",
			Email: "johann@test.com",
			Name:  "Johann Hemmann",
			Bio:   "Bananas are yellow!",
		},
		{
			Login: "sindresorhus",
			Email: "sth@sth.com",
			Name:  "sindresorhus",
			Bio:   "CODE or not to CODE, that is the question.",
		},
	}

	t.Run("ELASTIC UTIL: database functionality test", func(t *testing.T) {
		//insert all users of test array
		for sliceIndex, user := range testSlice {
			id, err := elasticClient.Insert(user, index)
			fmt.Println(err)
			require.Nil(t, err, "insert failed")
			require.NotEqual(t, "", id, "didn't return id @ insert")
			testSlice[sliceIndex].Id = id
		}

		//testuser with changes
		testUserID := testSlice[0].Id
		updateBio := "This was modified by a test"

		id, err := elasticClient.Update(map[string]interface{}{"bio": updateBio}, testUserID, index)
		require.Nil(t, err, "update failed")
		require.Equal(t, testUserID, id, "didn't return id @ update")

		time.Sleep(1 * time.Second)

		results, err := elasticClient.Search("code", index, "bio")
		require.Nil(t, err, "search failed")

		var dataSlice []*ElasticUser
		for _, message := range results {
			var userData ElasticUser
			err = json.Unmarshal(message, &userData)
			require.Nil(t, err, "can't unmarshal search result")
			dataSlice = append(dataSlice, &userData)
		}

		require.Equal(t, 1, len(dataSlice), "not enough or too many results")
		require.Equal(t, "sindresorhus", dataSlice[0].Login, "search returns wrong results")

	})

	_, err = client.DeleteIndex(index).Do(ctx)
	require.Nil(t, err, "index deletion failed")

}
