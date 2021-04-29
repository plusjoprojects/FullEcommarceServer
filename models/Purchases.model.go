// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Purchases ..
type Purchases struct {
	ClientsID      uint             `json:"clientsID"`
	Date           string           `json:"date"`
	No             string           `json:"no"`
	Discount       float64          `json:"discount"`
	Note           string           `json:"note"`
	Total          float64          `json:"total"`
	TaxValue       float64          `json:"taxValue"`
	Net            float64          `json:"net"`
	BillType       uint             `json:"billType"`
	Clients        Clients          `json:"clients" gorm:"foreignKey:ClientsID;references:ID"`
	PurchasesItems []PurchasesItems `json:"purchases_items" gorm:"foreignKey:PurchaseID;references:ID"`
	gorm.Model
}
