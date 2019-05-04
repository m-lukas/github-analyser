package db

import (
	"context"
	"errors"
	"time"

	"github.com/m-lukas/github-analyser/util"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbRoot *DatabaseRoot

const (
	ENV_PROD = "ENV_PROD"
	ENV_TEST = "ENV_TEST"
)

type DatabaseRoot struct {
	MongoClient *MongoClient
	RedisClient *RedisClient
	ScoreConfig *ScoreParams
}

func Init() error {
	var err error

	dbRoot = &DatabaseRoot{}

	if util.IsTesting() {
		return errors.New("not available in testing enviroment")
	}

	err = dbRoot.initMongoClient()
	if err != nil {
		return err
	}

	err = dbRoot.initRedisClient()
	if err != nil {
		return err
	}

	err = dbRoot.initScoreConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetMongo() (*MongoClient, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	if util.IsTesting() && root.MongoClient.Config.Enviroment != ENV_TEST {
		return nil, errors.New("using production database while in testing")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = root.MongoClient.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		err := root.initMongoClient()
		if err != nil {
			return nil, err
		}
	}

	mongoClient := root.MongoClient
	return mongoClient, nil

}

func GetRedis() (*RedisClient, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	if util.IsTesting() && root.RedisClient.Config.Enviroment != ENV_TEST {
		return nil, errors.New("using production database while in testing")
	}

	client := root.RedisClient.Client

	_, err = client.Ping().Result()
	if err != nil {

		err = root.initRedisClient()
		if err != nil {
			return nil, err
		}

	}

	return root.RedisClient, nil
}

func GetScoreConfig() (*ScoreParams, error) {

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	if root.ScoreConfig == nil {
		err := root.initScoreConfig()
		if err != nil {
			return nil, err
		}
	}

	return root.ScoreConfig, nil

}

func ReinitializeScoreConfig() error {
	root, err := getRoot()
	if err != nil {
		return err
	}

	err = root.initScoreConfig()
	if err != nil {
		return err
	}

	return nil
}
