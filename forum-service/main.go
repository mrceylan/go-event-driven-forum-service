package main

import (
	"forum-service/constants"
	dataaccess "forum-service/data-access"
	"forum-service/data-access/mongodb/repositories"
	eventserver "forum-service/event-server"
	"forum-service/event-server/rabbitmq"
	"forum-service/server"
	"forum-service/services/message"
	"forum-service/services/topic"
	"log"

	"github.com/spf13/viper"
)

var db dataaccess.IDatabase
var eventServer eventserver.IEventServer

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
	err := initEventServer()
	if err != nil {
		return err
	}
	topicSrv, messageSrv, err := initServices()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = initWebServer(topicSrv, messageSrv)
	if err != nil {
		return err
	}
	return err
}

func initServices() (topic.ITopicService, message.IMessageService, error) {
	var err error
	db, err = repositories.NewMongoDb(viper.GetString(constants.MONGO_CONNECTION), viper.GetInt(constants.MONGO_CONNECTION_TIMEOUT))
	if err != nil {
		return nil, nil, err
	}
	topicSrv := topic.Activate(db)
	messageSrv := message.Activate(db, eventServer)
	return topicSrv, messageSrv, nil
}

func initWebServer(topicSrv topic.ITopicService, messageSrv message.IMessageService) error {
	srv := server.NewServer(viper.GetInt(constants.PORT), topicSrv, messageSrv)
	err := srv.StartServer()
	return err
}

func initEventServer() error {
	settings := rabbitmq.RabbitMqServerSettings{
		ServerUrl: viper.GetString(constants.RABBITMQ_CONNECTION),
	}
	var err error
	eventServer, err = rabbitmq.NewRabbitMqServer(settings)
	if err != nil {
		return err
	}
	err = eventServer.DeclareQueue(viper.GetString(constants.MESSAGE_CREATED_QUEUE))
	if err != nil {
		return err
	}
	err = eventServer.DeclareQueue(viper.GetString(constants.MESSAGE_DELETED_QUEUE))
	if err != nil {
		return err
	}
	return nil
}

func closeConnections() {
	db.Disconnect(viper.GetInt(constants.MONGO_DISCONNECT_TIMEOUT))
	eventServer.Disconnect()
}
