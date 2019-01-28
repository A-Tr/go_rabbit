package handlers

import (
	"encoding/json"
	"go_rabbit/models"
	"go_rabbit/bus"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func HandleSend(w http.ResponseWriter, r *http.Request) {
	bus := bus.InitBus()
	err := bus.PublishMessage("Mi nombre es Pepe")
	if err != nil {
		log.Error("Error sending the message to Rabbit")
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Info("Everything went ok")
}

func HandleRead(w http.ResponseWriter, r *http.Request) {
	bus := bus.InitBus()
	resBytes, err := bus.ConsumeMessages()
	if err != nil {
		log.Error("Error sending the message to Rabbit")
		w.WriteHeader(http.StatusInternalServerError)
	}
	var mappedRes models.PostMessage
	json.Unmarshal(resBytes, &mappedRes)
	w.Write(resBytes)
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