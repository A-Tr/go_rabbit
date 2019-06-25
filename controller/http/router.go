package http

import (
	"go_rabbit/controller/bus"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	HTTPController Handler
	BusController  bus.BusController
}

type Middleware func(http.Handler) http.Handler

func (a *App) CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(RequestIdMw)
	router.HandleFunc("/liveness", HandleLiveness).Methods("GET")
	router.HandleFunc("/send", a.HTTPController.HandleSend).Methods("GET")
	// router.HandleFunc("/read", HandleRead)
	return router
}
