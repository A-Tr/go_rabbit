package repositories

import (
	"bytes"
	"encoding/json"
	"go_rabbit/models"
	"time"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitRepo struct {
	Q       amqp.Queue
	Ch      *amqp.Channel
	Conn    *amqp.Connection
	busErr  error
	SrvName string
}

func InitRabbitRepo(srvName string) (*RabbitRepo, error) {
	config := &RabbitRepo{
		SrvName: srvName,
	}

	config.Conn, config.busErr = amqp.Dial("amqp://guest:guest@localhost:5672")
	if config.busErr != nil {
		err := errors.Wrapf(config.busErr, "REPO ERROR")
		return nil, err
	}

	config.Ch, config.busErr = config.Conn.Channel()
	if config.busErr != nil {
		err := errors.Wrapf(config.busErr, "REPO ERROR")
		return nil, err
	}

	return config, nil
}

func (b *RabbitRepo) createMessage(msg string) (amqp.Publishing, error) {
	message := models.PostMessage{
		Message:   msg,
		Publisher: b.SrvName,
		Timestamp: time.Now(),
	}
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Error("Error Marshalling the Message")
		return amqp.Publishing{}, err
	}

	return amqp.Publishing{
		ContentType: "text/plain",
		Body:        msgBytes,
	}, nil
}

func (b *RabbitRepo) PublishMessage(msg, queue string, log *log.Entry) error {
	b.Q, b.busErr = b.Ch.QueueDeclare(queue, false, false, false, false, nil)
	if b.busErr != nil {
		err := errors.Wrapf(b.busErr, "QUEUE ERROR: %s", queue)
		return err
	}

	msgRdy, err := b.createMessage(msg)
	if err != nil {
		err = errors.Wrapf(err, "QUEUE ERROR: Couldn't create message %s for queue %s", msg, queue)
		return err
	}

	err = b.Ch.Publish("", queue, false, false, msgRdy)
	if err != nil {
		err = errors.Wrapf(err, "QUEUE ERROR: Couldn't sendMessage %s to queue %s", msg, queue)
		return err
	}
	return nil
}

func (b *RabbitRepo) ConsumeMessages(c chan []byte) error {
	msgs, err := b.Ch.Consume(
		"SOMEQUEUE", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Error("Error reading messages")
		return err
	}

	var buffer bytes.Buffer
	for d := range msgs {
		d.Ack(true)
		buffer.Write(d.Body)
	}

	defer b.Conn.Close()
	c <- buffer.Bytes()
	return nil
}
