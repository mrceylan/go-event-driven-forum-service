package repositories

import (
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
