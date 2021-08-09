package services

import (
	"forum-service/data-access/interfaces"
	"forum-service/models"
)

type TopicService struct {
	TopicRepo interfaces.TopicRepository
}

func (ts *TopicService) CreateTopic(header string, userId string) (models.Topic, error) {
	model, err := ts.TopicRepo.SaveTopic(header, userId)

	if err != nil {
		return models.Topic{}, err
	}

	return model, nil
}

func (ts *TopicService) GetTopicById(id string) (models.Topic, error) {
	model, err := ts.TopicRepo.GetTopicById(id)

	if err != nil {
		return models.Topic{}, err
	}

	return model, nil
}
