package rabbitmq

import (
	"log"
	"search-service/services/search"

	"github.com/streadway/amqp"
)

type RabbitMqServerSettings struct {
	ServerUrl     string
	SearchService search.ISearchService
}

type RabbitMqServer struct {
	Conn          *amqp.Connection
	SearchService search.ISearchService
}

func NewRabbitMqServer(settings RabbitMqServerSettings) (*RabbitMqServer, error) {
	defer log.Println("Event server started...")
	connectRabbitMQ, err := amqp.Dial(settings.ServerUrl)
	if err != nil {
		return nil, err
	}
	return &RabbitMqServer{connectRabbitMQ, settings.SearchService}, nil
}

func (c *RabbitMqServer) Disconnect() error {
	err := c.Conn.Close()
	return err
}
