package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"` // omit password from JSON responses
	Role     string `json:"role" gorm:"not null"` // "doctor" or "receptionist"
}
