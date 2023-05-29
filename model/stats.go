package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatHolder struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID           string             `json:"uuid" bson:"uuid"`
	Map            uint8              `json:"map" bson:"map"`
	Kills          uint64             `json:"kills" bson:"KILL,omitempty"`
	Deaths         uint64             `json:"deaths" bson:"DEATH,omitempty"`
	Longshot       uint64             `json:"longshot" bson:"LONGSHOT,omitempty"`
	EventCaptures  uint64             `json:"events" bson:"EVENT_CAPTURES,omitempty"`
	ExpEarned      uint64             `json:"exp" bson:"EXP_EARNED,omitempty"`
	Playtime       uint64             `json:"playtime,omitempty" bson:"PLAYTIME,omitempty"`
	MinedDiamonds  uint64             `json:"mined_diamonds,omitempty" bson:"MINED_DIAMONDS,omitempty"`
	MinedNetherite uint64             `json:"mined_netherite,omitempty" bson:"MINED_NETHERITE,omitempty"`
}

type EStatisticType string

const (
	KILL            EStatisticType = "kill"
	DEATH           EStatisticType = "death"
	EVENT_CAPTURES  EStatisticType = "event_captures"
	LONGSHOT        EStatisticType = "longshot"
	PLAYTIME        EStatisticType = "playtime"
	EXP_EARNED      EStatisticType = "exp_earned"
	MINED_DIAMONDS  EStatisticType = "mined_diamonds"
	MINED_NETHERITE EStatisticType = "mined_netherite"
	INVALID         EStatisticType = "invalid"
)
