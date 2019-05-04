package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Redis(t *testing.T) {

	var err error

	root := &DatabaseRoot{}
	err = root.initRedisClient()
	require.Nil(t, err, "failed to initialize redis client")
	redisClient := root.RedisClient
	require.NotNil(t, redisClient, "failed to initialize redis client")

	err = redisClient.Client.FlushDB().Err()
	require.Nil(t, err, "flushing of database failed")

	testSlice := map[string]interface{}{
		"test1": "test",
		"test2": 2,
		"test3": 3.14,
	}

	t.Run("database functionality test", func(t *testing.T) {
		err = redisClient.Insert(testSlice)
		require.Nil(t, err, "inserting of testdata failed")

		for key, value := range testSlice {
			redisValue, err := redisClient.Get(key)
			require.Nil(t, err, "retrieving of test key failed")
			expected := fmt.Sprintf("%v", value)
			require.Equal(t, expected, redisValue)
		}

		keyCount, err := redisClient.Client.DBSize().Result()
		require.Nil(t, err, "redis internal: db size failed")
		require.Equal(t, int64(3), keyCount)

		err = redisClient.HashInsert("testHash1", "field1", "test")
		require.Nil(t, err, "hash insert failed")
		err = redisClient.HashInsert("testHash1", "score1", 3.14)
		require.Nil(t, err, "hash insert failed")
		err = redisClient.HashInsert("testHash2", "fieldName", 5)
		require.Nil(t, err, "hash insert failed")

		keyCount, err = redisClient.Client.DBSize().Result()
		require.Nil(t, err, "redis internal: db size failed")
		require.Equal(t, int64(5), keyCount)

		hashLength, err := redisClient.Client.HLen("testHash1").Result()
		require.Nil(t, err, "redis internal: hash length failed")
		require.Equal(t, int64(2), hashLength)

		scoreValue := redisClient.GetScoreParam("testHash1", "score1")
		require.Equal(t, 3.14, scoreValue)
		scoreValue = redisClient.GetScoreParam("testHash1", "score2") //does not exist -> return: 1.0
		require.Equal(t, 1.0, scoreValue)
	})

	err = redisClient.Client.FlushDB().Err()
	require.Nil(t, err, "flushing of database failed")
}
