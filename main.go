package main

import (
	"go_rabbit/config"
	controller "go_rabbit/controller/http"
	rp "go_rabbit/repositories"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	cfg        config.ServiceConfig
	handlers   controller.Handler
	repository rp.Repository
	app        controller.App
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := envconfig.Process("local", &cfg)
	if err != nil {
		log.WithError(err).Fatal(err.Error())
	}

	mqProducer, err := rp.InitRabbitRepo(cfg.SrvName, "producer")
	if err != nil {
		log.WithError(err).Fatal("Error connecting to bus")
	}

	mqConsumer, err := rp.InitRabbitRepo(cfg.SrvName, "consumer")
	if err != nil {
		log.WithError(err).Fatal("Error connecting to bus")
	}

	handlers = controller.Handler{
		Repository: mqProducer,
	}

	app = controller.App{
		HTTPController: handlers,
		BusController:  mqConsumer,
	}
}

func main() {

	appRouter := app.CreateRouter()
	log.Print("Starting server " + cfg.SrvName + " on port " + cfg.Port)

	msgChan := make(chan []byte)

	go app.BusController.ConsumeMessages([]byte{}, msgChan)
	go func() {
		for d := range msgChan {
			go app.BusController.ConsumeMessages(d, msgChan)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	log.Fatal(http.ListenAndServe(cfg.Port, appRouter))

}
