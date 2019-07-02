package main

import (
	"github.com/sirupsen/logrus"
	"go_rabbit/config"
	controller "go_rabbit/controller/http"
	rp "go_rabbit/repositories"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	log "go_rabbit/logger"
)

var (
	cfg        config.ServiceConfig
	handlers   controller.Handler
	repository rp.Repository
	app        controller.App
	loggerCfg     log.LoggerConfig
)

func init() {
	log.Init(os.Stdout, logrus.InfoLevel)

	err := envconfig.Process("local", &cfg)
	if err != nil {
		logrus.WithError(err).Fatal(err.Error())
	}

	mqProducer, err := rp.InitRabbitRepo(cfg.SrvName, "producer")
	if err != nil {
		logrus.WithError(err).Fatal("Error connecting to bus")
	}

	mqConsumer, err := rp.InitRabbitRepo(cfg.SrvName, "consumer")
	if err != nil {
		logrus.WithError(err).Fatal("Error connecting to bus")
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

	logger := log.NewLogger(loggerCfg, "CONSUMER_LOGGER")
	appRouter := app.CreateRouter()
	logger.Print("Starting server " + cfg.SrvName + " on port " + cfg.Port)

	go func() {
		err :=  app.BusController.ConsumeMessages(logger)
		if err != nil {
			logger.WithError(err).Fatal("MAIN: Error consuming messages")
		}
	}()

	logger.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	logger.Fatal(http.ListenAndServe(cfg.Port, appRouter))

}
