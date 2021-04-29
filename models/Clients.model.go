// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Clients ..
type Clients struct {
	Name           string  `json:"name"`
	No             string  `json:"no"`
	TaxID          string  `json:"taxID"`
	Address        string  `json:"address"`
	Phone          string  `json:"phone"`
	Fax            string  `json:"fax"`
	Boss           string  `json:"boss"`
	BossPhone      string  `json:"bossPhone"`
	StarterBalance float64 `json:"starterBalance"`
	Limit          float64 `json:"limit"`
	Balance        float64 `json:"balance"`
	Type           string  `json:"type"`
	PriceType      uint    `json:"price_type" gorm:"default:0"`
	gorm.Model
}
