package main

import (
	"go_rabbit/config"
	controller "go_rabbit/controller/http"
	log "go_rabbit/logger"
	rp "go_rabbit/repositories"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var (
	cfg        config.ServiceConfig
	handlers   controller.Handler
	repository rp.Repository
	app        controller.App
	loggerCfg  log.LoggerConfig
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

	logger.Println(" [*] Waiting for logs. To exit press CTRL+C")
	//Queue should exist before consume
	go func() {
		logger.Fatal(http.ListenAndServe(cfg.Port, appRouter))
	}()

	go func() {
		err := app.BusController.ConsumeMessages(logger)
		if err != nil {
			logger.WithError(err).Fatal("MAIN: Error consuming messages")
		}
	}()

	logger.Println("Ready to produce and consume")

	// graceful stop
	quitSig := make(chan os.Signal)
	signal.Notify(quitSig, syscall.SIGTERM)
	signal.Notify(quitSig, syscall.SIGINT)
	select {
	case sig := <-quitSig:
		logger.Printf("Got %s signal. Handled gracefully.", sig)
		os.Exit(0)
	}
}
