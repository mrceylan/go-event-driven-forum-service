package schemas

import (
	"forum-service/data-access/mongodb/utils"
	"forum-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	TopicId    primitive.ObjectID `bson:"topicId,omitempty"`
	Message    string             `bson:"message,omitempty"`
	CreateDate time.Time          `bson:"createDate,omitempty"`
	CreatedBy  primitive.ObjectID `bson:"createdBy,omitempty"`
	IsDeleted  bool               `bson:"isDeleted,omitempty"`
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

	userObjectId, err := utils.ConvertStringToObjectId(model.CreatedBy)
	if err != nil {
		return err
	}
	topicId, err := utils.ConvertStringToObjectId(model.TopicId)
	if err != nil {
		return err
	}

	if model.Id != "" {
		id, err := utils.ConvertStringToObjectId(model.Id)
		if err != nil {
			return err
		}
		ms.Id = id
	}

	ms.CreatedBy = userObjectId
	ms.CreateDate = model.CreateDate
	ms.TopicId = topicId
	ms.Message = model.Message

	return nil
}
