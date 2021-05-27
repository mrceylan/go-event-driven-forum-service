package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topic struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Header     string             `bson:"header,omitempty"`
	CreateDate time.Time          `bson:"createDate,omitempty"`
	CreatedBy  primitive.ObjectID `bson:"createdBy,omitempty"`
}
