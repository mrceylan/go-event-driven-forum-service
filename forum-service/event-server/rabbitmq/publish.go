package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func (srv *RabbitMqServer) PublishEvent(queue string, eventPayload interface{}) {
	data, err := json.Marshal(eventPayload)
	log.Println(err)
	err = srv.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
	log.Println(err)
}
