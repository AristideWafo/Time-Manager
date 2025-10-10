package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Presence struct {
	ID        primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Type      string             `json:"Type" binding:"required" validate:"required"`
	Timestamp time.Time          `json:"timestamp" binding:"required" validate:"required"`
	CreatedAt time.Time          `json:"created_at" binding:"required" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at"`
	User      User               `json:"user" binding:"required" validate:"required"`
}
