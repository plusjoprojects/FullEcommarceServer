// Package models
package models

import (
	"github.com/jinzhu/gorm"
)

// Addresses ..
type Addresses struct {
	City    string  `json:"city"`
	Area    string  `json:"area"`
	Address string  `json:"address"`
	Note    string  `json:"note"`
	Lat     float64 `json:"lat" gorm:"default:0"`
	Long    float64 `json:"long" gorm:"default:0"`
	UserID  uint    `json:"user_id"`
	gorm.Model
}
