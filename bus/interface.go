package bus

import (
	"encoding/json"
	"time"
	"go_rabbit/models"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func InitRabbit() {
	conn := CreateConnection()
	ch := CreateChannel(conn)
	CreateQueue(ch)
	log.Print("Rabbit Ready")

	defer conn.Close()
	defer ch.Close()
}

func CreateConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Error("Couldn't connect to message queue")
	}

	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Couldn't connect to message queue")
	}

	return ch
}

func CreateQueue(ch *amqp.Channel) {
	_, err := ch.QueueDeclare("test", false, false, false, false, nil)
	if err != nil {
		log.Error("Couldn't create queue")
	}
}

func PublishMessage(ch *amqp.Channel) error {
	err := ch.Publish("", "test", false, false, CreateMessage())
	if err != nil {
		log.Error("Couldn't sendMessage")
		return err
	}
	return nil
}

func CreateMessage() amqp.Publishing {
	message := models.PostMessage{
		Title: "Hola",
		Subtitle: "Que tal",
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
