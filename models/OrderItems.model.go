// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// OrderItems ..
type OrderItems struct {
	ItemID  uint    `json:"item_id"`
	OrderID uint    `json:"order_id"`
	Qty     int64   `json:"qty"`
	Total   float64 `json:"total"`
	Item    Items   `json:"item" gorm:"foreignKey:ItemID;references:ID"`
	gorm.Model
}
