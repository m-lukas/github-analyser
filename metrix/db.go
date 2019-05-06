package metrix

import (
	"fmt"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
)

func save(pairs map[string]interface{}, field string) error {

	redisClient, err := db.GetRedis()
	if err != nil {
		return err
	}

	for key, value := range pairs {

		err = redisClient.HashInsert(key, field, value)
		if err != nil {
			return err
		}

	}

	return nil
}

func updateExisting() error {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return err
	}
	collectionName := "users"

	userArray, err := mongoClient.FindAllUsers(collectionName)
	if err != nil {
		return err
	}

	scoreConfig, err := db.GetScoreConfig()
	if err != nil {
		return err
	}

	for _, user := range userArray {

		user.Scores = controller.CalcScores(user, scoreConfig)

		user.ActivityScore = controller.CalcActivityScore(user.Scores, scoreConfig)
		user.PopularityScore = controller.CalcPopularityScore(user.Scores, scoreConfig)

		filter := bson.D{{"login", user.Login}}
		err = mongoClient.UpdateAll(filter, user, collectionName)
		if err != nil {
			return err
		}

		fmt.Printf("%s Updated score for user: %s\n", prefix, user.Login)

	}

	return nil

}

func getUserFromCache(login string) (*db.User, error) {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return nil, err
	}
	collectionName := "users"

	dbUser, err := mongoClient.FindUser(login, collectionName)
	if err != nil {
		return nil, err
	}

	return dbUser, nil

}
