package rabbitmq

type IRMQService interface {
	PublishMessage(messageBody []byte) error
	ConsumeMessage(stopChan chan bool, errorChan chan error)
}
