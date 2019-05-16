package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/olivere/elastic"
)

func (elasticClient *ElasticClient) Insert() {

}

func (elasticClient *ElasticClient) Search(term string, index string, indexType string, fields ...string) ([]json.RawMessage, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := elasticClient.Client
	query := elastic.NewMultiMatchQuery(term, fields...)

	searchResult, err := client.Search().Index(index).Type(indexType).Query(query).Do(ctx)
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

func (client *ElasticClient) Update() {

}
