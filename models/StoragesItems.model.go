// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// StoragesItems ..
type StoragesItems struct {
	StorageScope string  `json:"storage_scope"`
	ItemID       uint    `json:"item_id"`
	Qty          float64 `json:"qty"`
	gorm.Model
}

type StoragesItemsForApp struct {
	StorageScope string  `json:"storage_scope"`
	ItemID       uint    `json:"item_id"`
	Qty          float64 `json:"qty"`
	gorm.Model
}
