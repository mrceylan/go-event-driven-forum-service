package events

import (
	"context"
	"encoding/json"
	"log"
	"search-service/constants"
	eventserver "search-service/event-server"
	"search-service/models"
	"search-service/services/search"

	"github.com/spf13/viper"
)

type messsageCreatePayload struct {
	Id      string
	TopicId string
	Message string
}

func (payload *messsageCreatePayload) mapToModel() models.Message {
	return models.Message{
		Id:      payload.Id,
		TopicId: payload.TopicId,
		Message: payload.Message,
	}
}

type MessageCreatedEvent struct {
	EventServer   eventserver.IEventServer
	SearchService search.ISearchService
}

func getMessageCreatedQueueName() string {
	return viper.GetString(constants.MESSAGE_CREATED_QUEUE)
}

func (ms MessageCreatedEvent) SubscribeToEvents() {
	ch := make(chan []byte)
	ctx := context.Background()
	go ms.EventServer.StartConsumer(getMessageCreatedQueueName(), ch, ctx)
	for event := range ch {
		var payload messsageCreatePayload
		if err := json.Unmarshal(event, &payload); err != nil {
			log.Println(err)
		}
		log.Println(event)
		err := ms.SearchService.SaveMessage(ctx, payload.mapToModel())
		log.Println(err)
	}
}
