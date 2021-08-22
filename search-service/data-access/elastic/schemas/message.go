package schemas

import "search-service/models"

type Message struct {
	Id      string `json:"id"`
	TopicId string `json:"topicId"`
	Message string `json:"message"`
}

func (ms *Message) MapToModel() models.Message {
	return models.Message{
		Id:      ms.Id,
		TopicId: ms.TopicId,
		Message: ms.Message,
	}
}

func (ms *Message) MapFromModel(model models.Message) {

	ms.TopicId = model.TopicId
	ms.Message = model.Message
	ms.Id = model.Id
}
