package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func checkAPI(w http.ResponseWriter, r *http.Request) {
	db, err := db.Get()
	if err != nil {
		log.Fatal("Unable to establish database connection!")
	}
	collection := db.Collection("testing")

	result := struct {
		ID    primitive.ObjectID `bson:"_id,omitempty"`
		Value string             `bson:"Value"`
	}{}

	objectID, err := primitive.ObjectIDFromHex("5cab24edf3f1a30a16927aeb")
	if err != nil {
		log.Fatal("Could not convert test object id!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data := collection.FindOne(ctx, bson.M{"_id": objectID})

	err = data.Decode(&result)
	if err != nil {
		log.Fatal("Error while decoding mongo document data!")
	}
	w.Write([]byte(result.Value))
}
