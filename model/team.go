package model

import (
	"time"
)

type Team struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Users     []User    `json:"users"`
}
