package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

//Pipeline is a help structure to combine many stages to a MongoDB aggregation pipeline.
type Pipeline struct {
	Stages []bson.D
}

//Add adds a new bson.D Map as stage to the pipeline.
func (p *Pipeline) Add(stage bson.D) {
	p.Stages = append(p.Stages, stage)
}

//Run executes the stages of the Pipeline interface as aggregation function
//and decodes the result into the given interface.
//Given interface must be an struct with format: {Data: []<type>}
func (p *Pipeline) Run(result interface{}, collectionName string) error {

	//configurate mongo collection
	mongoClient, err := GetMongo()
	if err != nil {
		return err
	}
	collection := mongoClient.Database.Collection(collectionName)

	//add stages of Pipeline helper struct to mongo's pipeline struct
	pipe := mongo.Pipeline{}
	for _, stage := range p.Stages {
		pipe = append(pipe, stage)
	}

	//context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//execute aggregation on collection, retrieve cursor to current entry
	cursor, err := collection.Aggregate(ctx, pipe)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	//helper slice for merging all entries of the cursor into one slice
	var entrySlice []bson.D

	//loop throgh cursor, decode entry into bson.D Map, add to helper slice
	for cursor.Next(ctx) {
		entry := bson.D{}
		err := cursor.Decode(&entry)
		if err != nil {
			return err
		}
		entrySlice = append(entrySlice, entry)
	}

	//wrapper struct for helper slice (array -> struct)
	rawData := struct{ Data interface{} }{Data: entrySlice}

	//marshal data into []byte, unmarshal into helper struct
	bytes, err := bson.Marshal(rawData)
	if err != nil {
		return err
	}
	bson.Unmarshal(bytes, result)

	return nil
}
