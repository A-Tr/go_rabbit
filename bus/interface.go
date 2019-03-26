package bus

// Bus interface implements
// all neccesary bus methods
type Bus interface {
	SendMessage([]byte) error
	ConsumeMessages() ([]byte, error)
}

// Possible bus types
const (
	RABBIT = "RABBIT"
	TEST = "TEST"
)

// 
func NewBus(busType string) Bus {
	
	if(busType == "TEST") {
		return &FakeBus{}
	}

	config := &RabbitBus{}
	InitRabbitBus(config)
	
	return config
}