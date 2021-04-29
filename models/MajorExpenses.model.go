// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// MajorExpenses ..
type MajorExpenses struct {
	Title string `json:"title"`
	gorm.Model
}
