// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// User ..
type User struct {
	gorm.Model
	Name     string   `json:"name"`
	Phone    string   `json:"phone" gorm:"unique"`
	Password string   `json:"password"`
	RolesID  uint     `json:"roles_id"`
	Points   uint     `json:"points" gorm:"default:0"`
	Roles    Roles    `json:"roles" gorm:"foreignKey:RolesID;references:ID"`
	Employee Employee `json:"employee" gorm:"foreignKey:UserID;references:ID"`
}

// Login ...
type Login struct {
	Phone    string `json:"phone" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}
