package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbRoot *DatabaseRoot

type DatabaseRoot struct {
	MongoClient *mongo.Client
	RedisClient *redis.Client
	Config      *DatabaseConfig
}

type DatabaseConfig struct {
	MongoDatabaseName string
	MongoURI          string
}

func defaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		MongoDatabaseName: "core",
		MongoURI:          getMongoURI(),
	}
}

func ExampleClient(client *redis.Client) {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func Init() error {
	var err error

	dbRoot = &DatabaseRoot{
		Config: defaultDatabaseConfig(),
	}

	err = dbRoot.initMongoClient()
	if err != nil {
		return err
	}

	err = dbRoot.initRedisClient()
	if err != nil {
		return err
	}

	return nil
}

func Get() *DatabaseRoot {

	if dbRoot == nil {
		err := Init()
		if err != nil {
			log.Println(err)
		}
	}
	return dbRoot

}

/*
	Initializes the mongoDB Client to access databases and collections.
*/
func (root *DatabaseRoot) initMongoClient() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Config.MongoURI))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	root.MongoClient = client
	log.Println("Initialized mongo client!")

	return nil
}

func (root *DatabaseRoot) initRedisClient() error {

	client := redis.NewClient(&redis.Options{
		Addr:     getRedisURI(),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	root.RedisClient = client
	log.Println("Initialized redis client!")

	return nil

}

/*
	Returns configurated URI string for MongoDB.
*/
func getMongoURI() (uri string) {
	dbHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_NAME")
	dbUser := os.Getenv("MONGO_USER")
	dbPass := os.Getenv("MONGO_PASS")

	return fmt.Sprintf(`mongodb://%s:%s@%s/%s`, dbUser, dbPass, dbHost, dbName)
}

func getRedisURI() (uri string) {
	dbHost := os.Getenv("REDIS_HOST")
	return dbHost
}
