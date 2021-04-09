// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// PurchasesItems ..
type PurchasesItems struct {
	PurchaseID    uint    `json:"purchase_id"`
	ItemID        uint    `json:"item_id"`
	Qty           float64 `json:"qty"`
	PurchasePrice float64 `json:"purchase_price"`
	Tax           float64 `json:"tax"`
	Total         float64 `json:"total"`
	Storage       string  `json:"storage"`
	Item          Items   `json:"item" gorm:"foreignKey:ItemID;references:ID"`
	gorm.Model
}
