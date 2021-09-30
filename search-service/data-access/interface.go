package dataaccess

import (
	"context"
	"search-service/models"
)

type IDatabase interface {
	SaveMessage(ctx context.Context, model models.Message) error
	DeleteMessage(ctx context.Context, id string) error
	SearchMessages(ctx context.Context, searchString string) ([]models.MessageSearchResult, error)

	Disconnect(timeOut int) error
}
