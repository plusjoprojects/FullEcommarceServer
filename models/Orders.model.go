// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Orders ..
type Orders struct {
	Total         float64      `json:"total"`
	Type          string       `json:"type"`
	AddressID     uint         `json:"address_id"`
	UserID        uint         `json:"user_id"`
	EmployeeID    uint         `json:"employee_id" gorm:"default:0"`
	DiscountID    uint         `json:"discount_id"`
	Note          string       `json:"note"`
	PaymentMethod string       `json:"payment_method"`
	Status        uint         `json:"status" gorm:"default:0"`
	OrderItems    []OrderItems `json:"order_items" gorm:"foreignKey:OrderID;references:ID"`
	gorm.Model
}
