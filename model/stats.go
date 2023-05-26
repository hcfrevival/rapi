package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatHolder struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID          string             `json:"uuid" bson:"uuid"`
	Map           uint8              `json:"map" bson:"map"`
	Kills         uint64             `json:"kills" bson:"KILL"`
	Deaths        uint64             `json:"deaths" bson:"DEATH"`
	Longshot      uint64             `json:"longshot" bson:"LONGSHOT"`
	EventCaptures uint64             `json:"events" bson:"EVENT_CAPTURES"`
	ExpEarned     uint64             `json:"exp" bson:"EXP_EARNED"`
	Playtime      uint64             `json:"playtime" bson:"PLAYTIME"`
}

type EStatisticType string

const (
	KILL           EStatisticType = "kill"
	DEATH          EStatisticType = "death"
	EVENT_CAPTURES EStatisticType = "event_captures"
	LONGSHOT       EStatisticType = "longshot"
	PLAYTIME       EStatisticType = "playtime"
	EXP_EARNED     EStatisticType = "exp_earned"
	INVALID        EStatisticType = "invalid"
)
