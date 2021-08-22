package schemas

import (
	"forum-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	TopicId    primitive.ObjectID `bson:"topicId,omitempty"`
	Message    string             `bson:"message,omitempty"`
	CreateDate time.Time          `bson:"createDate,omitempty"`
	CreatedBy  primitive.ObjectID `bson:"createdBy,omitempty"`
}

func (ms *Message) MapToModel() models.Message {
	return models.Message{
		Id:         ms.Id.Hex(),
		TopicId:    ms.TopicId.Hex(),
		Message:    ms.Message,
		CreateDate: ms.CreateDate,
		CreatedBy:  ms.CreatedBy.Hex(),
	}
}

func (ms *Message) MapFromModel(model models.Message) error {

	userObjectId, err := primitive.ObjectIDFromHex(model.CreatedBy)
	if err != nil {
		log.Println("Invalid user id")
	}
	topicId, err := primitive.ObjectIDFromHex(model.TopicId)
	if err != nil {
		log.Println("Invalid topic id")
	}

	if model.Id != "" {
		id, err := primitive.ObjectIDFromHex(model.Id)
		if err != nil {
			log.Println("Invalid message id")
		}
		ms.Id = id
	}

	ms.CreatedBy = userObjectId
	ms.CreateDate = model.CreateDate
	ms.TopicId = topicId
	ms.Message = model.Message

	return nil
}
