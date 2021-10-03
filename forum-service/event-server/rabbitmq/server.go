package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMqServerSettings struct {
	ServerUrl string
}

type RabbitMqServer struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMqServer(settings RabbitMqServerSettings) (*RabbitMqServer, error) {
	defer log.Println("Event server started...")
	connectRabbitMQ, err := amqp.Dial(settings.ServerUrl)
	if err != nil {
		return nil, err
	}
	channel, err := connectRabbitMQ.Channel()
	if err != nil {
		return nil, err
	}

	server := &RabbitMqServer{connectRabbitMQ, channel}

	return server, nil
}

func (srv *RabbitMqServer) Disconnect() error {
	err := srv.Conn.Close()
	return err
}

func (srv *RabbitMqServer) DeclareQueue(queue string) error {
	_, err := srv.Channel.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
