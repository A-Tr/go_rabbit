package bus

import (
	"encoding/json"
	"go_rabbit/models"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type BusConfig struct {
	Q      amqp.Queue
	Ch     *amqp.Channel
	Conn   *amqp.Connection
	busErr error
}

func InitBus() *BusConfig {
	config := &BusConfig{}
	config.Conn, config.busErr = amqp.Dial("amqp://guest:guest@localhost:5672")
	if config.busErr != nil {
		log.Error("Couldn't connect to message queue")
	}

	config.Ch, config.busErr = config.Conn.Channel()
	if config.busErr != nil {
		log.Error("Couldn't Create Channel to message queue")
	}

	config.Q, config.busErr = config.Ch.QueueDeclare("test", false, false, false, false, nil)
	if config.busErr != nil {
		log.Error("Couldn't create Queue")
	}

	return config
}

func createMessage() amqp.Publishing {
	message := models.PostMessage{
		Title:     "Hola",
		Subtitle:  "Que tal",
		Timestamp: time.Now(),
	}
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Error("Error Marshalling the Message")
	}

	return amqp.Publishing{
		ContentType: "text/plain",
		Body:        msgBytes,
	}
}

func (*BusConfig) PublishMessage(Ch *amqp.Channel) error {
	err := Ch.Publish("", "test", false, false, createMessage())
	if err != nil {
		log.Error("Couldn't sendMessage")
		return err
	}
	return nil
}

func ConsumeMessages(Ch *amqp.Channel) {
	msgs, err := Ch.Consume(
		"test", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Error("Error reading messages")
	}
	for d := range msgs {
		log.Print(d)
	}
}
