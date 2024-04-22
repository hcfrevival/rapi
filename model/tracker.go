package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type EventTracker struct {
	ID           primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	EventName    string               `json:"event_name" bson:"event_name"`
	StartTime    uint64               `json:"start_time" bson:"start_time"`
	EndTime      uint64               `json:"end_time" bson:"end_time"`
	Duration     uint64               `json:"duration" bson:"duration"`
	Leaderboard  map[string]uint64    `json:"leaderboard,omitempty" bson:"leaderboard,omitempty"`
	Participants []EventTrackerPlayer `json:"participants,omitempty" bson:"participants,omitempty"`
	Entries      []EventTrackerEntry  `json:"entries,omitempty" bson:"entries,omitempty"`
}

type EventTrackerPlayer struct {
	PlayerUUID string            `json:"uuid" bson:"uuid"`
	PlayerName string            `json:"username,omitempty" bson:"username,omitempty"`
	Values     map[string]uint64 `json:"values,omitempty" bson:"values,omitempty,truncate"`
}

type EventTrackerEntry struct {
	Type           string   `json:"type" bson:"type"`
	Time           uint64   `json:"time" bson:"time"`
	FactionName    string   `json:"faction_name,omitempty" bson:"faction_name,omitempty"`
	FactionID      string   `json:"faction_id,omitempty" bson:"faction_id,omitempty"`
	KillerID       string   `json:"killer_id,omitempty" bson:"killer_id,omitempty"`
	SlainID        string   `json:"slain_id,omitempty" bson:"slain_id,omitempty"`
	KillerUsername string   `json:"killer_username,omitempty" bson:"killer_username,omitempty"`
	SlainUsername  string   `json:"slain_username,omitempty" bson:"slain_username,omitempty"`
	Location       Location `json:"location,omitempty" bson:"location,omitempty"`
}
