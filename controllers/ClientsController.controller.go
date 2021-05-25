// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// StoreClient ..
func StoreClient(c *gin.Context) {
	var client models.Clients

	c.ShouldBindJSON(&client)

	config.DB.Create(&client)

	c.JSON(200, gin.H{
		"client": client,
	})
}

// IndexClients ..
func IndexClients(c *gin.Context) {
	var clients []models.Clients
	config.DB.Find(&clients)

	c.JSON(200, gin.H{
		"clients": clients,
	})
}

// UpdateClient ..
func UpdateClient(c *gin.Context) {
	var client models.Clients

	c.ShouldBindJSON(&client)

	config.DB.Model(&client).Update(&client)

	c.JSON(200, gin.H{
		"client": client,
	})
}

// DestroyClient ..
func DestroyClient(c *gin.Context) {
	ID := c.Param("id")

	config.DB.Delete(&models.Clients{}, ID)

	var clients []models.Clients

	config.DB.Find(&clients)

	c.JSON(200, gin.H{
		"clients": clients,
	})
}

// IndexClientsForPurchase ..
func IndexClientsForPurchase(c *gin.Context) {
	var clients []models.Clients

	config.DB.Where("type = ?", "مورد").Find(&clients)

	c.JSON(200, gin.H{
		"clients": clients,
	})
}

// IndexClientsForSales ..
func IndexClientsForSales(c *gin.Context) {
	var clients []models.Clients

	config.DB.Where("type = ?", "زبون").Find(&clients)

	c.JSON(200, gin.H{
		"clients": clients,
	})
}

// ShowClient ..
func ShowClient(c *gin.Context) {
	ID := c.Param("id")
	var client models.Clients
	config.DB.Where("id = ?", ID).First(&client)

	// Get Sales Not Payment Completed
	var notPaymentSales []models.Sales

	config.DB.Preload("Clients").Preload("SalesItems", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Item")
	}).Where("clients_id = ?", ID).
		Where("payment_collected = ?", 0).
		Find(&notPaymentSales)

	c.JSON(200, gin.H{
		"client":          client,
		"notPaymentSales": notPaymentSales,
	})
}
