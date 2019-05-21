package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/olivere/elastic/v7"
)

//Insert a new document into the given elastic index
func (elasticClient *ElasticClient) Insert(item interface{}, index string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := elasticClient.Client
	resp, err := client.Index().Index(index).BodyJson(item).Do(ctx)
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}

//Search for all documents in the index where the fields are matching the given search term
func (elasticClient *ElasticClient) Search(term string, index string, fields ...string) ([]json.RawMessage, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := elasticClient.Client
	query := elastic.NewMultiMatchQuery(term, fields...)

	searchResult, err := client.Search().Index(index).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}

	var rawData []json.RawMessage

	for _, hit := range searchResult.Hits.Hits {
		var item json.RawMessage
		err = json.Unmarshal(hit.Source, &item)
		if err != nil {
			return nil, err
		}
		rawData = append(rawData, item)
	}

	return rawData, nil
}

//Update the document with the given id in the index using the provided map
func (elasticClient *ElasticClient) Update(updateData map[string]interface{}, id string, index string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := elasticClient.Client
	resp, err := client.Update().Index(index).Id(id).Doc(updateData).Do(ctx)
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
