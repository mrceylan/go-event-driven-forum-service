package repositories

import (
	"context"
	"forum-service/data-access/mongodb/schemas"
	"forum-service/data-access/mongodb/utils"
	"forum-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (cl *MongoClient) SaveMessage(ctx context.Context, message models.Message) (models.Message, error) {
	MessageCollection := cl.getForumDatabase().getMessagesCollection()

	var entity schemas.Message
	err := entity.MapFromModel(message)
	entity.IsDeleted = false

	if err != nil {
		return models.Message{}, err
	}

	insertResult, err := MessageCollection.InsertOne(ctx, entity)

	if err != nil {
		return models.Message{}, err
	}

	message.Id = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return message, nil
}

func (cl *MongoClient) DeleteMessage(ctx context.Context, id string) error {
	MessageCollection := cl.getForumDatabase().getMessagesCollection()

	objectId, err := utils.ConvertStringToObjectId(id)
	if err != nil {
		return err
	}

	_, err = MessageCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectId},
		bson.D{
			{"$set", bson.D{{"isDeleted", true}}},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (cl *MongoClient) GetMessageById(ctx context.Context, id string) (models.Message, error) {
	MessageCollection := cl.getForumDatabase().getMessagesCollection()

	var entity schemas.Message

	objectId, err := utils.ConvertStringToObjectId(id)
	if err != nil {
		return models.Message{}, err
	}

	if err := MessageCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return models.Message{}, err
	}

	return entity.MapToModel(), nil

}

func (cl *MongoClient) GetTopicMessages(ctx context.Context, id string) ([]models.Message, error) {
	MessageCollection := cl.getForumDatabase().getMessagesCollection()

	var result []models.Message

	objectId, err := utils.ConvertStringToObjectId(id)
	if err != nil {
		return []models.Message{}, err
	}

	var cursor *mongo.Cursor

	if cursor, err = MessageCollection.Find(ctx, bson.M{"topicId": objectId}); err != nil {
		return []models.Message{}, err
	}

	var entities []schemas.Message
	err = cursor.All(ctx, &entities)
	if err != nil {
		return []models.Message{}, err
	}

	for _, entity := range entities {
		result = append(result, entity.MapToModel())
	}

	return result, nil

}
