package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/liveness", HandleLiveness)
	router.HandleFunc("/api/chat", HandleSend).Methods("POST")
	// router.HandleFunc("/read", hd.HandleRead)
	return router
}
