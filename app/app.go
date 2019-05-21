package app

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

//InitRouter defines the middleware for the chi router and initializes the endpoints
func InitRouter(apiPath string) *chi.Mux {

	router := chi.NewRouter()

	//use default middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	//ENDPOINT /api
	router.Get(apiPath,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API is running!"))
		},
	)

	//Wrapper for specific routes
	InitUserRoutes(router, apiPath)

	return router
}

//HandleCors adds cors support to the router
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
