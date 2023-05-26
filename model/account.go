package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Account struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Username  string             `json:"username" bson:"username"`
	FirstSeen time.Time          `json:"first_seen" bson:"first_seen"`
	LastSeen  time.Time          `json:"last_seen" bson:"last_seen"`
}
