package app

import (
	"fmt"

	"github.com/m-lukas/github-analyser/controller"
	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/errorutil"
	"github.com/m-lukas/github-analyser/httputil"
	"github.com/m-lukas/github-analyser/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Response struct with main scores
type scoreResponse struct {
	ActivityScore   float64
	PopularityScore float64
}

//Wrapper struct for aggregations regarding ActivityScore
type ActivityAggregationResponse struct {
	Data []*ActivityAggregation
}

//Struct for unmarshaling aggregation result
type ActivityAggregation struct {
	ID            primitive.ObjectID `bson:"_id"`
	Login         string             `bson:"login"`
	ActivityScore float64            `bson:"activity_score"`
	Difference    float64            `bson:"difference"`
}

//Wrapper struct for aggregations with full user data response
type UserSearchAggregationResponse struct {
	Data []*db.User
}

//doGetUser uses the controller functions to fetch and/or get the user by login
func doGetUser(userName string) (*db.User, *httputil.ErrorResponse) {

	if userName == "" {
		resp := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.NullOrEmptyError{Param: "userName"}.Error())
		return nil, resp
	}

	user, err := controller.GetUser(userName)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to get user: %s, error: %s", user.Login, err.Error()))
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return user, nil

}

//doGetScore gets the user data and returns the main score fields
func doGetScore(userName string) (*scoreResponse, *httputil.ErrorResponse) {

	if userName == "" {
		resp := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.NullOrEmptyError{Param: "userName"}.Error())
		return nil, resp
	}

	user, err := controller.GetUser(userName)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to get user: %s, error: %s", user.Login, err.Error()))
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	//convert to selective response
	resp := &scoreResponse{
		ActivityScore:   user.ActivityScore,
		PopularityScore: user.PopularityScore,
	}

	return resp, nil
}

//doGetNearestUserByScore returns the login and score of the user who's ActivityScore is nearest to the given score
func doGetNearestUserByScore(score int, collectionName string) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}

	//building aggregation pipeline
	pipeline.Add(bson.D{{"$project", bson.D{{"login", 1}, {"activity_score", 1}, {"difference", bson.D{{"$abs", bson.D{{"$subtract", bson.A{score, "$activity_score"}}}}}}}}})
	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$gt", 0}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"difference", 1}}}})
	pipeline.Add(bson.D{{"$limit", 1}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to run pipeline: %s, error: %s", "NearestByScore", err.Error()))
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return &result, nil
}

//doGetNextUsersByScore returns a list of n (=entries) users whose score is bigger than the given one in ascending order
func doGetNextUsersByScore(score int, entries int, collectionName string) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}

	//building the pipeline
	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$gt", score}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"activity_score", 1}}}})
	pipeline.Add(bson.D{{"$limit", entries}})
	pipeline.Add(bson.D{{"$project", bson.D{{"activity_score", 1}, {"login", 1}}}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to run pipeline: %s, error: %s", "NextUserByScore", err.Error()))
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return &result, nil
}

//doGetPreviousUsersByScore returns a list of n (=entries) users whose score is lower than the given one in descending order
func doGetPreviousUsersByScore(score int, entries int, collectionName string) (*ActivityAggregationResponse, *httputil.ErrorResponse) {

	pipeline := db.Pipeline{}

	//building the pipeline
	pipeline.Add(bson.D{{"$match", bson.D{{"activity_score", bson.D{{"$lt", score}}}}}})
	pipeline.Add(bson.D{{"$sort", bson.D{{"activity_score", -1}}}})
	pipeline.Add(bson.D{{"$limit", entries}})
	pipeline.Add(bson.D{{"$project", bson.D{{"activity_score", 1}, {"login", 1}}}})

	var result ActivityAggregationResponse

	err := pipeline.Run(&result, collectionName)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to run pipeline: %s, error: %s", "PreviousUserByScore", err.Error()))
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return &result, nil
}

//doSearch searches multiple fields in the elastic index for the provided search term/query
func doSearch(query string) ([]*db.ElasticUser, *httputil.ErrorResponse) {

	results, err := controller.SearchUser(query)
	if err != nil {
		logger.Error(fmt.Sprintf("(Elastic) Failed to search: '%s', error: %s", query, err.Error()))
		return nil, httputil.NewError(httputil.SERVER_ERROR, errorutil.InternalServerError)
	}

	return results, nil
}
