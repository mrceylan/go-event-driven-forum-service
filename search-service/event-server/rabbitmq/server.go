package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMqServerSettings struct {
	ServerUrl string
}

type RabbitMqServer struct {
	Conn *amqp.Connection
}

func NewRabbitMqServer(settings RabbitMqServerSettings) (*RabbitMqServer, error) {
	defer log.Println("Event server started...")
	connectRabbitMQ, err := amqp.Dial(settings.ServerUrl)
	if err != nil {
		return nil, err
	}
	return &RabbitMqServer{connectRabbitMQ}, nil
}

func (c *RabbitMqServer) Disconnect() error {
	err := c.Conn.Close()
	return err
}
