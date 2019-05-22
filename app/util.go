package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/m-lukas/github-analyser/util"
)

//getParam retrieves the url param from chi
func getParam(r *http.Request, name string) string {
	param := chi.URLParam(r, name)
	return param
}

//getQueryParam retrieves the query param from chi
func getQueryParam(r *http.Request, name string) string {
	query := r.URL.Query()
	param := query.Get(name)
	return param
}

//getUserCollectionName get default user collection for the current enviroment
func getUserCollectionName() string {
	collectionName := "users"
	if util.IsTesting() {
		collectionName = "test_getuser"
	}
	return collectionName
}

//checkAndConvertScore converts score parameter and checks range
func checkAndConvertScore(score string) (int, error) {
	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		return 0, err
	}
	if scoreInt < 0 || scoreInt > 100 {
		return 0, errors.New("Score must be in range 0-100!")
	}

	return scoreInt, nil
}
