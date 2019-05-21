package app

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/m-lukas/github-analyser/errorutil"
	"github.com/m-lukas/github-analyser/httputil"
)

//InitUserRoutes initializes all specific endpoint following the /api path
func InitUserRoutes(router *chi.Mux, basePath string) {

	//User route cluster
	router.Route(basePath+"/user", func(r chi.Router) {

		//ENDPOINT: /user/<login>

		r.Get("/{login}", func(w http.ResponseWriter, r *http.Request) {
			userName := getParam(r, "login")

			resp, err := doGetUser(userName)
			if err != nil {
				err.WriteError(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

		//ENDPOINT: /user/<login>/score

		r.Get("/{login}/score", func(w http.ResponseWriter, r *http.Request) {
			userName := getParam(r, "login")

			resp, err := doGetScore(userName)
			if err != nil {
				err.WriteError(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

	})

	//Score route cluster
	router.Route(basePath+"/score", func(r chi.Router) {

		//ENDPOINT: /score/<score>

		r.Get("/{score}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")

			scoreInt, e := checkAndConvertScore(score)
			if e != nil {
				err := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.ConversionError{Err: e, Param: "score", Value: score}.Error())
				err.WriteError(w)
			}

			collectionName := getUserCollectionName()

			resp, err := doGetNearestUserByScore(scoreInt, collectionName)
			if err != nil {
				err.WriteError(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

		//ENDPOINT: /score/<score>/next/<entries>

		r.Get("/{score}/next/{entries}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")
			entries := getParam(r, "entries")

			scoreInt, e := checkAndConvertScore(score)
			if e != nil {
				err := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.ConversionError{Err: e, Param: "score", Value: score}.Error())
				err.WriteError(w)
			}
			entriesInt, e := strconv.Atoi(entries)
			if e != nil {
				err := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.ConversionError{Err: e, Param: "entries", Value: entries}.Error())
				err.WriteError(w)
			}

			collectionName := getUserCollectionName()

			resp, err := doGetNextUsersByScore(scoreInt, entriesInt, collectionName)
			if err != nil {
				err.WriteError(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

		//ENDPOINT: /score/<score>/previous/<entries>

		r.Get("/{score}/previous/{entries}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")
			entries := getParam(r, "entries")

			scoreInt, e := checkAndConvertScore(score)
			if e != nil {
				err := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.ConversionError{Err: e, Param: "score", Value: score}.Error())
				err.WriteError(w)
			}
			entriesInt, e := strconv.Atoi(entries)
			if e != nil {
				err := httputil.NewError(httputil.INVALID_ARGUMENTS, errorutil.ConversionError{Err: e, Param: "entries", Value: entries}.Error())
				err.WriteError(w)
			}

			collectionName := getUserCollectionName()

			resp, err := doGetPreviousUsersByScore(scoreInt, entriesInt, collectionName)
			if err != nil {
				err.WriteError(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

	})

	//Search route cluster
	router.Route(basePath+"/search", func(r chi.Router) {

		//ENDPOINT: /search

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			query := getQueryParam(r, "search")

			resp, err := doSearch(query)
			if err != nil {
				err.WriteError(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

	})

}
