package interfaces

import "forum-service/models"

type MessageRepository interface {
	SaveMessage(message models.Message) (models.Message, error)
	GetMessage(id string) (models.Message, error)
}
