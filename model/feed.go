package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kill struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	KillerID       string             `json:"killer_id,omitempty" bson:"killer_id,omitempty"`
	KillerUsername string             `json:"killer_username,omitempty" bson:"killer_username,omitempty"`
	KilledID       string             `json:"slain_id" bson:"slain_id"`
	KilledUsername string             `json:"slain_username" bson:"slain_username"`
	DeathMessage   string             `json:"death_message,omitempty" bson:"death_message,omitempty"`
	Map            uint8              `json:"map" bson:"map"`
	Date           uint64             `json:"date" bson:"date"`
}

type Death struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	KilledID       string             `json:"slain_id" bson:"slain_id"`
	KilledUsername string             `json:"slain_username" bson:"slain_username"`
	DeathMessage   string             `json:"death_message" bson:"death_message"`
	Map            uint8              `json:"map" bson:"map"`
	Date           uint64             `json:"date" bson:"date"`
}

type EventCapture struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FactionID      string             `json:"faction_id" bson:"faction_id"`
	FactionName    string             `json:"faction_name" bson:"faction_name"`
	EventName      string             `json:"event_name" bson:"event_name"`
	FactionMembers []string           `json:"faction_members" bson:"faction_members"`
	Map            uint8              `json:"map" bson:"map"`
	Date           uint64             `json:"date" bson:"date"`
}
