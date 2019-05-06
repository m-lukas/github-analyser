package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func getParam(r *http.Request, name string) string {
	param := chi.URLParam(r, name)
	return param
}
