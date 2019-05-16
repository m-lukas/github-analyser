package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/m-lukas/github-analyser/util"
	"github.com/olivere/elastic"
)

//ElasticClient contains the elastic db client, its config and the default database
type ElasticClient struct {
	Client *elastic.Client
	Config *ElasticConfig
}

//ElasticConfig contains config to init elastic db client
type ElasticConfig struct {
	ElasticURI          string
	SniffOpt            bool
	HealthCheckInterval time.Duration
	Enviroment          string
}

//getDefaultConfig return config in dev/prod
func (client *ElasticClient) getDefaultConfig() *ElasticConfig {
	return &ElasticConfig{
		ElasticURI:          getElasticURI(),
		SniffOpt:            false,
		HealthCheckInterval: 10 * time.Second,
		Enviroment:          ENV_PROD,
	}
}

//getTestConfig return config in test
func (client *ElasticClient) getTestConfig() *ElasticConfig {
	return &ElasticConfig{
		ElasticURI:          "localhost:9201",
		SniffOpt:            false,
		HealthCheckInterval: 10 * time.Second,
		Enviroment:          ENV_TEST,
	}
}

//InitElasticClient establishes a connection to the elasticDB instance
func (root *DatabaseRoot) InitElasticClient() error {

	return nil

	elasticClient := &ElasticClient{}
	//assign config according to the enviroment
	if util.IsTesting() {
		elasticClient.Config = elasticClient.getTestConfig()
	} else {
		elasticClient.Config = elasticClient.getDefaultConfig()
	}

	config := elasticClient.Config

	//configurate client
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticURI),
		elastic.SetSniff(config.SniffOpt),
		elastic.SetHealthcheckInterval(config.HealthCheckInterval),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//ping client
	_, _, err = client.Ping(config.ElasticURI).Do(ctx)
	if err != nil {
		//not reachable
		return err
	}

	elasticClient.Client = client
	root.ElasticClient = elasticClient
	log.Println("Initialized elastic client!")

	return nil
}

/*
	Returns configurated URI string for ElasticDB.
*/
func getElasticURI() (uri string) {
	dbHost := os.Getenv("ELASTIC_HOST")
	return dbHost
}
