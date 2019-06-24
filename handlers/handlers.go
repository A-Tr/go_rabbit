package handlers

import (
	"github.com/satori/go.uuid"
	"encoding/json"
	"go_rabbit/models"
	"go_rabbit/bus"
	"net/http"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Bus *bus.BusConfig
}

func (h *Handler) HandleSend(w http.ResponseWriter, r *http.Request) {
	reqId, _ := uuid.NewV4()
	logger := log.WithField("request-id", reqId)

	err :=  h.Bus.PublishMessage("Mi nombre es Pepe", logger)
	if err != nil {
		logger.Error("Error sending the message to Rabbit")
		w.WriteHeader(http.StatusInternalServerError)
	}
	logger.Info("Everything went ok")
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