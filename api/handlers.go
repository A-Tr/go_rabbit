package api

import (
	"encoding/json"
	"go_rabbit/bus"
	"go_rabbit/models"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Handlers struct {
	EventBus bus.Bus
}

func (h *Handlers) HandleSend(w http.ResponseWriter, r *http.Request) {

	var msg models.PostMessage
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		log.Error("Error decoding JSON msg")
		return
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Error("Error preparing JSON for sending")
		return
	}

	err = h.EventBus.SendMessage(msgBytes)
	if err != nil {
		log.Error("Error sending json")
	}
}

// func HandleRead(w http.ResponseWriter, r *http.Request) {
// 	bus := bus.InitBus()
// 	resBytes, err := bus.ConsumeMessages()
// 	if err != nil {
// 		log.Error("Error sending the message to Rabbit")
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	var mappedRes models.PostMessage
// 	json.Unmarshal(resBytes, &mappedRes)
// 	w.Write(resBytes)
// 	log.Info("Everything went ok")
// }

func (h *Handlers) HandleLiveness(w http.ResponseWriter, r *http.Request) {
	log.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}

func (h *Handlers) HandleReadiness(w http.ResponseWriter, r *http.Request) {
	log.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}
