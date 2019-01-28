package main

import (
	"github.com/gorilla/mux"
	log"github.com/sirupsen/logrus"
	"net/http"
	"go_rabbit/bus"
)


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	router := mux.NewRouter()
	log.Print("Starting server")
	router.HandleFunc("/liveness", handleLiveness)
	router.HandleFunc("/send", handleSend)

	bus.InitRabbit()

	log.Fatal(http.ListenAndServe(":3000", router))
}

func handleLiveness (w http.ResponseWriter, r *http.Request) {
	log.Print("TODO EN ORDEN")
	w.WriteHeader(200)
} 

func handleSend (w http.ResponseWriter, r *http.Request) {
	conn := bus.CreateConnection()
	ch := bus.CreateChannel(conn)
	err := bus.PublishMessage(ch)
	if err != nil {
		log.Error("Error sending the message to Rabbit")
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Info("Everything went ok")
}

