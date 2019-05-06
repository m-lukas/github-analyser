package app

import (
	"net/http"
	"strconv"

	"bitbucket.org/timbertom/backend/httputil"
	"bitbucket.org/timbertom/backend/translate"
	"github.com/go-chi/chi"
)

func InitUserRoutes(router *chi.Mux) {

	router.Route("/user", func(r chi.Router) {

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

	router.Route("/score", func(r chi.Router) {

		r.Get("/{score}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")

			scoreInt, e := strconv.Atoi(score)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}

			resp, err := doGetNearestUserByScore(scoreInt)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})
		r.Get("/{score}/next/{entries}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")
			entries := getParam(r, "entries")

			scoreInt, e := strconv.Atoi(score)
			entriesInt, e := strconv.Atoi(entries)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}

			resp, err := doGetNextUsersByScore(scoreInt, entriesInt)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})
		r.Get("/{score}/previous/{entries}", func(w http.ResponseWriter, r *http.Request) {
			score := getParam(r, "score")
			entries := getParam(r, "entries")

			scoreInt, e := strconv.Atoi(score)
			entriesInt, e := strconv.Atoi(entries)
			if e != nil {
				err := httputil.FromTranslationKey(400, translate.MissingParameter)
				err.Write(w)
			}

			resp, err := doGetPreviousUsersByScore(scoreInt, entriesInt)
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

			resp, err := doSearch(query)
			if err != nil {
				err.Write(w)
				return
			}

			httputil.WriteSuccess(w, 200, resp)
		})

	})

}
