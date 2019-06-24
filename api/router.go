package api

import (
	"net/http"
	hd "go_rabbit/handlers"

	"github.com/gorilla/mux"
)

type App struct {
	Handlers hd.Handler
}

type Middleware func(http.Handler) http.Handler

func (a *App) CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(RequestIdMw)
	router.HandleFunc("/liveness", hd.HandleLiveness).Methods("GET")
	router.HandleFunc("/send", a.Handlers.HandleSend).Methods("GET")
	router.HandleFunc("/read", hd.HandleRead)
	return router
}
