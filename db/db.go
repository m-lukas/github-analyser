package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DB_MONGO = "DB_MONGO"
	DB_REDIS = "DB_REDIS"
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

func Get(dbType string) (interface{}, error) {

	var err error

	if dbRoot == nil {
		err = Init()
		if err != nil {
			log.Println(err)
		}
	}

	switch dbType {
	case DB_MONGO:
		db, err := dbRoot.getMongo()
		if err != nil {
			return nil, err
		}
		return db, err
	case DB_REDIS:
		db, err := dbRoot.getRedis()
		if err != nil {
			return nil, err
		}
		return db, err
	default:
		return nil, errors.New("database does not exist!")
	}

}

func (root *DatabaseRoot) getMongo() (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := root.MongoClient.Ping(ctx, readpref.Primary())
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

func (root *DatabaseRoot) getRedis() (*redis.Client, error) {

	client := root.RedisClient

	_, err := client.Ping().Result()
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
