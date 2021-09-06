package rabbitmq

import (
	"forum-service/models"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMqServerSettings struct {
	ServerUrl      string
	PublishChannel <-chan models.PublishEvent
}

type RabbitMqServer struct {
	Conn           *amqp.Connection
	PublishChannel <-chan models.PublishEvent
}

func NewRabbitMqServer(settings RabbitMqServerSettings) (*RabbitMqServer, error) {
	defer log.Println("Event server started...")
	connectRabbitMQ, err := amqp.Dial(settings.ServerUrl)
	if err != nil {
		return nil, err
	}
	declareQueues(connectRabbitMQ)
	return &RabbitMqServer{connectRabbitMQ, settings.PublishChannel}, nil
}

func (c *RabbitMqServer) Disconnect() error {
	err := c.Conn.Close()
	return err
}

func declareQueues(conn *amqp.Connection) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	_, err = channel.QueueDeclare(
		"deneme", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}
	return nil
}
