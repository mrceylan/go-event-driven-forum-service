package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertStringToObjectId(id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid objectId %s", id)
	}

	return objectId, nil
}
