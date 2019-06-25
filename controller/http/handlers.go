package http

import (
	"encoding/json"
	rp "go_rabbit/repositories"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Repository rp.Repository
}

func (h *Handler) HandleSend(w http.ResponseWriter, r *http.Request) {
	logger := log.WithField("request-id", r.Header.Get("request-id"))

	err := h.Repository.PublishMessage("nuevomensajedefinitivo", "SOMEQUEUE", logger)
	if err != nil {
		HandleError(err, http.StatusInternalServerError, w, logger)
		return
	}
	logger.Info("Everything went ok")
}

// func HandleRead(w http.ResponseWriter, r *http.Request) {
// 	bus, err := bus.InitBus()

// 	logger := log.WithField("request-id", r.Header.Get("request-id"))

// 	if err != nil {
// 		HandleError(err, http.StatusInternalServerError, w, logger)
// 		return
// 	}

// 	resBytes, err := bus.ConsumeMessages()
// 	if err != nil {
// 		HandleError(err, http.StatusInternalServerError, w, logger)
// 		return
// 	}
// 	var mappedRes models.PostMessage
// 	json.Unmarshal(resBytes, &mappedRes)
// 	w.Write(resBytes)
// 	log.Info("Everything went ok")
// }

func HandleLiveness(w http.ResponseWriter, r *http.Request) {
	log.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	log.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}

func HandleError(err error, code int, w http.ResponseWriter, logger *log.Entry) {
	logger.WithError(err).Error("Error processing request")
	eR := ErrorResponse{code, err.Error()}
	erb, _ := json.Marshal(eR)
	w.WriteHeader(code)
	w.Write(erb)
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
