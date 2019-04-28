package db

import (
	"errors"
	"log"
	"strconv"

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

func RedisHashInsert(client *redis.Client, key string, field string, value interface{}) error {

	err := client.HSet(key, field, value).Err()
	if err != nil {
		return err
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

//TODO: use RedisGet with error wrapper
func GetScoreParam(client *redis.Client, key string, field string) float64 {

	value, err := client.HGet(key, field).Result()

	if err == redis.Nil {
		err = client.HSet(key, field, 1.0).Err()
		if err != nil {
			log.Fatal(err)
		}
		return 1.0
	} else if err != nil {
		log.Fatal(err)
	} else {
		float, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal("false field type in redis hash!")
		}
		return float
	}

	return 1.0

}
