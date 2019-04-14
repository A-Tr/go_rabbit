package bus

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitBus struct {
	Q      amqp.Queue
	Ch     *amqp.Channel
	Conn   *amqp.Connection
	busErr error
	done   chan error
}

func ConfigRabbitBus(bus *RabbitBus) *RabbitBus {
	log.Info("Starting bus")

	bus.done = make(chan error)
	bus.Conn, bus.busErr = amqp.Dial("amqp://guest:guest@localhost:5672")
	if bus.busErr != nil {
		log.Error("Couldn't connect to message queue")
	}

	bus.Ch, bus.busErr = bus.Conn.Channel()
	if bus.busErr != nil {
		log.Error("Couldn't Create Channel to message queue")
	}

	return bus
}

func createMessage(msg []byte) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "text/json",
		Body:        msg,
	}
}

func (b *RabbitBus) SendMessage(msg []byte) error {

	log.Info("sending message to Rabbit: ", string(msg))
	b.Q, b.busErr = b.Ch.QueueDeclare("definitivo", false, false, false, false, nil)
	if b.busErr != nil {
		log.Error("Couldn't create Queue or Queue doesn't exist")
	}

	err := b.Ch.Publish("", "definitivo", true, false, createMessage(msg))
	if err != nil {
		log.Error("Couldn't sendMessage: ", err.Error())
		return err
	}

	log.Info("Message succesfully sent")
	return nil
}

func (b *RabbitBus) ConsumeMessages() error {
	b.Q, b.busErr = b.Ch.QueueDeclare("definitivo", false, false, false, false, nil)
	if b.busErr != nil {
		log.Error("Couldn't create Queue or Queue doesn't exist")
	}
	deliveries, err := b.Ch.Consume(
		"definitivo", // name
		"consumer",   // consumerTag,
		false,        // noAck
		false,        // exclusive
		false,        // noLocal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Consume: %s", err)
	}

	go handle(deliveries, b.done)

	return nil
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		HandleMessage(d.Body)
		d.Ack(true)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
