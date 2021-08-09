package main

import (
	"log"
	"user-service/constants"
	dataaccess "user-service/data-access"
	"user-service/data-access/mongodb/repository"
	"user-service/server"
	"user-service/services/auth"
	"user-service/services/user"
	"user-service/utils"

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
	authSrv, userSrv, err := initServices()
	if err != nil {
		return err
	}
	err = initWebServer(authSrv, userSrv)
	return err
}

func initServices() (auth.IAuthService, user.IUserService, error) {
	var err error
	db, err = repository.NewMongoDb(viper.GetString(constants.MONGO_CONNECTION), viper.GetInt(constants.MONGO_CONNECTION_TIMEOUT))
	if err != nil {
		return nil, nil, err
	}
	jwtutil := utils.NewJwtUtil(viper.GetString(constants.JWT_SECRET_KEY), viper.GetInt(constants.JWT_EXPIRATION_MINUTES))
	authSrv := auth.Activate(db, jwtutil)
	userSrv := user.Activate(db)
	return authSrv, userSrv, nil
}

func initWebServer(authSrv auth.IAuthService, userSrv user.IUserService) error {
	srv := server.NewServer(viper.GetInt(constants.PORT), userSrv, authSrv)
	err := srv.StartServer()
	return err
}

func closeConnections() {
	db.Disconnect(viper.GetInt(constants.MONGO_DISCONNECT_TIMEOUT))
}
