package main

import (
	"log"
	"search-service/constants"
	dataaccess "search-service/data-access"
	"search-service/data-access/elastic/repositories"
	"search-service/event-server/events"
	"search-service/event-server/rabbitmq"
	"search-service/server"
	"search-service/services/search"

	"github.com/spf13/viper"
)

var db dataaccess.IDatabase

func main() {
	initConfig()
	err := initApp()
	defer closeConnections()

	if err != nil {
		log.Fatal(err)
	}

}

func initConfig() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func initApp() error {
	searchSrv, err := initServices()
	if err != nil {
		return err
	}
	err = initEventServer(searchSrv)
	if err != nil {
		return err
	}
	err = initWebServer(searchSrv)
	return err
}

func initServices() (search.ISearchService, error) {
	var err error
	db, err = repositories.NewElasticClient(repositories.ElasticConfig{
		Addresses: viper.GetStringSlice(constants.ELASTIC_ADDRESSES),
		UserName:  viper.GetString(constants.ELASTIC_USERNAME),
		Password:  viper.GetString(constants.ELASTIC_PASSWORD),
	})
	if err != nil {
		return nil, err
	}
	searchSrv := search.Activate(db)
	return searchSrv, nil
}

func initWebServer(searchSrv search.ISearchService) error {
	srv := server.NewServer(viper.GetInt(constants.PORT), searchSrv)
	err := srv.StartServer()
	return err
}

func initEventServer(searchSrv search.ISearchService) error {
	srv, err := rabbitmq.NewRabbitMqServer(rabbitmq.RabbitMqServerSettings{
		ServerUrl: viper.GetString(constants.RABBITMQ_CONNECTION),
	})
	if err != nil {
		return err
	}

	go events.MessageCreatedEvent{EventServer: srv, SearchService: searchSrv}.SubscribeToEvents()
	go events.MessageDeletedEvent{EventServer: srv, SearchService: searchSrv}.SubscribeToEvents()

	return nil
}

func closeConnections() {
	db.Disconnect(0)
}
