package db

import (
	"context"
	"errors"
	"time"

	"github.com/m-lukas/github-analyser/util"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//dbRoot singleton
var dbRoot *DatabaseRoot

//TestRoot (only for testing purposes)
var TestRoot *DatabaseRoot

const (
	//ENV_PROD is config flag for prod env
	ENV_PROD = "ENV_PROD"
	//ENV_TEST is config flag for prod env
	ENV_TEST = "ENV_TEST"
)

//DatabaseRootstructure contains all database clients
type DatabaseRoot struct {
	MongoClient *MongoClient
	RedisClient *RedisClient
	ScoreConfig *ScoreParams
}

//Init initializes the dbRoot and triggers the initialization of all db clients
func Init() error {
	var err error

	//break in testing enviroment to avoid cross-execution of funtions
	if util.IsTesting() {
		return errors.New("not available in testing enviroment")
	}

	dbRoot = &DatabaseRoot{}

	err = dbRoot.InitMongoClient()
	if err != nil {
		return err
	}

	err = dbRoot.InitRedisClient()
	if err != nil {
		return err
	}

	err = dbRoot.InitScoreConfig()
	if err != nil {
		return err
	}

	return nil
}

//GetMongo pings the mongo client and returns a *MongoClient object if reachable
func GetMongo() (*MongoClient, error) {
	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	//check if db config has right enviroment flag
	if util.IsTesting() && root.MongoClient.Config.Enviroment != ENV_TEST {
		return nil, errors.New("using production database while in testing")
	}

	//ping mongo client
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = root.MongoClient.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		//try to initialize mongo client again if pinging failed
		err := root.InitMongoClient()
		if err != nil {
			//return err if not reachable
			return nil, err
		}
	}

	mongoClient := root.MongoClient
	return mongoClient, nil
}

//GetRedis pings the redis client and returns a *RedisClient object if reachable
func GetRedis() (*RedisClient, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	//check if db config has right enviroment flag
	if util.IsTesting() && root.RedisClient.Config.Enviroment != ENV_TEST {
		return nil, errors.New("using production database while in testing")
	}

	//ping existing client
	client := root.RedisClient.Client
	_, err = client.Ping().Result()
	if err != nil {
		//reinitialize if pinging failed
		err = root.InitRedisClient()
		if err != nil {
			//return err if not reachable
			return nil, err
		}
	}

	return root.RedisClient, nil
}

//GetScoreConfig retrieves the all score parameters or if nil, reinitializes them
func GetScoreConfig() (*ScoreParams, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	if root.ScoreConfig == nil {
		err := root.InitScoreConfig()
		if err != nil {
			return nil, err
		}
	}

	return root.ScoreConfig, nil

}

//ReinitializeScoreConfig forces reinitialization of all score parmeters
func ReinitializeScoreConfig() error {
	root, err := getRoot()
	if err != nil {
		return err
	}

	err = root.InitScoreConfig()
	if err != nil {
		return err
	}

	return nil
}
