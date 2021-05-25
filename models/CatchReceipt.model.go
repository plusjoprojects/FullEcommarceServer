// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// CatchReceipt ..
type CatchReceipt struct {
	ClientsID uint    `json:"clientsID"`
	Date      string  `json:"date"`
	No        string  `json:"no"`
	Price     float64 `json:"price"`
	Note      string  `json:"note"`
	PriceText string  `json:"priceText"`
	Type      uint    `json:"type" gorm:"default:0"`
	CheckNo   string  `json:"checkNo"`
	DueDate   string  `json:"dueDate"`
	Clients   Clients `json:"clients" gorm:"foreignKey:ClientsID;references:ID"`
	gorm.Model
}
