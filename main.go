package main

import (
	"go_rabbit/api"
	"go_rabbit/config"
	rp "go_rabbit/repositories"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	cfg          config.ServiceConfig
	handlers     api.Handler
	repositories rp.Repository
	app          api.App
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := envconfig.Process("local", &cfg)
	if err != nil {
		log.WithError(err).Fatal(err.Error())
	}

	repositories, err = rp.InitRabbitRepo(cfg.SrvName)
	if err != nil {
		log.WithError(err).Fatal("Error connecting to bus")
	}

	handlers = api.Handler{
		Repository: repositories,
	}

	app = api.App{
		Handlers: handlers,
	}
}

func main() {

	appRouter := app.CreateRouter()
	log.Print("Starting server " + cfg.SrvName + " on port " + cfg.Port)

	log.Fatal(http.ListenAndServe(cfg.Port, appRouter))
}
