package app

import (
	"log"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/errorutil"
	"github.com/m-lukas/github-analyser/httputil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func doGetUser(userName string) (*db.User, *httputil.ErrorResponse) {

	if userName == "" {
		resp := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.NullOrEmptyError{Param: "userName"}.Error())
		return nil, resp
	}

	user, err := controller.GetUser(userName)
	if err != nil {
		log.Println(err)
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return user, nil

}

func doGetScore(userName string) (*scoreResponse, *httputil.ErrorResponse) {

	if userName == "" {
		resp := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.NullOrEmptyError{Param: "userName"}.Error())
		return nil, resp
	}

	user, err := controller.GetUser(userName)
	if err != nil {
		log.Println(err)
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	resp := &scoreResponse{
		ActivityScore:   user.ActivityScore,
		PopularityScore: user.PopularityScore,
	}

	return resp, nil
}

func doGetNearestUserByScore(score int, collectionName string) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}

	pipeline.Add(bson.D{{"$project", bson.D{{"login", 1}, {"activity_score", 1}, {"difference", bson.D{{"$abs", bson.D{{"$subtract", bson.A{score, "$activity_score"}}}}}}}}})
	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$gt", 0}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"difference", 1}}}})
	pipeline.Add(bson.D{{"$limit", 1}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		log.Println(err)
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return &result, nil
}

func doGetNextUsersByScore(score int, entries int, collectionName string) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}

	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$gt", score}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"activity_score", 1}}}})
	pipeline.Add(bson.D{{"$limit", entries}})
	pipeline.Add(bson.D{{"$project", bson.D{{"activity_score", 1}, {"login", 1}}}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		log.Println(err)
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return &result, nil
}

func doGetPreviousUsersByScore(score int, entries int, collectionName string) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}

	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$lt", score}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"activity_score", -1}}}})
	pipeline.Add(bson.D{{"$limit", entries}})
	pipeline.Add(bson.D{{"$project", bson.D{{"activity_score", 1}, {"login", 1}}}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		log.Println(err)
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return &result, nil
}

func doSearch(query string) ([]*db.ElasticUser, *httputil.ErrorResponse) {

	results, err := controller.SearchUser(query)
	if err != nil {
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return results, nil
}
