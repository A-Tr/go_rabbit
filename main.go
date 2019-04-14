package main

import (
	"go_rabbit/api"
	"go_rabbit/bus"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func init() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func main() {
	eB := &bus.RabbitBus{}
	eB = bus.ConfigRabbitBus(eB)

	hd := &api.Handlers{
		EventBus: eB,
	}
	router := api.NewRouter(hd)
	log.Print("Starting server")

	eb := bus.NewBus("RABBIT")
	go eb.ConsumeMessages()

	log.Fatal(http.ListenAndServe(":3001", router))
}
