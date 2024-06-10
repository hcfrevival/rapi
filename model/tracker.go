package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type EventTracker struct {
	ID           primitive.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty"`
	Metadata     EventTrackerMetadata      `json:"metadata,omitempty" bson:"metadata"`
	Leaderboard  map[string]uint64         `json:"leaderboard,omitempty" bson:"leaderboard,omitempty"`
	Winner       EventTrackerFactionFields `json:"winner,omitempty" bson:"winner,omitempty"`
	Participants []EventTrackerPlayer      `json:"participants,omitempty" bson:"participants,omitempty"`
	Entries      []EventTrackerEntry       `json:"entries,omitempty" bson:"entries,omitempty"`
}

type EventTrackerMetadata struct {
	EventName string `json:"name" bson:"name"`
	StartTime uint64 `json:"start_time" bson:"start_time"`
	EndTime   uint64 `json:"end_time" bson:"end_time"`
	Duration  uint64 `json:"duration" bson:"duration"`
}

type EventTrackerPlayer struct {
	UUID   string            `json:"id" bson:"id"`
	Name   string            `json:"name,omitempty" bson:"name,omitempty"`
	Values map[string]uint64 `json:"values,omitempty" bson:"values,omitempty,truncate"`
}

type EventTrackerFactionFields struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type EventTrackerPlayerFields struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type EventTrackerEntry struct {
	Type           string                    `json:"type" bson:"type"`
	Time           uint64                    `json:"time" bson:"time"`
	SlainPlayer    EventTrackerPlayerFields  `json:"slain,omitempty" bson:"slain,omitempty"`
	KillerPlayer   EventTrackerPlayerFields  `json:"killer,omitempty" bson:"killer,omitempty"`
	Faction        EventTrackerFactionFields `json:"faction,omitempty" bson:"faction,omitempty"`
	NewTicketCount int8                      `json:"new_ticket_count,omitempty" bson:"new_ticket_count,omitempty"`
	TicketsLost    int8                      `json:"tickets_lost,omitempty" bson:"tickets_lost,omitempty"`
	Location       Location                  `json:"location,omitempty" bson:"location,omitempty"`
}
