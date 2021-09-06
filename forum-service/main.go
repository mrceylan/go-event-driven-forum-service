package main

import (
	"forum-service/constants"
	dataaccess "forum-service/data-access"
	"forum-service/data-access/mongodb/repositories"
	eventserver "forum-service/event-server"
	"forum-service/models"
	"forum-service/server"
	"forum-service/services/message"
	"forum-service/services/topic"
	"log"

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
	publishChannel := make(chan models.PublishEvent, 100)
	topicSrv, messageSrv, err := initServices(publishChannel)
	if err != nil {
		return err
	}
	err = initEventServer(eventserver.EventServerSettings{PublishChannel: publishChannel})
	if err != nil {
		return err
	}
	err = initWebServer(topicSrv, messageSrv)
	if err != nil {
		return err
	}
	return err
}

func initServices(publishChannel chan<- models.PublishEvent) (topic.ITopicService, message.IMessageService, error) {
	var err error
	db, err = repositories.NewMongoDb(viper.GetString(constants.MONGO_CONNECTION), viper.GetInt(constants.MONGO_CONNECTION_TIMEOUT))
	if err != nil {
		return nil, nil, err
	}
	topicSrv := topic.Activate(db)
	messageSrv := message.Activate(db, publishChannel)
	return topicSrv, messageSrv, nil
}

func initWebServer(topicSrv topic.ITopicService, messageSrv message.IMessageService) error {
	srv := server.NewServer(viper.GetInt(constants.PORT), topicSrv, messageSrv)
	err := srv.StartServer()
	return err
}

func initEventServer(settings eventserver.EventServerSettings) error {
	srv, err := eventserver.NewEventServer(settings)
	if err != nil {
		return err
	}
	go srv.StartPublisher()
	return nil
}

func closeConnections() {
	db.Disconnect(viper.GetInt(constants.MONGO_DISCONNECT_TIMEOUT))
}
