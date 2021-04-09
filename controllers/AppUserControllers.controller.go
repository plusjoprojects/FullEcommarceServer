// Package controllers ..
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreUserAddress ..
func StoreUserAddress(c *gin.Context) {
	var address models.Addresses
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"address": address,
	})
}

// DestroyUserAddress ..
func DestroyUserAddress(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Addresses{}, ID)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

// IndexUserAddresses ..
func IndexUserAddresses(c *gin.Context) {
	ID := c.Param("id")
	var addresses []models.Addresses
	config.DB.Where("user_id = ?", ID).Find(&addresses)
	c.JSON(200, gin.H{
		"addresses": addresses,
	})
}

// StoreAppOrderType ..
type StoreAppOrderType struct {
	Order      models.Orders       `json:"order"`
	OrderItems []models.OrderItems `json:"order_items"`
}

// CheckDiscountCodeValid ..
func CheckDiscountCodeValid(c *gin.Context) {
	code := c.Param("code")
	userID := c.Param("user_id")

	var codeChecker models.DiscountCode
	err := config.DB.Where("code = ?", code).Find(&codeChecker).Error
	if err != nil {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "No Record Found",
		})
		return
	}

	var useCodeCount int64
	config.DB.Model(&models.Orders{}).Where("user_id = ?", userID).
		Where("discount_id = ?", codeChecker.ID).
		Count(&useCodeCount)

	if useCodeCount == 0 {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "Code Valid",
			"code":    codeChecker,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"status":  2,
			"message": "Code Used",
		})
	}

}

// StoreAppOrder ..
func StoreAppOrder(c *gin.Context) {
	var data StoreAppOrderType
	c.ShouldBindJSON(&data)

	order := data.Order
	config.DB.Create(&order)

	for _, item := range data.OrderItems {
		item.OrderID = order.ID
		config.DB.Create(&item)
	}

	c.JSON(200, gin.H{
		"order": order,
	})
}

// AppShowOrder ..
func AppShowOrder(c *gin.Context) {
	orderID := c.Param("id")

	var order models.Orders
	var orderItems []models.OrderItems

	config.DB.Where("id = ?", orderID).First(&order)
	config.DB.Preload("Item").Where("order_id = ?", orderID).Find(&orderItems)

	c.JSON(200, gin.H{
		"order":      order,
		"orderItems": orderItems,
	})

}

// AppSearch ..
func AppSearch(c *gin.Context) {
	text := c.Param("text")

	var items []models.Items
	config.DB.Scopes(models.WithTranslation("items")).Where("title LIKE ?", "%"+text+"%").Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

// AppIndexItemsWithIDSModel ..
type AppIndexItemsWithIDSModel struct {
	IDS []uint `json:"ids"`
}

// AppIndexItemsWithIDS ..
func AppIndexItemsWithIDS(c *gin.Context) {
	var data AppIndexItemsWithIDSModel
	c.ShouldBindJSON(&data)

	var items []models.Items
	config.DB.Scopes(models.WithTranslation("items")).Find(&items, data.IDS)

	c.JSON(200, gin.H{
		"items": items,
	})
}

// StoreNotificationToken ...
func StoreNotificationToken(c *gin.Context) {
	// Register var on model
	var notificationToken models.NotificationToken
	c.ShouldBindJSON(&notificationToken) // Bind JSON

	// Check count in the storage
	var tokenCount int64
	config.DB.Model(&models.NotificationToken{}).Where("token = ?", notificationToken.Token).Count(&tokenCount)

	// Check if have user with it or not
	if notificationToken.UserID == 0 {
		if tokenCount == 0 {
			config.DB.Create(&notificationToken)
		}
	} else {

		if tokenCount == 0 {
			config.DB.Create(&notificationToken)
		} else {
			config.DB.Model(&models.NotificationToken{}).Where("token = ?", notificationToken.Token).Updates(&models.NotificationToken{
				UserID: notificationToken.UserID,
			})
		}
	}

	var userNotificationToken models.NotificationToken
	config.DB.Where("token = ?", notificationToken.Token).Find(&userNotificationToken)

	c.JSON(200, gin.H{
		"message":           "Success",
		"notificationToken": userNotificationToken,
	})
}

// UpdateNotificationTokenDetails ..
func UpdateNotificationTokenDetails(c *gin.Context) {
	var notificationToken models.NotificationToken
	if err := c.ShouldBindJSON(&notificationToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&notificationToken).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"notificationToken": notificationToken,
	})
}
