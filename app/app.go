package app

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func InitRouter(apiPath string) *chi.Mux {

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get(apiPath,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API is running!"))
		},
	)

	InitUserRoutes(router)

	return router
}

func HandleCors(h http.Handler) http.Handler {

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "*"},
		Debug:            false,
	})
	handler := cors.Handler(h)
	return handler

}
