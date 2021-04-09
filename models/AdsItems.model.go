// Package models
package models

import (
	"github.com/jinzhu/gorm"
)

// AdsItems ..
type AdsItems struct {
	AdsID  uint  `json:"ads_id"`
	ItemID uint  `json:"item_id"`
	Item   Items `json:"item" gorm:"foreignKey:ItemID;references:ID"`
	gorm.Model
}
