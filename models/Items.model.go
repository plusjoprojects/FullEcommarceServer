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

// WithStorageCounts ..
func WithStorageCounts() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(&Items{}).Select("storages_items.qty,items.*").Joins("LEFT JOIN storages_items ON storages_items.item_id = items.ID AND storages_items.storage_scope = 'All'")
	}
}

// Items ..
type Items struct {
	Title              string                  `json:"title"`
	Description        string                  `json:"description"`
	Image              string                  `json:"image"`
	Price              float64                 `json:"price"`
	Discount           float64                 `json:"discount" gorm:"default:0"`
	OriginalPrice      float64                 `json:"original_price" gorm:"default:0"`
	New                bool                    `json:"new" gorm:"default:false"`
	Barcode            string                  `json:"barcode"`
	Color              string                  `json:"color"`
	Size               string                  `json:"size"`
	Weight             string                  `json:"weight"`
	Packing            string                  `json:"packing"`
	CategoriesID       uint                    `json:"categories_id"`
	SubCategoriesID    uint                    `json:"sub_categories_id"`
	Section            string                  `json:"section"`
	Stand              string                  `json:"stand"`
	Row                string                  `json:"row"`
	Part               string                  `json:"part"`
	BuyPrice           float64                 `json:"buyPrice"`
	JomlaPrice         float64                 `json:"jomlaPrice"`
	JomltAljomlaPrice  float64                 `json:"jomltAljomlaPrice"`
	Reorder            float64                 `json:"reorder"`
	PayOnApp           bool                    `json:"payOnApp" gorm:"default:1"`
	StoragesItems      StoragesItems           `json:"storagesItems" gorm:"foreignKey:ItemID;references:ID"`
	Categories         Categories              `json:"categories" gorm:"foreignKey:CategoriesID;references:ID"`
	SubCategories      SubCategories           `json:"subCategories" gorm:"foreignKey:SubCategoriesID;references:ID"`
	Translations       []Translations          `json:"translations" gorm:"foreignKey:ForeignKey;references:ID"`
	TranslationsLocale TranslationsLocaleItems `json:"translationsLocale"`
	gorm.Model
}
