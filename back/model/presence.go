package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Presence struct {
	ID        bson.ObjectID `json:"_id" bson:"_id,omitempty" binding:"required" validate:"required"`
	Type      string        `json:"Type" binding:"required" validate:"required"`
	Timestamp time.Time     `json:"Timestamp" bson:"Timestamp" binding:"required" validate:"required"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"CreatedAt" binding:"required" validate:"required"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"UpdatedAt" binding:"required" validate:"required"`
	User      bson.ObjectID `json:"User" bson:"User,omitempty" binding:"required" validate:"required"`
}

func (presence *Presence) Validate() error {
	return ValidateModel(presence)
}

func (presence *Presence) ValidatePreSave() error {
	return ValidatePreSave(presence)
}

func (presence *Presence) SetID(_id bson.ObjectID) {
	presence.ID = _id
}

func (presence *Presence) GetID() bson.ObjectID {
	return presence.ID
}

func (presence *Presence) SetCreatedAtAndUpdatedAt() {
	presence.CreatedAt = time.Now()
	presence.UpdatedAt = presence.CreatedAt
}
