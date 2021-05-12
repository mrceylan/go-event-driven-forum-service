package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Db  *mongo.Database
	Ctx context.Context
}

func (db *Database) GetUserCollection() *mongo.Collection {
	return db.Db.Collection("Users")
}
