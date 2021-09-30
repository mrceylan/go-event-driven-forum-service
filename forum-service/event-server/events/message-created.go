package events

import (
	"forum-service/constants"
	eventserver "forum-service/event-server"
	"forum-service/models"

	"github.com/spf13/viper"
)

type messsageCreatePayload struct {
	Id      string
	TopicId string
	Message string
}

func (payload *messsageCreatePayload) mapFromModel(model models.Message) {
	*payload = messsageCreatePayload{
		Id:      model.Id,
		TopicId: model.TopicId,
		Message: model.Message,
	}
}

type MessageCreatedEvent struct {
	EventServer eventserver.IEventServer
}

func getMessageCreatedQueueName() string {
	return viper.GetString(constants.MESSAGE_CREATED_QUEUE)
}

func (ms MessageCreatedEvent) PublishEvent(model models.Message) {
	var payload messsageCreatePayload
	payload.mapFromModel(model)
	ms.EventServer.PublishEvent(getMessageCreatedQueueName(), payload)
}
