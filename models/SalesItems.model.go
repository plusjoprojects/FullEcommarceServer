// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SalesItems ..
type SalesItems struct {
	SalesID uint    `json:"salesID"`
	ItemID  uint    `json:"item_id"`
	Qty     float64 `json:"qty"`
	Price   float64 `json:"price"`
	Total   float64 `json:"total"`
	Item    Items   `json:"item" gorm:"foreignKey:ItemID;references:ID"`
	gorm.Model
}
