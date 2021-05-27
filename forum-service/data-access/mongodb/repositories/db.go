package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Db  *mongo.Database
	Ctx context.Context
}

func (db *Database) GetTopicCollection() *mongo.Collection {
	return db.Db.Collection("Topics")
}

func (db *Database) GetMessageCollection() *mongo.Collection {
	return db.Db.Collection("Messages")
}
