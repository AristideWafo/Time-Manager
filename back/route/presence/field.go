package presence

import (
	"time"
)

type PresenceOutput struct {
	Type      string    `json:"Type" binding:"required" validate:"required"`
	Timestamp time.Time `json:"Timestamp" bson:"Timestamp" binding:"required" validate:"required"`
}

type CreatePresenceInput struct {
	Type      string    `json:"Type" binding:"required" validate:"required"`
	Timestamp time.Time `json:"Timestamp" bson:"Timestamp" binding:"required" validate:"required"`
	User      string    `json:"User" bson:"User,omitempty" binding:"required" validate:"required"`
}
