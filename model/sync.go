package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SyncServer struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ServerID        uint8              `json:"id" bson:"id"`
	Name            string             `json:"name" bson:"name"`
	ProxyName       string             `json:"proxy_name" bson:"proxy_name"`
	Description     []string           `json:"description" bson:"description"`
	OnlineUsernames []string           `json:"online_usernames" bson:"online_usernames"`
	Type            string             `json:"type" bson:"type"`
	Status          string             `json:"status" bson:"status"`
}
