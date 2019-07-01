package bus

type BusController interface {
	ConsumeMessages([]byte, chan []byte) error
}