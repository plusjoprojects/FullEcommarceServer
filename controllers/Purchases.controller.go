// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddPurchasesType ..
type AddPurchasesType struct {
	Purchases      models.Purchases        `json:"purchases"`
	PurchasesItems []models.PurchasesItems `json:"purchases_items"`
}

// AddPurchases .. Function
func AddPurchases(c *gin.Context) {
	// Get All The Bind Json
	var addPurchase AddPurchasesType
	if err := c.ShouldBindJSON(&addPurchase); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create Purchases Table
	config.DB.Create(&addPurchase.Purchases)

	// check if payment type is 'ذمم'
	if addPurchase.Purchases.BillType == 1 {
		var client models.Clients
		config.DB.Where("ID = ?", addPurchase.Purchases.ClientsID).First(&client)

		lastBalance := client.Balance
		newBalance := lastBalance - addPurchase.Purchases.Net

		config.DB.Model(&models.Clients{}).Where("ID = ?", addPurchase.Purchases.ClientsID).Updates(&models.Clients{
			Balance: newBalance,
		})
	}

	// Create Purchases Items

	for _, item := range addPurchase.PurchasesItems {
		// Add Item To Purchases Items
		PurchaseItem := models.PurchasesItems{
			PurchaseID:    addPurchase.Purchases.ID,
			ItemID:        item.ID,
			Qty:           item.Qty,
			PurchasePrice: item.PurchasePrice,
			Total:         item.Total,
			Storage:       item.Storage,
		}
		config.DB.Create(&PurchaseItem) // Save And End
		// Now Store Item in Storages Items
		// First Check if the item with same scope is avaliable
		var StorageItemsCount int64
		config.DB.Model(models.StoragesItems{}).Where(models.StoragesItems{StorageScope: item.Storage, ItemID: item.ID}).
			Count(&StorageItemsCount)

		// If No Count
		if StorageItemsCount == 0 {
			config.DB.Create(&models.StoragesItems{
				StorageScope: item.Storage,
				ItemID:       item.ID,
				Qty:          item.Qty,
			})
		} else {
			// Get The Last Value Of Qty
			var storageItems models.StoragesItems
			config.DB.Where(models.StoragesItems{StorageScope: item.Storage, ItemID: item.ID}).
				First(&storageItems)

			// The Last value
			lastQty := storageItems.Qty
			newQty := lastQty + item.Qty
			config.DB.Model(models.StoragesItems{}).Where(models.StoragesItems{StorageScope: item.Storage, ItemID: item.ID}).
				Updates(&models.StoragesItems{
					Qty: newQty,
				})
		}

	}

	c.JSON(200, gin.H{
		"message": "Success",
	})

}

// IndexPurchases ..
func IndexPurchases(c *gin.Context) {
	var purchases []models.Purchases
	config.DB.Preload("Company").Preload("PurchasesItems", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Item")
	}).Order("id asc").Find(&purchases)

	c.JSON(200, gin.H{
		"purchases": purchases,
	})
}
