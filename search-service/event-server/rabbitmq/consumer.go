package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"search-service/models"
)

func (srv *RabbitMqServer) StartConsumer() {
	logger := log.Default()
	eventChannel, err := srv.Conn.Channel()
	if err != nil {
		logger.Print(err)
	}
	msgs, err := eventChannel.Consume(
		"deneme",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Print(err)
	}

	forever := make(chan bool)
	ctx := context.Background()
	go func() {
		for d := range msgs {
			var model models.Message
			if err := json.Unmarshal(d.Body, &model); err != nil {
				logger.Print(err)
			}
			err = srv.SearchService.SaveMessage(ctx, model)
			if err != nil {
				logger.Print(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
