package model

type Location struct {
	World string  `json:"world,omitempty" bson:"world,omitempty"`
	X     float32 `json:"x" bson:"x,omitempty,truncated"`
	Y     float32 `json:"y" bson:"y,omitempty,truncate"`
	Z     float32 `json:"z" bson:"z,omitempty,truncate"`
}
