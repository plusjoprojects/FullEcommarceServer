// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DelegatesStorage ..
type DelegatesStorage struct {
	DelegatesID uint    `json:"delegatesID"`
	ItemID      uint    `json:"ItemID"`
	Qty         float64 `json:"qty"`
	gorm.Model
}
