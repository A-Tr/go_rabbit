package http

import (
	"github.com/sirupsen/logrus"
	"encoding/json"
	rp "go_rabbit/repositories"
	"net/http"

	log "go_rabbit/logger"
)

type Handler struct {
	Repository rp.Repository
}

func (h *Handler) HandleSend(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger(log.LoggerConfig{Component: "Send Message Handler"}, r.Header.Get("request-id"))

	err := h.Repository.PublishMessage("nuevomensajedefinitivo", "SOMEQUEUE", logger)
	if err != nil {
		HandleError(err, http.StatusInternalServerError, w, logger)
		return
	}

	logger.Info("Everything went ok")
}

func HandleLiveness(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger(log.LoggerConfig{Component: "Liveness Handler"}, r.Header.Get("request-id"))

	logger.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger(log.LoggerConfig{Component: "Readiness Handler"}, r.Header.Get("request-id"))

	logger.Print("TODO EN ORDEN")
	w.WriteHeader(200)
}

func HandleError(err error, code int, w http.ResponseWriter, logger *logrus.Entry) {
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
