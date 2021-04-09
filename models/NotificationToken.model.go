// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// NotificationToken ..
type NotificationToken struct {
	UserID             uint   `json:"user_id" gorm:"default:0"`
	Token              string `json:"token"`
	AllNotification    bool   `json:"all_notification" gorm:"default:true"`
	OrdersNotification bool   `json:"orders_notification" gorm:"default:true"`
	gorm.Model
}
