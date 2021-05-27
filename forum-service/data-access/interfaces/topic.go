package interfaces

import "forum-service/models"

type TopicRepository interface {
	SaveTopic(header string, userId string) (models.Topic, error)
	GetTopic(id string) (models.Topic, error)
}
