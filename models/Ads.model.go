// Package models
package models

import (
	"github.com/jinzhu/gorm"
)

// Ads ..
type Ads struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Active      bool   `json:"active"`
	Image       string `json:"image"`
	gorm.Model
}
