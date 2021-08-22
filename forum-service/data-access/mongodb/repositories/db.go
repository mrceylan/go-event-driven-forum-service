package repositories

import (
	"context"
	"forum-service/constants"
	dataaccess "forum-service/data-access"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

type MongoDb struct {
	Database *mongo.Database
}

func NewMongoDb(conn string, timeOut int) (dataaccess.IDatabase, error) {
	defer log.Println("Database connection created...")
	client, err := mongo.NewClient(options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(timeOut))
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &MongoClient{client}, nil
}

func (cl *MongoClient) forumDatabase() *MongoDb {
	return &MongoDb{cl.Client.Database(viper.GetString(constants.MONGO_DB_NAME))}
}

func (db *MongoDb) topicsCollection() *mongo.Collection {
	return db.Database.Collection("Topics")
}

func (db *MongoDb) messagesCollection() *mongo.Collection {
	return db.Database.Collection("Messages")
}

func (cl *MongoClient) Disconnect(timeOut int) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(timeOut))
	err := cl.Client.Disconnect(ctx)

	return err
}
