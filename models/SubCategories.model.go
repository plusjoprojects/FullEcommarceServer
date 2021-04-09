// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SubCategories ..
type SubCategories struct {
	Title              string                       `json:"title"`
	Description        string                       `json:"description"`
	Image              string                       `json:"image"`
	CategoriesID       uint                         `json:"categories_id"`
	Categories         Categories                   `json:"categories" gorm:"foreignKey:CategoriesID;references:ID"`
	Translations       []Translations               `json:"translations" gorm:"foreignKey:ForeignKey;references:ID"`
	TranslationsLocale TranslationsLocaleCategories `json:"translationsLocale" gorm:"foreignKey:ForeignKey;references:ID"`
	gorm.Model
}
