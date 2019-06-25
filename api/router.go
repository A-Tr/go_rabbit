package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Handlers Handler
}

type Middleware func(http.Handler) http.Handler

func (a *App) CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(RequestIdMw)
	router.HandleFunc("/liveness", HandleLiveness).Methods("GET")
	router.HandleFunc("/send", a.Handlers.HandleSend).Methods("GET")
	// router.HandleFunc("/read", HandleRead)
	return router
}
