package main

import (
	log"github.com/sirupsen/logrus"
	"net/http"
	"go_rabbit/bus"
	"go_rabbit/api"
)


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	router := api.NewRouter()
	log.Print("Starting server")

	bus.InitBus()

	log.Fatal(http.ListenAndServe(":3000", router))
}

