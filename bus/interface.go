package bus

import (
	"os"
)


type Bus interface {
	SendMessage([]byte) error
}

func InitBus() Bus {
	
	if(os.Getenv("ENV") == "TEST") {
		return &FakeBus{}
	}

	config := &RabbitBus{}
	InitRabbitBus(config)
	
	return config
}