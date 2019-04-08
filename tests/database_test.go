package tests

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/joho/godotenv"
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDatabase(t *testing.T) {

	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/m-lukas/github-analyser/.env"))
	if err != nil {
		t.Errorf("Unable to load .env!")
		t.FailNow()
	}

	err = db.Init()
	if err != nil {
		t.Errorf("Unable to initialize mongo client!")
		t.FailNow()
	}

	db, err := db.Get()
	if err != nil {
		log.Println(err)
		t.Errorf("Unable to establish database connection!")
		t.FailNow()
	}
	collection := db.Collection("testing")

	result := struct {
		ID    primitive.ObjectID `bson:"_id,omitempty"`
		Value string             `bson:"Value"`
	}{}

	objectID, err := primitive.ObjectIDFromHex("5cab24edf3f1a30a16927aeb")
	if err != nil {
		t.Errorf("Could not find test object in database!")
		t.FailNow()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data := collection.FindOne(ctx, bson.M{"_id": objectID})

	err = data.Decode(&result)
	if err != nil {
		t.Errorf("Error while decoding mongo document data!")
		t.FailNow()
	}

	if result.Value != "testvalue" {
		t.Errorf("Wrong return value!")
		t.FailNow()
	}

}
