package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"search-service/data-access/elastic/schemas"
	"search-service/models"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func (cl *ElasticClient) SaveMessage(ctx context.Context, message models.Message) error {

	var entity schemas.Message
	entity.MapFromModel(message)

	dataJSON, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      messageIndex(),
		DocumentID: entity.Id,
		Body:       strings.NewReader(string(dataJSON)),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, cl.Client)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), entity.Id)
	}

	return nil
}

func (cl *ElasticClient) SearchMessages(ctx context.Context, searchString string) ([]models.Message, error) {

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"message": searchString,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := cl.Client.Search(
		cl.Client.Search.WithContext(ctx),
		cl.Client.Search.WithIndex(messageIndex()),
		cl.Client.Search.WithBody(&buf),
		cl.Client.Search.WithTrackTotalHits(true),
		cl.Client.Search.WithPretty(),
	)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			// Print the response status and error information.
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	log.Println(res)

	var entities []schemas.Message

	if err := json.NewDecoder(res.Body).Decode(&entities); err != nil {
		return nil, err
	}

	var result []models.Message

	for _, entity := range entities {
		result = append(result, entity.MapToModel())
	}

	return result, nil
}
