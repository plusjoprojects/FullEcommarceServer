// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Employee ..
type Employee struct {
	gorm.Model
	UserID          uint    `json:"user_id"`
	Birthday        string  `json:"birthday"`
	Address         string  `json:"address"`
	Salary          float64 `json:"salary"`
	EmergencyNumber string  `json:"emergency_number"`
}
