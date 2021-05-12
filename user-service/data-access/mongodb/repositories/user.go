package repositories

import (
	"log"
	"user-service/data-access/mongodb/schemas"
	"user-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Database) SaveUser(u models.User, password string) (models.User, error) {
	entity := schemas.User{
		UserName:     u.UserName,
		Email:        u.Email,
		PasswordHash: password,
	}

	UserCollection := db.GetUserCollection()

	insertResult, err := UserCollection.InsertOne(db.Ctx, entity)

	if err != nil {
		return models.User{}, err
	}

	u.Id = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return u, nil
}

func (db *Database) GetUserById(id string) (models.User, error) {
	UserCollection := db.GetUserCollection()

	var entity schemas.User

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	if err := UserCollection.FindOne(db.Ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:       entity.Id.Hex(),
		UserName: entity.UserName,
		Email:    entity.Email,
	}, nil

}

func (db *Database) GetUserByEmail(email string) (models.User, error) {
	UserCollection := db.GetUserCollection()

	var entity schemas.User

	if err := UserCollection.FindOne(db.Ctx, bson.M{"email": email}).Decode(&entity); err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:       entity.Id.Hex(),
		UserName: entity.UserName,
		Email:    entity.Email,
	}, nil

}

func (db *Database) GetUserPasswordById(id string) string {
	UserCollection := db.GetUserCollection()

	var entity schemas.User

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	if err := UserCollection.FindOne(db.Ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return ""
	}

	return entity.PasswordHash

}
