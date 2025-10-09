package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Password  string             `json:"password" binding:"required" validate:"required"`
	FirstName string             `json:"first_name" binding:"required" validate:"required"`
	LastName  string             `json:"last_name" binding:"required" validate:"required"`
	Email     string             `json:"email" binding:"required" validate:"required,email"`
	Role      string             `json:"role" binding:"required" validate:"required"`
	Team      string             `json:"team"`
	CreatedAt time.Time          `json:"created_at" binding:"required" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at"`
}
