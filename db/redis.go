package db

import (
	"log"
	"os"
	"strconv"

	"github.com/m-lukas/github-analyser/util"

	"github.com/go-redis/redis"
)

//RedisClient contains the redis db client and its config
type RedisClient struct {
	Client *redis.Client
	Config *RedisConfig
}

//RedisConfig contains config to init redis db client
type RedisConfig struct {
	URI        string
	Password   string
	DatabaseID int
	Enviroment string
}

//getDefaultConfig return config in dev/prod
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

//getTestConfig return config in test
func (client *RedisClient) getTestConfig() *RedisConfig {
	return &RedisConfig{
		URI:        "localhost:6379",
		Password:   "",
		DatabaseID: 1,
		Enviroment: ENV_TEST,
	}
}

//InitRedisClient establishes a connection to the redis instance
func (root *DatabaseRoot) InitRedisClient() error {

	redisClient := &RedisClient{}
	//assign config according to the enviroment
	if util.IsTesting() {
		redisClient.Config = redisClient.getTestConfig()
	} else {
		redisClient.Config = redisClient.getDefaultConfig()
	}

	//configurate client
	client := redis.NewClient(&redis.Options{
		Addr:     redisClient.Config.URI,
		Password: redisClient.Config.Password,
		DB:       redisClient.Config.DatabaseID,
	})

	//ping client
	_, err := client.Ping().Result()
	if err != nil {
		//not reachable
		return err
	}

	redisClient.Client = client
	root.RedisClient = redisClient
	log.Println("Initialized redis client!")

	return nil

}

/*
	Returns configurated URI string for Redis.
*/
func getRedisURI() (uri string) {
	dbHost := os.Getenv("REDIS_HOST")
	return dbHost
}
