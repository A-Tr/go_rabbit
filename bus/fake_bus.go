package bus

import (
	log "github.com/sirupsen/logrus"
)

type FakeBus struct {}

func (b *FakeBus) SendMessage(msg []byte) error {
	log.Print("TEST BUS: Message received: ", string(msg))
	return nil
}
