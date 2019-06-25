package bus

type BusController interface {
	ConsumeMessages(c chan []byte) error
}