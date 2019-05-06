package db

import (
	"log"
	"os"
	"strconv"

	"github.com/m-lukas/github-analyser/util"

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
	Enviroment string
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
		Enviroment: ENV_PROD,
	}
}

func (client *RedisClient) getTestConfig() *RedisConfig {
	return &RedisConfig{
		URI:        "localhost:6379",
		Password:   "",
		DatabaseID: 1,
		Enviroment: ENV_TEST,
	}
}

func (root *DatabaseRoot) InitRedisClient() error {

	redisClient := &RedisClient{}
	if util.IsTesting() {
		redisClient.Config = redisClient.getTestConfig()
	} else {
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
