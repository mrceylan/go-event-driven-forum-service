package dataaccess

import (
	"context"
	"forum-service/models"
)

type IDatabase interface {
	SaveMessage(ctx context.Context, message models.Message) (models.Message, error)
	GetMessageById(ctx context.Context, id string) (models.Message, error)
	GetTopicMessages(ctx context.Context, id string) ([]models.Message, error)

	SaveTopic(ctx context.Context, model models.Topic) (models.Topic, error)
	GetTopicById(ctx context.Context, id string) (models.Topic, error)

	Disconnect(timeOut int) error
}
