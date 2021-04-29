// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Delegates ..
type Delegates struct {
	Name            string  `json:"name"`
	Phone           string  `json:"phone"`
	EmergencyNumber string  `json:"emergency_number"`
	Address         string  `json:"address"`
	Birthday        string  `json:"birthday"`
	Tier1           float64 `json:"tier1"`
	Tier2           float64 `json:"tier2"`
	Tier3           float64 `json:"tier3"`
	gorm.Model
}
