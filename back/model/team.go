package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Team struct {
	ID        bson.ObjectID `json:"_id" bson:"_id,omitempty" binding:"required" validate:"required"`
	Name      string        `json:"Name" bson:"Name" binding:"required" validate:"required"`
	CreatedAt time.Time     `json:"CreatedAt" bson:"CreatedAt" binding:"required" validate:"required"`
	UpdatedAt time.Time     `json:"UpdatedAt" bson:"UpdatedAt"`
}

func (team *Team) Validate() error {
	return ValidateModel(team)
}

func (team *Team) ValidatePreSave() error {
	return ValidatePreSave(team)
}

func (team *Team) SetID(_id bson.ObjectID) {
	team.ID = _id
}

func (team *Team) GetID() bson.ObjectID {
	return team.ID
}

func (team *Team) SetCreatedAtAndUpdatedAt() {
	team.CreatedAt = time.Now()
	team.UpdatedAt = team.CreatedAt
}
