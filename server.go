package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/m-lukas/github-analyser/app"
	"github.com/m-lukas/github-analyser/db"

	"github.com/go-chi/chi"

	"github.com/joho/godotenv"
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
		err := server.HTTPServer.ListenAndServe()
		if err != nil {
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

	server := &Server{
		Config: defaultServerConfig(),
	}

	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := db.Get().GetRedis()
	if err != nil {
		log.Fatal(err)
	}
	db.ExampleClient(redisClient)

	/*
		data, err := controller.GetUser("m-lukas")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
	*/

	server.Router = app.InitRouter(server.Config.APIPath)

	handler := app.HandleCors(server.Router)
	server.HTTPServer = configHTTPServer(server.Config, handler)

	log.Println("Server has been configurated!")

	runServer(server)
}
