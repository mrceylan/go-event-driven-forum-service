package eventserver

type IEventServer interface {
	DeclareQueue(queue string) error
	PublishEvent(queue string, eventPayload interface{})
	Disconnect() error
}
