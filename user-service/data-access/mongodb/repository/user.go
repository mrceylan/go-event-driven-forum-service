package repository

import (
	"context"
	"log"
	"user-service/data-access/mongodb/schemas"
	"user-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cl *MongoClient) SaveUser(ctx context.Context, u models.User, password string) (models.User, error) {
	entity := schemas.User{
		UserName:     u.UserName,
		Email:        u.Email,
		PasswordHash: password,
	}

	collection := cl.usersDatabase().usersCollection()

	insertResult, err := collection.InsertOne(ctx, entity)

	if err != nil {
		return models.User{}, err
	}

	u.Id = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return u, nil
}

func (cl *MongoClient) GetUserById(ctx context.Context, id string) (models.User, error) {
	collection := cl.usersDatabase().usersCollection()

	var entity schemas.User

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	if err := collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:       entity.Id.Hex(),
		UserName: entity.UserName,
		Email:    entity.Email,
	}, nil

}

func (cl *MongoClient) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	collection := cl.usersDatabase().usersCollection()

	var entity schemas.User

	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&entity); err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:       entity.Id.Hex(),
		UserName: entity.UserName,
		Email:    entity.Email,
	}, nil

}

func (cl *MongoClient) GetUserPasswordById(ctx context.Context, id string) string {
	collection := cl.usersDatabase().usersCollection()

	var entity schemas.User

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	if err := collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return ""
	}

	return entity.PasswordHash

}
