package db

import (
	"errors"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

func (redisClient *RedisClient) Insert(pairs map[string]interface{}) error {

	client := redisClient.Client
	for key, value := range pairs {

		err := client.Set(key, value, 0).Err()
		if err != nil {
			return err
		}

	}
	return nil
}

func (redisClient *RedisClient) HashInsert(key string, field string, value interface{}) error {

	client := redisClient.Client
	err := client.HSet(key, field, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisClient *RedisClient) Get(key string) (interface{}, error) {

	client := redisClient.Client
	value, err := client.Get(key).Result()

	if err == redis.Nil {
		return nil, errors.New("key does not exist!")
	} else if err != nil {
		return nil, err
	} else {
		return value, nil
	}

}

func (redisClient *RedisClient) GetScoreParam(key string, field string) float64 {

	client := redisClient.Client
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
