package main

import (
	"context"
	"log"
	"user-service/data-access/mongodb/repositories"
	"user-service/server"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("UserService")
	myDb := repositories.Database{
		Db:  db,
		Ctx: ctx,
	}
	StartWebServer(myDb)
	log.Println("Server started to listening..")
}

func StartWebServer(db repositories.Database) {
	srv := server.Server{
		Port:       8080,
		Repository: &db,
	}
	srv.CreateServer()
}
