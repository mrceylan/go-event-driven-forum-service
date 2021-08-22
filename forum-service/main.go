package main

import (
	"forum-service/constants"
	dataaccess "forum-service/data-access"
	"forum-service/data-access/mongodb/repositories"
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
	topicSrv, messageSrv, err := initServices()
	if err != nil {
		return err
	}
	err = initWebServer(topicSrv, messageSrv)
	return err
}

func initServices() (topic.ITopicService, message.IMessageService, error) {
	var err error
	db, err = repositories.NewMongoDb(viper.GetString(constants.MONGO_CONNECTION), viper.GetInt(constants.MONGO_CONNECTION_TIMEOUT))
	if err != nil {
		return nil, nil, err
	}
	topicSrv := topic.Activate(db)
	messageSrv := message.Activate(db)
	return topicSrv, messageSrv, nil
}

func initWebServer(topicSrv topic.ITopicService, messageSrv message.IMessageService) error {
	srv := server.NewServer(viper.GetInt(constants.PORT), topicSrv, messageSrv)
	err := srv.StartServer()
	return err
}

func closeConnections() {
	db.Disconnect(viper.GetInt(constants.MONGO_DISCONNECT_TIMEOUT))
}
