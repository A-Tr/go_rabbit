package bus

import (
	log "github.com/sirupsen/logrus"
)

func HandleMessage(msg []byte) error {
	log.Print("Message Received: ", string(msg))
	return nil
}
