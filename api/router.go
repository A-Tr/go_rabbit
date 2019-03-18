package api

import (
	hd "go_rabbit/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/liveness", hd.HandleLiveness)
	router.HandleFunc("/api/chat", hd.HandleSend).Methods("POST")
	router.HandleFunc("/read", hd.HandleRead)
	return router
}
