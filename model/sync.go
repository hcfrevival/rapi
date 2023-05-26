package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SyncServer struct {
	ID              primitive.ObjectID
	ServerID        uint8
	Name            string
	ProxyName       string
	Description     []string
	OnlineUsernames []string
	Type            string
	Status          string
}
