package main

import (
	"go_rabbit/api"
	"go_rabbit/bus"
	"go_rabbit/handlers"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	eb, err := bus.InitBus()
	if err != nil {
		log.WithError(err).Fatal("Error connecting to bus")
	}
	handler := handlers.Handler{
		Bus: eb,
	}
	app := api.App{
		Handlers: handler,
	}

	appRouter := app.CreateRouter()
	log.Print("Starting server on port 3000")

	log.Fatal(http.ListenAndServe(":3000", appRouter))
}
