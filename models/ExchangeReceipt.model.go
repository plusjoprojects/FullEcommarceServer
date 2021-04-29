// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// ExchangeReceipt ..
type ExchangeReceipt struct {
	MajorExpensesID uint    `json:"majorExpensesID"`
	Date            string  `json:"date"`
	No              string  `json:"no"`
	Price           float64 `json:"price"`
	Note            string  `json:"note"`
	gorm.Model
}
