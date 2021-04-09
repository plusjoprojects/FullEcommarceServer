// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Companies ..
type Companies struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	gorm.Model
}
