package events

import (
	"forum-service/constants"
	eventserver "forum-service/event-server"

	"github.com/spf13/viper"
)

type messsageDeletePayload struct {
	Id string
}

type MessageDeletedEvent struct {
	EventServer eventserver.IEventServer
}

func getMessageDeletedQueueName() string {
	return viper.GetString(constants.MESSAGE_DELETED_QUEUE)
}

func (ms MessageDeletedEvent) PublishEvent(id string) {
	payload := messsageDeletePayload{id}
	ms.EventServer.PublishEvent(getMessageDeletedQueueName(), payload)
}
