package repositories

import (
	"context"
	"forum-service/data-access/mongodb/schemas"
	"forum-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cl *MongoClient) SaveTopic(ctx context.Context, model models.Topic) (models.Topic, error) {
	TopicCollection := cl.forumDatabase().topicsCollection()

	userObjectId, err := primitive.ObjectIDFromHex(model.CreatedBy)
	if err != nil {
		log.Println("Invalid user id")
	}
	entity := schemas.Topic{
		Header:     model.Header,
		CreatedBy:  userObjectId,
		CreateDate: model.CreateDate,
	}

	insertResult, err := TopicCollection.InsertOne(ctx, entity)

	if err != nil {
		return models.Topic{}, err
	}

	model.Id = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return model, nil
}

func (cl *MongoClient) GetTopicById(ctx context.Context, id string) (models.Topic, error) {
	TopicCollection := cl.forumDatabase().topicsCollection()

	var entity schemas.Topic

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	if err := TopicCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return models.Topic{}, err
	}

	return models.Topic{
		Id:         entity.Id.Hex(),
		Header:     entity.Header,
		CreatedBy:  entity.CreatedBy.Hex(),
		CreateDate: entity.CreateDate,
	}, nil

}
