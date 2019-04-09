package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbInstance *Database

type Database struct {
	MongoClient *mongo.Client
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
	dbInstance = &Database{
		Config: defaultDatabaseConfig(),
	}

	mongoClient, err := initMongoClient()
	if err != nil {
		return err
	}

	dbInstance.MongoClient = mongoClient
	return nil
}

/*
	Returns default mongo database.
*/
func Get() (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := dbInstance.MongoClient.Ping(ctx, readpref.Primary())

	if err == nil {
		client := dbInstance.MongoClient
		config := dbInstance.Config
		return client.Database(config.MongoDatabaseName), nil
	}

	mongoClient, err := initMongoClient()
	if err != nil {
		return nil, err
	}

	dbInstance.MongoClient = mongoClient

	client := dbInstance.MongoClient
	config := dbInstance.Config
	return client.Database(config.MongoDatabaseName), nil
}

/*
	Initializes the mongoDB Client to access databases and collections.
*/
func initMongoClient() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbInstance.Config.MongoURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	log.Println("Initialized mongo client!")
	return client, err

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
