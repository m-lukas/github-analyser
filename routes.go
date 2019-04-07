package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/m-lukas/github-analyser/db"
)

func checkAPI(w http.ResponseWriter, r *http.Request) {

	db, err := db.Get()
	if err != nil {
		log.Fatal(err)
	}
	collection := db.Collection("test")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID
	fmt.Println(id)

	w.Write([]byte("API is running!"))
}
