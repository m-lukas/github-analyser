package app

import (
	"fmt"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"bitbucket.org/timbertom/backend/httputil"
	"bitbucket.org/timbertom/backend/translate"
)

type scoreResponse struct {
	ActivityScore   float64
	PopularityScore float64
}

type ActivityAggregationResponse struct {
	Data []*ActivityAggregation
}

type ActivityAggregation struct {
	ID            primitive.ObjectID `bson:"_id"`
	Login         string             `bson:"login"`
	ActivityScore float64            `bson:"activity_score"`
	Difference    float64            `bson:"difference"`
}

type UserSearchAggregationResponse struct {
	Data []*db.User
}

func doGetScore(userName string) (*scoreResponse, *httputil.ErrorResponse) {

	if userName == "" {
		return nil, httputil.FromTranslationKey(400, translate.MissingParameter)
	}

	user, err := controller.GetUser(userName)
	if err != nil {
		return nil, httputil.FromTranslationKey(500, translate.ServerError)
	}

	resp := &scoreResponse{
		ActivityScore:   user.ActivityScore,
		PopularityScore: user.PopularityScore,
	}

	return resp, nil
}

func doGetUser(userName string) (*db.User, *httputil.ErrorResponse) {

	if userName == "" {
		return nil, httputil.FromTranslationKey(400, translate.MissingParameter)
	}

	user, err := controller.GetUser(userName)
	if err != nil {
		return nil, httputil.FromTranslationKey(500, translate.ServerError)
	}

	return user, nil

}

func doGetNearestUserByScore(score int) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}
	collectionName := "users"

	pipeline.Add(bson.D{{"$project", bson.D{{"login", 1}, {"activity_score", 1}, {"difference", bson.D{{"$abs", bson.D{{"$subtract", bson.A{score, "$activity_score"}}}}}}}}})
	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$gt", 0}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"difference", 1}}}})
	pipeline.Add(bson.D{{"$limit", 1}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		fmt.Println(err)
		return nil, httputil.FromTranslationKey(500, translate.ServerError)
	}

	return &result, nil
}

func doGetNextUsersByScore(score int, entries int) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}
	collectionName := "users"

	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$gt", score}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"activity_score", 1}}}})
	pipeline.Add(bson.D{{"$limit", entries}})
	pipeline.Add(bson.D{{"$project", bson.D{{"activity_score", 1}, {"login", 1}}}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		fmt.Println(err)
		return nil, httputil.FromTranslationKey(500, translate.ServerError)
	}

	return &result, nil
}

func doGetPreviousUsersByScore(score int, entries int) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}
	collectionName := "users"

	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$lt", score}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"activity_score", -1}}}})
	pipeline.Add(bson.D{{"$limit", entries}})
	pipeline.Add(bson.D{{"$project", bson.D{{"activity_score", 1}, {"login", 1}}}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		fmt.Println(err)
		return nil, httputil.FromTranslationKey(500, translate.ServerError)
	}

	return &result, nil
}

func doSearch(query string) (*UserSearchAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}
	collectionName := "users"

	pipeline.Add(bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", query}}}}}})

	var result UserSearchAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		fmt.Println(err)
		return nil, httputil.FromTranslationKey(500, translate.ServerError)
	}

	return &result, nil
}
