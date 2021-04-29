// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// ItemsToDelegate ..
type ItemsToDelegate struct {
	DelegatesID uint   `json:"delegatesID"`
	Date        string `json:"date"`
	Note        string `json:"note"`
	No          string `json:"no"`
	gorm.Model
}
