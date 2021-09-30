package message

import (
	"context"
	dataaccess "forum-service/data-access"
	eventserver "forum-service/event-server"
	"forum-service/event-server/events"
	"forum-service/models"
)

type IMessageService interface {
	SaveMessage(ctx context.Context, message models.Message) (models.Message, error)
	DeleteMessage(ctx context.Context, id string) error
	GetMessageById(ctx context.Context, id string) (models.Message, error)
}

type MessageService struct {
	Database    dataaccess.IDatabase
	EventServer eventserver.IEventServer
}

func Activate(db dataaccess.IDatabase, eventSrv eventserver.IEventServer) IMessageService {
	return &MessageService{db, eventSrv}
}

func (ms *MessageService) SaveMessage(ctx context.Context, model models.Message) (models.Message, error) {
	model, err := ms.Database.SaveMessage(ctx, model)

	if err != nil {
		return models.Message{}, err
	}

	events.MessageCreatedEvent{EventServer: ms.EventServer}.PublishEvent(model)

	return model, nil
}

func (ms *MessageService) DeleteMessage(ctx context.Context, id string) error {
	err := ms.Database.DeleteMessage(ctx, id)

	if err != nil {
		return err
	}

	events.MessageDeletedEvent{EventServer: ms.EventServer}.PublishEvent(id)

	return nil
}

func (ms *MessageService) GetMessageById(ctx context.Context, id string) (models.Message, error) {
	model, err := ms.Database.GetMessageById(ctx, id)

	if err != nil {
		return models.Message{}, err
	}

	return model, nil
}
