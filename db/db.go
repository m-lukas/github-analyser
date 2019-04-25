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

func getRoot() (*DatabaseRoot, error) {

	var err error

	err = checkDbRoot()
	if err != nil {
		return nil, err
	}

	return dbRoot, nil

}

func GetMongo() (*mongo.Database, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = root.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		err := root.initMongoClient()
		if err != nil {
			return nil, err
		}
	}

	client := root.MongoClient
	config := root.Config
	db := client.Database(config.MongoDatabaseName)
	return db, nil
}

func GetRedis() (*redis.Client, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	client := root.RedisClient
	_, err = client.Ping().Result()
	if err != nil {

		err = root.initRedisClient()
		if err != nil {
			return nil, err
		}

	}

	return client, nil
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

func checkDbRoot() error {

	if dbRoot == nil {
		err := Init()
		if err != nil {
			log.Println(err)
		}
	}

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
