package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

const (
	apiPath = "/api"
	host    = "localhost"
	port    = "8080"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get(apiPath, checkAPI)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "*"},
		Debug:            false,
	})

	handler := cors.Handler(r)

	server := &http.Server{
		Addr:         fmt.Sprintf(`%s:%s`, host, port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	log.Println("Server has been configurated!")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println(fmt.Sprintf(`Server has been started on port: %s`, port))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	log.Println("Server was stopped!")
	os.Exit(0)
	/*
		collection := client.Database("test").Collection("test")

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		if err != nil {
			log.Fatal(err)
		}

		id := res.InsertedID
		fmt.Println(id)
	*/
}
