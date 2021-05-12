package interfaces

import (
	"user-service/models"
)

type Repository interface {
	SaveUser(u models.User, password string) (models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserPasswordById(id string) string
}
