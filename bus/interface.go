package bus


type Bus interface {
	InitBus() *BusConfig
	PublishMessage(string) error
}