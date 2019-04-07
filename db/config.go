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

var (
	globalClient       *mongo.Client
	globalDatabase     *mongo.Database
	globalDatabaseType *string
)

const (
	mongoDatabaseName = "core"
)

/*
	Returns default mongo database.
*/
func Get() (*mongo.Database, error) {

	if globalDatabase == nil {
		client, err := initMongoClient()
		if err != nil {
			return nil, err
		}
		globalDatabase = client.Database(mongoDatabaseName)
		log.Println("Initialized mongo database!")
		return globalDatabase, nil
	}
	return globalDatabase, nil

}

/*
	Initializes the mongoDB Client to access databases and collections.
*/
func initMongoClient() (*mongo.Client, error) {

	if globalClient == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() //what does this do?
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoDBConfig()))
		if err != nil {
			return nil, err
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			return nil, err
		}

		globalClient = client
		log.Println("Initialized mongo client!")
		return globalClient, nil
	}
	return globalClient, nil

}

/*
	Returns configurated URI string for MongoDB.
*/
func getMongoDBConfig() (uri string) {
	dbHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_NAME")
	dbUser := os.Getenv("MONGO_USER")
	dbPass := os.Getenv("MONGO_PASS")

	return fmt.Sprintf(`mongodb://%s:%s@%s/%s`, dbUser, dbPass, dbHost, dbName)
}

func getDatabaseType() *string {
	if globalDatabaseType == nil {
		databaseType := os.Getenv("DATABASE_TYPE")
		globalDatabaseType = &databaseType
	}
	return globalDatabaseType
}
