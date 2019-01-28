package bus

import (
	"encoding/json"
	"go_rabbit/models"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type BusConfig struct {
	q      amqp.Queue
	ch     *amqp.Channel
	conn   *amqp.Connection
	busErr	error
}

func InitBus() *BusConfig {
	config := &BusConfig{}
	config.conn, config.busErr = amqp.Dial("amqp://guest:guest@localhost:5672")
	if config.busErr != nil {
		log.Error("Couldn't connect to message queue")
	}

	config.ch, config.busErr = config.conn.Channel()
	if config.busErr != nil {
		log.Error("Couldn't Create Channel to message queue")
	}

	config.q, config.busErr = config.ch.QueueDeclare("test", false, false, false, false, nil)
	if config.busErr != nil {
		log.Error("Couldn't create Queue")
	}
	
	return config
}

func (*BusConfig) CreateMessage() amqp.Publishing {
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

func ConsumeMessages(ch *amqp.Channel) {
	msgs, err := ch.Consume(
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
