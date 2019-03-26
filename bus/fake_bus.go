package bus

import (
	log "github.com/sirupsen/logrus"
)

type FakeBus struct {}

func (b *FakeBus) SendMessage(msg []byte) error {
	log.Print("TEST BUS: Message received: ", string(msg))
	return nil
}

func (b *FakeBus) ConsumeMessages() ([]byte, error) {
	log.Print("TEST BUS: Message consumed: ")
	return nil, nil
}
