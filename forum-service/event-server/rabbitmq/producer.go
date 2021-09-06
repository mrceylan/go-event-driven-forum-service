package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func (c *RabbitMqServer) StartPublisher() {
	logger := log.Default()
	eventChannel, err := c.Conn.Channel()
	if err != nil {
		logger.Print(err)
	}
	for event := range c.PublishChannel {
		logger.Print(event)
		data, err := json.Marshal(event.Event)
		logger.Print(err)
		err = eventChannel.Publish(
			"", // exchange
			event.Topic,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			})
		logger.Print(err)
	}
}
