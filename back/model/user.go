package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `json:"_id" bson:"_id,omitempty" binding:"required" validate:"required"`
	Password   string        `json:"Password" bson:"Password" binding:"required" validate:"required"`
	FirstName  string        `json:"FirstName" bson:"FirstName" binding:"required" validate:"required"`
	LastName   string        `json:"LastName" bson:"LastName" binding:"required" validate:"required"`
	Email      string        `json:"Email" binding:"required" bson:"Email" validate:"required,email"`
	Role       string        `json:"Role" binding:"required" bson:"Role" validate:"required"`
	Team       bson.ObjectID `json:"Team" bson:"Team,omitempty"`
	ShiftStart time.Time     `json:"ShiftStart" bson:"ShiftStart,omitempty"`
	ShiftEnd   time.Time     `json:"ShiftEnd" bson:"ShiftEnd,omitempty"`
	CreatedAt  time.Time     `json:"CreatedAt" bson:"CreatedAt" binding:"required" validate:"required"`
	UpdatedAt  time.Time     `json:"UpdatedAt" bson:"UpdatedAt" binding:"required" validate:"required"`
}

func (user *User) Validate() error {
	return ValidateModel(user)
}

func (user *User) ValidatePreSave() error {
	return ValidatePreSave(user)
}

func (user *User) SetID(_id bson.ObjectID) {
	user.ID = _id
}

func (user *User) GetID() bson.ObjectID {
	return user.ID
}

func (user *User) SetCreatedAtAndUpdatedAt() {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
}
