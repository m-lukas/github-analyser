package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func getParam(r *http.Request, name string) string {
	param := chi.URLParam(r, name)
	return param
}

func getQueryParam(r *http.Request, name string) string {
	query := r.URL.Query()
	param := query.Get(name)
	return param
}
