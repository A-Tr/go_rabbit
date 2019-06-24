package bus

import (
	"github.com/pkg/errors"
	"bytes"
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

func InitBus() (*BusConfig, error) {
	config := &BusConfig{}
	config.Conn, config.busErr = amqp.Dial("amqp://guest:guest@localhost:5672")
	if config.busErr != nil {
		err := errors.Wrapf(config.busErr, "Couldn't connect to message queue")
		return nil, err
	}

	config.Ch, config.busErr = config.Conn.Channel()
	if config.busErr != nil {
		err := errors.Wrapf(config.busErr, "Couldn't Create Channel to message queue")
		return nil, err
	}

	return config, nil
}

func createMessage(msg string) amqp.Publishing {
	message := models.PostMessage{
		Title:     msg,
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

func (b *BusConfig) PublishMessage(msg, queue string, log *log.Entry) error {
	b.Q, b.busErr = b.Ch.QueueDeclare(queue, false, false, false, false, nil)
	if b.busErr != nil {
		err := errors.Wrapf(b.busErr, "QUEUE ERROR: %s", queue)
		return err
	}


	err := b.Ch.Publish("", queue, false, false, createMessage(msg))
	if err != nil {
		err = errors.Wrapf(err, "QUEUE ERROR: Couldn't sendMessage %s to queue %s", msg, queue)
		return err
	}
	return nil
}

func (b *BusConfig) ConsumeMessages() ([]byte, error) {
	msgs, err := b.Ch.Consume(
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
		return nil, err
	}

	var buffer bytes.Buffer
	for d := range msgs {
		d.Ack(true)
		buffer.Write(d.Body)
	}
	return buffer.Bytes(), nil
}
