package message

import (
	"context"
	dataaccess "forum-service/data-access"
	"forum-service/models"
)

type IMessageService interface {
	SaveMessage(ctx context.Context, message models.Message) (models.Message, error)
	GetMessageById(ctx context.Context, id string) (models.Message, error)
}

type MessageService struct {
	Database       dataaccess.IDatabase
	PublishChannel chan<- models.PublishEvent
}

func Activate(db dataaccess.IDatabase, publishChannel chan<- models.PublishEvent) IMessageService {
	return &MessageService{db, publishChannel}
}

func (ms *MessageService) SaveMessage(ctx context.Context, model models.Message) (models.Message, error) {
	model, err := ms.Database.SaveMessage(ctx, model)

	if err != nil {
		return models.Message{}, err
	}

	ms.PublishChannel <- models.PublishEvent{Topic: "deneme", Event: model}

	return model, nil
}

func (ms *MessageService) GetMessageById(ctx context.Context, id string) (models.Message, error) {
	model, err := ms.Database.GetMessageById(ctx, id)

	if err != nil {
		return models.Message{}, err
	}

	return model, nil
}
