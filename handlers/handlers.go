package handlers

import (
	"go_rabbit/bus"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HandleSend(w http.ResponseWriter, r *http.Request) {
	prodBus := bus.InitBus()
	err := prodBus.PublishMessage(prodBus.Ch)
	if err != nil {
		log.Error("Error sending the message to Rabbit")
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Info("Everything went ok")
}

func HandleLiveness(w http.ResponseWriter, r *http.Request) {
	log.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	log.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}