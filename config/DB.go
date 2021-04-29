// Package config ...
package config

import (
	"server/models"

	"github.com/jinzhu/gorm"
	// Connect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// models
)

// SetupDB ...

// DB ..
var DB *gorm.DB

// SetupDB ..
func SetupDB() {
	database, err := gorm.Open("mysql", "root:00962s00962!S@tcp(127.0.0.1:3306)/ecommarce?charset=utf8mb4&parseTime=True&loc=Local")

	// If Error in Connect
	if err != nil {
		panic(err)
	}
	// User Models Setup
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Employee{})
	database.AutoMigrate(&models.Delegates{})
	database.AutoMigrate(&models.DelegatesStorage{})
	database.AutoMigrate(&models.Clients{})
	database.AutoMigrate(&models.AuthClients{})
	database.AutoMigrate(&models.AuthTokens{})
	database.AutoMigrate(&models.Roles{})
	database.AutoMigrate(&models.Categories{})
	database.AutoMigrate(&models.SubCategories{})
	database.AutoMigrate(&models.Colors{})
	database.AutoMigrate(&models.Items{})
	database.AutoMigrate(&models.ItemsToDelegate{})
	database.AutoMigrate(&models.Companies{})
	database.AutoMigrate(&models.Storages{})
	database.AutoMigrate(&models.Purchases{})
	database.AutoMigrate(&models.Sales{})
	database.AutoMigrate(&models.SalesItems{})
	database.AutoMigrate(&models.CatchReceipt{})
	database.AutoMigrate(&models.PaymentReceipt{})
	database.AutoMigrate(&models.ExchangeReceipt{})
	database.AutoMigrate(&models.MajorExpenses{})
	database.AutoMigrate(&models.PurchasesItems{})
	database.AutoMigrate(&models.StoragesItems{})
	database.AutoMigrate(&models.StorageItemsLogs{})
	database.AutoMigrate(&models.Addresses{})
	database.AutoMigrate(&models.Orders{})
	database.AutoMigrate(&models.OrderItems{})
	database.AutoMigrate(&models.Ads{})
	database.AutoMigrate(&models.AdsItems{})
	database.AutoMigrate(&models.NotificationToken{})
	database.AutoMigrate(&models.DiscountCode{})
	database.AutoMigrate(&models.Translations{})
	// Confirm the DB variables
	DB = database
}
