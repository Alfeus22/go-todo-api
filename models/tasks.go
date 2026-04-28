package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/v2/bson"
)

type Task struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	IsDone bool               `json:"isdone" bson:"isdone"`

	UserId primitive.ObjectID `bson:"user_id" json:"user_id"`
}
