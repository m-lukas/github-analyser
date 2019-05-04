package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/m-lukas/github-analyser/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
	Config   *MongoConfig
}

type MongoConfig struct {
	MongoDatabaseName string
	MongoURI          string
}

func (client *MongoClient) getDefaultConfig() *MongoConfig {
	return &MongoConfig{
		MongoDatabaseName: os.Getenv("MONGO_DB"),
		MongoURI:          getMongoURI(),
	}
}

func (client *MongoClient) getTestConfig() *MongoConfig {
	return &MongoConfig{
		MongoDatabaseName: "core_test",
		MongoURI:          "mongodb://user:hd63gdf5df5g@localhost:27018/admin",
	}
}

func (root *DatabaseRoot) initMongoClient() error {

	mongoClient := &MongoClient{}
	if util.IsTesting() {
		mongoClient.Config = mongoClient.getTestConfig()
	} else {
		mongoClient.Config = mongoClient.getDefaultConfig()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoClient.Config.MongoURI))
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	mongoClient.Client = client
	mongoClient.Database = client.Database(mongoClient.Config.MongoDatabaseName)

	root.MongoClient = mongoClient
	log.Println("Initialized mongo client!")

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
