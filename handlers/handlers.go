package handlers

import (
	"encoding/json"
	"go_rabbit/bus"
	"go_rabbit/models"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HandleSend(w http.ResponseWriter, r *http.Request) {

	bus := bus.InitBus()
	var msg models.PostMessage
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		log.Error("Error decoding JSON msg")
	}

	for i := 0; i < 10000; i++ {
		msg.Message = "quease5"
		err = bus.PublishMessage(msg)
		log.Info("MESSAGE ", msg, " sent")
		if err != nil {
			log.Error("Error sending the message to Rabbit")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}	
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
