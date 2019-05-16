package db

import (
	"context"
	"errors"
	"fmt"
)

const (
	user_index = "user_index"
)

func (client *ElasticClient) initIndexes() {

	client.Indexes = make(map[string]string, 0)
	client.Indexes[user_index] = userMapping

}

func (elasticClient *ElasticClient) checkIndexes() error {

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

const userMapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"records":{
			"properties":{
				"login":{
					"type":"keyword"
				},
				"firstname":{
					"type":"keyword"
				},
				"lastname":{
					"type":"keyword"
				},
				"bio":{
					"type":"keyword"
				}
			}
		}
	}
}`
