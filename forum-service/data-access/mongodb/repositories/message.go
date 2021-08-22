package repositories

import (
	"context"
	"forum-service/data-access/mongodb/schemas"
	"forum-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (cl *MongoClient) SaveMessage(ctx context.Context, message models.Message) (models.Message, error) {
	MessageCollection := cl.forumDatabase().messagesCollection()

	var entity schemas.Message
	err := entity.MapFromModel(message)

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

func (cl *MongoClient) GetMessageById(ctx context.Context, id string) (models.Message, error) {
	MessageCollection := cl.forumDatabase().messagesCollection()

	var entity schemas.Message

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	if err := MessageCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&entity); err != nil {
		return models.Message{}, err
	}

	return entity.MapToModel(), nil

}

func (cl *MongoClient) GetTopicMessages(ctx context.Context, id string) ([]models.Message, error) {
	MessageCollection := cl.forumDatabase().messagesCollection()

	var result []models.Message

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
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
