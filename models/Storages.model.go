// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Storages ..
type Storages struct {
	Title string `json:"title"`
	Scope string `json:"scope"`
	gorm.Model
}
