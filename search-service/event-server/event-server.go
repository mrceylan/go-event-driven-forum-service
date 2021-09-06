package eventserver

import (
	"search-service/event-server/rabbitmq"
	"search-service/services/search"
)

type IEventServer interface {
	StartConsumer()
	Disconnect() error
}

type EventServerSettings struct {
	SearchService search.ISearchService
}

func NewEventServer(settings EventServerSettings) (IEventServer, error) {
	srv, err := rabbitmq.NewRabbitMqServer(
		rabbitmq.RabbitMqServerSettings{
			ServerUrl:     "amqp://guest:guest@localhost:5672/",
			SearchService: settings.SearchService,
		})

	if err != nil {
		return nil, err
	}

	return srv, nil
}
