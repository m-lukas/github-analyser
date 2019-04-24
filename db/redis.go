package db

import (
	"errors"

	"github.com/go-redis/redis"
)

func RedisInsert(client *redis.Client, pairs map[string]interface{}) error {

	for key, value := range pairs {

		err := client.Set(key, value, 0).Err()
		if err != nil {
			return err
		}

	}

	return nil

}

func RedisGet(client *redis.Client, key string) (interface{}, error) {

	value, err := client.Get(key).Result()

	if err == redis.Nil {
		return nil, errors.New("key does not exist!")
	} else if err != nil {
		return nil, err
	} else {
		return value, nil
	}

}
