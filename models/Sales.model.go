// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Sales ..
type Sales struct {
	ClientsID   uint         `json:"clientsID"`
	DelegatesID uint         `json:"delegatesID"`
	Date        string       `json:"date"`
	No          string       `json:"no"`
	Discount    float64      `json:"discount"`
	Note        string       `json:"note"`
	Total       float64      `json:"total"`
	TaxValue    float64      `json:"taxValue"`
	Net         float64      `json:"net"`
	BillType    uint         `json:"billType"`
	DueDate     string       `json:"due_date"`
	SalesItems  []SalesItems `json:"salesItems" gorm:"foreignKey:SalesID;references:ID"`
	Deletegates Delegates    `json:"delegates" gorm:"foreignKey:DelegatesID;references:ID"`
	Clients     Delegates    `json:"clients" gorm:"foreignKey:ClientsID;references:ID"`
	gorm.Model
}
