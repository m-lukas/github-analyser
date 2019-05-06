package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type Pipeline struct {
	Stages []bson.D
}

func (p *Pipeline) Add(stage bson.D) {
	p.Stages = append(p.Stages, stage)
}

func (p *Pipeline) Run(result interface{}, collectionName string) error {

	mongoClient, err := GetMongo()
	if err != nil {
		return err
	}
	collection := mongoClient.Database.Collection(collectionName)

	pipe := mongo.Pipeline{}
	for _, stage := range p.Stages {
		pipe = append(pipe, stage)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Aggregate(ctx, pipe)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var entrySlice []bson.D

	for cursor.Next(ctx) {
		entry := bson.D{}
		err := cursor.Decode(&entry)
		if err != nil {
			return err
		}
		entrySlice = append(entrySlice, entry)
	}

	rawData := struct{ Data interface{} }{Data: entrySlice}

	bytes, err := bson.Marshal(rawData)
	if err != nil {
		return err
	}
	bson.Unmarshal(bytes, result)

	return nil
}
