// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// TranslationsLocaleCategories ..
type TranslationsLocaleCategories struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Categories ..
type Categories struct {
	Title              string                       `json:"title"`
	Description        string                       `json:"description"`
	Image              string                       `json:"image"`
	Translations       []Translations               `json:"translations" gorm:"foreignKey:ForeignKey;references:ID"`
	TranslationsLocale TranslationsLocaleCategories `json:"translationsLocale" gorm:"foreignKey:ForeignKey;references:ID"`
	gorm.Model
}

// CategoriesWithItems ..
type CategoriesWithItems struct {
	Category Categories `json:"category"`
	Items    []Items    `json:"items"`
}
