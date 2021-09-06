package repositories

import (
	"encoding/json"
	"log"
	dataaccess "search-service/data-access"

	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticConfig struct {
	Addresses []string
	UserName  string
	Password  string
}

type ElasticClient struct {
	Client *elasticsearch.Client
}

func NewElasticClient(config ElasticConfig) (dataaccess.IDatabase, error) {
	defer log.Println("Database connection created...")
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.Addresses,
		Username:  config.UserName,
		Password:  config.Password,
	})
	if err != nil {
		return nil, err
	}
	return &ElasticClient{client}, nil
}

func (cl *ElasticClient) Disconnect(_timeOut int) error {
	return nil
}

func messageIndex() string {
	return "messages"
}

type ElasticSearchResult struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string          `json:"_index"`
			Type   string          `json:"_type"`
			ID     string          `json:"_id"`
			Score  float64         `json:"_score"`
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
