package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/machinebox/graphql"

	"github.com/joho/godotenv"
	"github.com/m-lukas/github-analyser/app"
	"github.com/m-lukas/github-analyser/db"
)

/*
	Server helper structure to setup the HTTP-Server
*/
type Server struct {
	Router     *chi.Mux
	HTTPServer *http.Server
	Config     *ServerConfig
}

/*
	ServerConfig provides a structure for setting and transmitting the Servers config.
*/
type ServerConfig struct {
	Host    string
	Port    int
	APIPath string
}

func defaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Host:    "localhost",
		Port:    8080,
		APIPath: "/api",
	}
}

func configHTTPServer(config *ServerConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(`%s:%d`, config.Host, config.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}
}

func runServer(server *Server) {
	go func() {
		if err := server.HTTPServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println(fmt.Sprintf(`Server has been started on port: %d`, server.Config.Port))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	log.Println("Server was stopped!")
	os.Exit(0)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	testGraphQL()

	server := &Server{
		Config: defaultServerConfig(),
	}

	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	server.Router = app.InitRouter(server.Config.APIPath)

	handler := app.HandleCors(server.Router)
	server.HTTPServer = configHTTPServer(server.Config, handler)

	log.Println("Server has been configurated!")

	runServer(server)

}

type Data struct {
	repositoryOwner RepositoryOwner
}

type RepositoryOwner struct {
	login string
	email string
	bio   string
}

func testGraphQL() {

	client := graphql.NewClient("https://api.github.com/graphql")

	req := graphql.NewRequest(`
		query getUserData($name: String!){
			repositoryOwner(login: $name) {
			login
			... on User {
				email
				bio
			}
			}
		}
	`)

	req.Var("name", "m-lukas")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response := struct {
		data Data
	}{}

	if err := client.Run(ctx, req, &response); err != nil {
		log.Fatal(err)
	}

	fmt.Println("test")
	fmt.Println(response)

}
