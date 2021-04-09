// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// TranslationsLocaleItems ..
type TranslationsLocaleItems struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Items ..
type Items struct {
	Title              string                  `json:"title"`
	Description        string                  `json:"description"`
	Image              string                  `json:"image"`
	Price              float64                 `json:"price"`
	Discount           float64                 `json:"discount" gorm:"default:0"`
	New                bool                    `json:"new" gorm:"default:false"`
	Barcode            string                  `json:"barcode"`
	Color              string                  `json:"color"`
	Size               string                  `json:"size"`
	Weight             string                  `json:"weight"`
	Packing            string                  `json:"packing"`
	CategoriesID       uint                    `json:"categories_id"`
	SubCategoriesID    uint                    `json:"sub_categories_id"`
	Categories         Categories              `json:"categories" gorm:"foreignKey:CategoriesID;references:ID"`
	SubCategories      SubCategories           `json:"subCategories" gorm:"foreignKey:SubCategoriesID;references:ID"`
	Translations       []Translations          `json:"translations" gorm:"foreignKey:ForeignKey;references:ID"`
	TranslationsLocale TranslationsLocaleItems `json:"translationsLocale"`
	gorm.Model
}
