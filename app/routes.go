package app

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/m-lukas/github-analyser/httputil"
	"github.com/m-lukas/github-analyser/translate"
)

func InitUserRoutes(router *chi.Mux, basePath string) {

	router.Route(basePath+"/user", func(r chi.Router) {

		r.Get("/{login}", func(w http.ResponseWriter, r *http.Request) {
			userName := getParam(r, "login")

			resp, err := doGetUser(userName)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})
		r.Get("/{login}/score", func(w http.ResponseWriter, r *http.Request) {
			userName := getParam(r, "login")

			resp, err := doGetScore(userName)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

	})

	router.Route(basePath+"/score", func(r chi.Router) {

		r.Get("/{score}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")

			scoreInt, e := checkAndConvertScore(score)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}

			collectionName := getUserCollectionName()

			resp, err := doGetNearestUserByScore(scoreInt, collectionName)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})
		r.Get("/{score}/next/{entries}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")
			entries := getParam(r, "entries")

			scoreInt, e := checkAndConvertScore(score)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}
			entriesInt, e := strconv.Atoi(entries)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}

			collectionName := getUserCollectionName()

			resp, err := doGetNextUsersByScore(scoreInt, entriesInt, collectionName)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})
		r.Get("/{score}/previous/{entries}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")
			entries := getParam(r, "entries")

			scoreInt, e := checkAndConvertScore(score)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}
			entriesInt, e := strconv.Atoi(entries)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}

			collectionName := getUserCollectionName()

			resp, err := doGetPreviousUsersByScore(scoreInt, entriesInt, collectionName)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

	})

	router.Route("/search", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			query := getQueryParam(r, "search")

			collectionName := getUserCollectionName()

			resp, err := doSearch(query, collectionName)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

	})

}
