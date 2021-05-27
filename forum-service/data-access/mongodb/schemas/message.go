package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Message    string             `bson:"message,omitempty"`
	CreateDate time.Time          `bson:"createDate,omitempty"`
	CreatedBy  primitive.ObjectID `bson:"createdBy,omitempty"`
}
