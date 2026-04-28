package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	// Di v2, ObjectID diambil langsung dari paket bson
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}
