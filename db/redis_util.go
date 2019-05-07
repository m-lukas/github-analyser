package db

import (
	"errors"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

//Insert inserts a new key-value pair / overrides key if already exist
func (redisClient *RedisClient) Insert(pairs map[string]interface{}) error {

	client := redisClient.Client
	for key, value := range pairs {

		//SET operation for every pair
		err := client.Set(key, value, 0).Err()
		if err != nil {
			return err
		}

	}
	return nil
}

//HashInsert inserts a field-value pair into the hash value of a key
func (redisClient *RedisClient) HashInsert(key string, field string, value interface{}) error {

	//get client and perform HSET
	client := redisClient.Client
	err := client.HSet(key, field, value).Err()
	if err != nil {
		return err
	}
	return nil
}

//Get retrieves the value of a specific key and returns err if not existing
func (redisClient *RedisClient) Get(key string) (interface{}, error) {

	client := redisClient.Client
	//perform GET with given key
	value, err := client.Get(key).Result()

	if err == redis.Nil {
		return nil, errors.New("key does not exist!")
	} else if err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

//GetScoreParam is a helper function to get the HashValue of a given key in combination with a given value.
//Sets the field and return 1.0 if not existing!
func (redisClient *RedisClient) GetScoreParam(key string, field string) float64 {

	client := redisClient.Client
	//perform HGET
	value, err := client.HGet(key, field).Result()

	//HSET to 1.0 if not existing and return 1.0, or err if internal failure
	if err == redis.Nil {
		err = client.HSet(key, field, 1.0).Err()
		if err != nil {
			log.Fatal(err)
		}
		return 1.0
	} else if err != nil {
		log.Fatal(err)
	} else {
		//get value if it exists and parse to float64
		float, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal("false field type in redis hash!")
		}
		return float
	}

	return 1.0
}
