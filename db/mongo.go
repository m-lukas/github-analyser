package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/m-lukas/github-analyser/logger"
	"github.com/m-lukas/github-analyser/util"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MongoClient contains the mongo db client, its config and the default database
type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
	Config   *MongoConfig
}

//MongoConfig contains config to init mongo db client
type MongoConfig struct {
	MongoDatabaseName string
	MongoURI          string
	Enviroment        string
}

//getDefaultConfig return config in dev/prod
func (client *MongoClient) getDefaultConfig() *MongoConfig {
	return &MongoConfig{
		MongoDatabaseName: os.Getenv("MONGO_DB"),
		MongoURI:          getMongoURI(),
		Enviroment:        ENV_PROD,
	}
}

//getTestConfig return config in test
func (client *MongoClient) getTestConfig() *MongoConfig {
	return &MongoConfig{
		MongoDatabaseName: "core_test",
		// MongoURI:          "mongodb://user:hd63gdf5df5g@localhost:27018/admin",
		MongoURI:   "mongodb://localhost:27018/admin",
		Enviroment: ENV_TEST,
	}
}

//InitMongoClient establishes a connection to the mongoDB instance
func (root *DatabaseRoot) InitMongoClient() error {

	mongoClient := &MongoClient{}
	//use config according to the enviroment
	if util.IsTesting() {
		mongoClient.Config = mongoClient.getTestConfig()
	} else {
		mongoClient.Config = mongoClient.getDefaultConfig()
	}

	//add timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//try to establish connection to mongoDB Instance
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoClient.Config.MongoURI))
	if err != nil {
		return errors.Wrap(err, "mongo")
	}
	//ping instance, return err if not reachable
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	//add client and default database to struct
	mongoClient.Client = client
	mongoClient.Database = client.Database(mongoClient.Config.MongoDatabaseName)

	root.MongoClient = mongoClient
	logger.Info("Initialized mongo client!")

	return nil
}

/*
	Returns configurated URI string for MongoDB.
*/
func getMongoURI() (uri string) {
	dbHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_NAME")
	// dbUser := os.Getenv("MONGO_USER")
	// dbPass := os.Getenv("MONGO_PASS")

	// return fmt.Sprintf(`mongodb://%s:%s@%s/%s`, dbUser, dbPass, dbHost, dbName)
	return fmt.Sprintf(`mongodb://%s/%s`, dbHost, dbName)
}
