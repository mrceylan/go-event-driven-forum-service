package dataaccess

import (
	"context"
	"user-service/models"
)

type IDatabase interface {
	SaveUser(Ctx context.Context, u models.User, password string) (models.User, error)
	GetUserById(Ctx context.Context, id string) (models.User, error)
	GetUserByEmail(Ctx context.Context, email string) (models.User, error)
	GetUserPasswordById(Ctx context.Context, id string) string

	Disconnect(timeOut int) error
}
