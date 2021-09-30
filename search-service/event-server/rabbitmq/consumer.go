package rabbitmq

import (
	"context"
	"log"
)

func (srv *RabbitMqServer) StartConsumer(queue string, ch chan<- []byte, ctx context.Context) {
	logger := log.Default()
	eventChannel, err := srv.Conn.Channel()
	if err != nil {
		logger.Println(err)
	}
	msgs, err := eventChannel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Println(err)
	}

	for d := range msgs {
		ch <- d.Body
	}

}
