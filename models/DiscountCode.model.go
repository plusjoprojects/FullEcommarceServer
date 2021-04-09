// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DiscountCode ..
type DiscountCode struct {
	gorm.Model
	Title string `json:"title"`
	Code  string `json:"code"`
	Value int64  `json:"value"`
}
