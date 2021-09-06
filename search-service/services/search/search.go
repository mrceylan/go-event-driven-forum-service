package search

import (
	"context"
	dataaccess "search-service/data-access"
	"search-service/models"
)

type ISearchService interface {
	SaveMessage(ctx context.Context, message models.Message) error
	SearchMessages(ctx context.Context, searchString string) ([]models.MessageSearchResult, error)
}

type SearchService struct {
	Database dataaccess.IDatabase
}

func Activate(db dataaccess.IDatabase) ISearchService {
	return &SearchService{db}
}

func (ss *SearchService) SaveMessage(ctx context.Context, model models.Message) error {
	err := ss.Database.SaveMessage(ctx, model)

	if err != nil {
		return err
	}

	return nil
}

func (ss *SearchService) SearchMessages(ctx context.Context, searchString string) ([]models.MessageSearchResult, error) {
	result, err := ss.Database.SearchMessages(ctx, searchString)

	if err != nil {
		return nil, err
	}

	return result, nil
}
