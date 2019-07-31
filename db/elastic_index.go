package db

import (
	"context"
	"errors"
	"fmt"
)

const (
	user_index      = "users"
	user_index_test = "users_test"
)

func (client *ElasticClient) initIndexes() {

	client.Indexes = make(map[string]string, 0)
	client.Indexes[user_index] = userMapping
	client.Indexes[user_index_test] = userMapping

}

func (elasticClient *ElasticClient) CheckIndexes() error {

	client := elasticClient.Client

	ctx := context.Background()
	for name, mapping := range elasticClient.Indexes {

		exists, err := client.IndexExists(name).Do(ctx)
		if err != nil {
			return err
		}

		if !exists {

			createIndex, err := client.CreateIndex(name).BodyString(mapping).Do(ctx)
			if err != nil {
				return err
			}
			if !createIndex.Acknowledged {
				return errors.New(fmt.Sprintf("index creation failed for: %s", name))
			}

		}
	}

	return nil

}

func (elasticClient *ElasticClient) DeleteIndex(name string) error {

	client := elasticClient.Client

	ctx := context.Background()

	exists, err := client.IndexExists(name).Do(ctx)
	if err != nil {
		return err
	}

	if exists {

		deleteIndex, err := elasticClient.Client.DeleteIndex(name).Do(ctx)
		if err != nil {
			return err
		}
		if !deleteIndex.Acknowledged {
			return errors.New(fmt.Sprintf("index deletion failed for: %s", name))
		}

	}

	return nil

}

const userMapping = `
{
	"mappings":{
		"properties":{
			"login":{
				"type":"keyword"
			},
			"email":{
				"type":"keyword"
			},
			"name":{
				"type":"text"
			},
			"bio":{
				"type":"text"
			}
		}
	}
}`
