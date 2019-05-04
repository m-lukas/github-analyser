package db

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	Client *redis.Client
	Config *RedisConfig
}

type RedisConfig struct {
	URI        string
	Password   string
	DatabaseID int
}

func (client *RedisClient) getDefaultConfig() *RedisConfig {
	databaseID, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		databaseID = 0
	}

	return &RedisConfig{
		URI:        getRedisURI(),
		Password:   os.Getenv("REDIS_PASS"),
		DatabaseID: databaseID,
	}
}

func (client *RedisClient) getTestConfig() *RedisConfig {
	return &RedisConfig{
		URI:        "localhost:6379",
		Password:   "",
		DatabaseID: 1,
	}
}

func (root *DatabaseRoot) initRedisClient() error {

	redisClient := &RedisClient{}
	switch root.Enviroment {
	case ENV_TEST:
		redisClient.Config = redisClient.getTestConfig()
	default:
		redisClient.Config = redisClient.getDefaultConfig()
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisClient.Config.URI,
		Password: redisClient.Config.Password,
		DB:       redisClient.Config.DatabaseID,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	redisClient.Client = client
	root.RedisClient = redisClient
	log.Println("Initialized redis client!")

	return nil

}

func getRedisURI() (uri string) {
	dbHost := os.Getenv("REDIS_HOST")
	return dbHost
}
