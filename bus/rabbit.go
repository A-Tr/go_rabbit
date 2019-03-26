package bus

import (
	"bytes"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitBus struct {
	Q      amqp.Queue
	Ch     *amqp.Channel
	Conn   *amqp.Connection
	busErr error
}

func InitRabbitBus(config *RabbitBus) *RabbitBus {
	log.Info("Starting bus")
	config.Conn, config.busErr = amqp.Dial("amqp://guest:guest@localhost:5672")
	if config.busErr != nil {
		log.Error("Couldn't connect to message queue")
	}

	config.Ch, config.busErr = config.Conn.Channel()
	if config.busErr != nil {
		log.Error("Couldn't Create Channel to message queue")
	}

	config.Q, config.busErr = config.Ch.QueueDeclare("chatroom", false, false, false, false, nil)
	if config.busErr != nil {
		log.Error("Couldn't create Queue")
	}

	return config
}
func createMessage(msg []byte) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "text/json",
		Body:        msg,
	}
}

func (b *RabbitBus) SendMessage(msg []byte) error {

	log.Info("sending message to Rabbit: ", string(msg))
	err := b.Ch.Publish("", "chatroom", false, false, createMessage(msg))
	if err != nil {
		log.Error("Couldn't sendMessage: ", err.Error())
		return err
	}

	log.Info("Message succesfully sent")
	return nil
}

func (b *RabbitBus) ConsumeMessages() ([]byte, error) {
	msgs, err := b.Ch.Consume(
		"chatroom", // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Error("Error reading messages")
		return nil, err
	}

	var buffer bytes.Buffer
	for d := range msgs {
		d.Ack(true)
		// buffer.Write(d.Body)
	}
	return buffer.Bytes(), nil
}
