// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// PaymentReceipt ..
type PaymentReceipt struct {
	ClientsID uint    `json:"clientsID"`
	Date      string  `json:"date"`
	No        string  `json:"no"`
	Price     float64 `json:"price"`
	Note      string  `json:"note"`
	gorm.Model
}
