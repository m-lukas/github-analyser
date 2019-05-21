package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/m-lukas/github-analyser/app"
	"github.com/m-lukas/github-analyser/db"
	"github.com/m-lukas/github-analyser/mailer"
	"github.com/m-lukas/github-analyser/util"

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

	//start mail worker
	go mailer.StartWorker()

	env := os.Getenv("ENV")
	setupFlag := util.ReadBoolFlag("FLAG_DO_SETUP")
	metrixFlag := util.ReadBoolFlag("FLAG_DO_METRIX")

	flag.String("enviroment", env, "Possible values: <dev|prod>")

	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	if setupFlag {
		setupInit()
	}
	if metrixFlag {
		metrixInit()
	}

	server := &Server{
		Config: defaultServerConfig(),
	}

	server.Router = app.InitRouter(server.Config.APIPath)

	handler := app.HandleCors(server.Router)
	server.HTTPServer = configHTTPServer(server.Config, handler)

	log.Println("Server has been configurated!")

	runServer(server)
}
