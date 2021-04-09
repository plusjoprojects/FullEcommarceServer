// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Colors ..
type Colors struct {
	Title string `json:"title"`
	Color string `json:"color"`
	gorm.Model
}
