package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

func main() {
	client, err := InitMongo()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("test")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	fmt.Println(id)
}
