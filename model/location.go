package model

type Location struct {
	World string  `json:"world" bson:"world"`
	X     float32 `json:"x" bson:"x,truncate"`
	Y     float32 `json:"y" bson:"y,truncate"`
	Z     float32 `json:"z" bson:"z,truncate"`
}
