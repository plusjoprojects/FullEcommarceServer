// Package models ..
package models

import "github.com/jinzhu/gorm"

type Translations struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	ForeignKey uint   `json:"foreign_key"`
	Locale     string `json:"locale"`
	Value      string `json:"value"`
	*gorm.Model
}

func WithTranslation(table_name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		scope := "table_name = " + "'" + table_name + "'"
		return db.Preload("Translations", scope)
	}
}
