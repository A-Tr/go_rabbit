package api

import (
	"github.com/gorilla/mux"
)

func NewRouter(hd *Handlers) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/liveness", hd.HandleLiveness)
	router.HandleFunc("/api/chat", hd.HandleSend).Methods("POST")
	// router.HandleFunc("/read", hd.HandleRead)
	return router
}
