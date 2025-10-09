package model

import "time"

type User struct {
	ID        string    `json:"_id" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	Team      string    `json:"team"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
