package bus

// Bus interface implements
// all neccesary bus methods
type Bus interface {
	SendMessage([]byte) error
	ConsumeMessages() error
}
