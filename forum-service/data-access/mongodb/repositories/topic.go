package repositories

import (
	"forum-service/data-access/mongodb/schemas"
	"forum-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Database) SaveTopic(header string, userId string) (models.Topic, error) {
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println("Invalid user id")
	}
	entity := schemas.Topic{
		Header:    header,
		CreatedBy: userObjectId,
	}

	TopicCollection := db.GetTopicCollection()

	insertResult, err := TopicCollection.InsertOne(db.Ctx, entity)

	if err != nil {
		return models.Topic{}, err
	}
	insertResult.InsertedID
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
