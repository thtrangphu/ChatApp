package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	DOB      time.Time `json:"dob"`
	IsOnline bool      `json:"isOnline"`
}
