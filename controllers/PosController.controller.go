// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IndexPosItems ..
func IndexPosItems(c *gin.Context) {
	var items []models.Items
	config.DB.Order("id desc").Limit(20).Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

// IndexPosItemsWithCategoriesID ..
func IndexPosItemsWithCategoriesID(c *gin.Context) {
	categoryID := c.Param("id")

	var items []models.Items
	config.DB.Where("categories_id = ?", categoryID).Order("id desc").Limit(20).Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

// SearchPosItems ..
func SearchPosItems(c *gin.Context) {
	text := c.Param("text")

	var items []models.Items
	config.DB.Where("title LIKE ?", "%"+text+"%").
		Limit(20).
		Order("id desc").
		Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})

}

// StorePosOrderType ..
type StorePosOrderType struct {
	Order      models.Orders       `json:"order"`
	OrderItems []models.OrderItems `json:"order_items"`
}

// StorePosOrder ..
func StorePosOrder(c *gin.Context) {
	var data StorePosOrderType
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

// GetBarcodeItem ..
func GetBarcodeItem(c *gin.Context) {
	barcode := c.Param("barcode")

	var item models.Items
	err := config.DB.Where("barcode = ?", barcode).First(&item).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Item Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

// IndexPosTodayOrders ..
func IndexPosTodayOrders(c *gin.Context) {
	// created_at BETWEEN ? AND ?
	startOfToday, endOfToday := vendors.BetwenToday()

	var orders []models.Orders
	config.DB.
		Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Item")
		}).
		Where("created_at BETWEEN ? AND ?", startOfToday, endOfToday).
		Where("type = ?", "Pos").
		Order("id desc").
		Find(&orders)

	c.JSON(200, gin.H{
		"orders": orders,
	})
}
