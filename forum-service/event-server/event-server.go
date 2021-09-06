package eventserver

import (
	"forum-service/event-server/rabbitmq"
	"forum-service/models"
)

type IEventServer interface {
	StartPublisher()
	Disconnect() error
}

type EventServerSettings struct {
	PublishChannel <-chan models.PublishEvent
}

func NewEventServer(settings EventServerSettings) (IEventServer, error) {
	srv, err := rabbitmq.NewRabbitMqServer(
		rabbitmq.RabbitMqServerSettings{
			ServerUrl: "amqp://guest:guest@localhost:5672/",
		})

	if err != nil {
		return nil, err
	}

	return srv, nil
}
