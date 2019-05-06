package app

import (
	"net/http"

	"bitbucket.org/timbertom/backend/httputil"
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

}
