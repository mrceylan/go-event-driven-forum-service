package repositories

import (
	"forum-service/data-access/mongodb/schemas"
	"forum-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Database) SaveTopic(header string, userId string) (models.Topic, error) {
	TopicCollection := db.GetTopicCollection()

	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println("Invalid user id")
	}
	entity := schemas.Topic{
		Header:    header,
		CreatedBy: userObjectId,
	}

	insertResult, err := TopicCollection.InsertOne(db.Ctx, entity)

	if err != nil {
		return models.Topic{}, err
	}

	return models.Topic{
		Id:        insertResult.InsertedID.(primitive.ObjectID).Hex(),
		Header:    header,
		CreatedBy: userId,
	}, nil
}

func (db *Database) GetTopicById(id string) (models.Topic, error) {
	TopicCollection := db.GetTopicCollection()

	var entity schemas.Topic

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	if err := TopicCollection.FindOne(db.Ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return models.Topic{}, err
	}

	return models.Topic{
		Id:         entity.Id.Hex(),
		Header:     entity.Header,
		CreatedBy:  entity.CreatedBy.Hex(),
		CreateDate: entity.CreateDate,
	}, nil

}
