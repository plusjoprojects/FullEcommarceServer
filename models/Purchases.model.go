// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Purchases ..
type Purchases struct {
	CompanyID      uint             `json:"company_id"`
	Date           string           `json:"date"`
	No             string           `json:"no"`
	Discount       float64          `json:"discount"`
	Note           string           `json:"note"`
	Total          float64          `json:"total"`
	Net            float64          `json:"net"`
	Company        Companies        `json:"company" gorm:"foreignKey:CompanyID;references:ID"`
	PurchasesItems []PurchasesItems `json:"purchases_items" gorm:"foreignKey:PurchaseID;references:ID"`
	gorm.Model
}
