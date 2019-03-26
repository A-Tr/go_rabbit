package main

import (
	log"github.com/sirupsen/logrus"
	"net/http"
	"go_rabbit/api"
	"go_rabbit/bus"

)

func init() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors: true,
	})
}

func main() {
	router := api.NewRouter()
	log.Print("Starting server")

	eb := bus.NewBus("RABBIT")
	go eb.ConsumeMessages()

	log.Fatal(http.ListenAndServe(":3001", router))
}

