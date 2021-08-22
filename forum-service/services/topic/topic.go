package topic

import (
	"context"
	dataaccess "forum-service/data-access"
	"forum-service/models"
)

type ITopicService interface {
	CreateTopic(ctx context.Context, model models.Topic) (models.Topic, error)
	GetTopicById(ctx context.Context, id string) (models.Topic, error)
	GetTopicMessages(ctx context.Context, id string) ([]models.Message, error)
}

type TopicService struct {
	Database dataaccess.IDatabase
}

func Activate(db dataaccess.IDatabase) ITopicService {
	return &TopicService{db}
}

func (ts *TopicService) CreateTopic(ctx context.Context, model models.Topic) (models.Topic, error) {
	model, err := ts.Database.SaveTopic(ctx, model)

	if err != nil {
		return models.Topic{}, err
	}

	return model, nil
}

func (ts *TopicService) GetTopicById(ctx context.Context, id string) (models.Topic, error) {
	model, err := ts.Database.GetTopicById(ctx, id)

	if err != nil {
		return models.Topic{}, err
	}

	return model, nil
}

func (ts *TopicService) GetTopicMessages(ctx context.Context, id string) ([]models.Message, error) {
	result, err := ts.Database.GetTopicMessages(ctx, id)

	if err != nil {
		return []models.Message{}, err
	}

	return result, nil
}
