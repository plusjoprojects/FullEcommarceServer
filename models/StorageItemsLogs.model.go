// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// StorageItemsLogs ..
type StorageItemsLogs struct {
	ItemID  uint    `json:"item_id"`
	Actions string  `json:"actions"`
	Qty     float64 `json:"qty"`
	Note    string  `json:"note"`
	gorm.Model
}
