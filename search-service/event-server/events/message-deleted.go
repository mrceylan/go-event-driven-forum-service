package events

import (
	"context"
	"encoding/json"
	"log"
	"search-service/constants"
	eventserver "search-service/event-server"
	"search-service/services/search"

	"github.com/spf13/viper"
)

type messsageDeletePayload struct {
	Id string
}

type MessageDeletedEvent struct {
	EventServer   eventserver.IEventServer
	SearchService search.ISearchService
}

func getMessageDeletedQueueName() string {
	return viper.GetString(constants.MESSAGE_DELETED_QUEUE)
}

func (ms MessageDeletedEvent) SubscribeToEvents() {
	ch := make(chan []byte)
	ctx := context.Background()
	go ms.EventServer.StartConsumer(getMessageDeletedQueueName(), ch, ctx)
	for event := range ch {
		var payload messsageDeletePayload
		if err := json.Unmarshal(event, &payload); err != nil {
			log.Println(err)
		}
		err := ms.SearchService.DeleteMessage(ctx, payload.Id)
		log.Println(err)
	}
}
