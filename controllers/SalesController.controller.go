// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddSalesType ..
type AddSalesType struct {
	Sales      models.Sales        `json:"sales"`
	SalesItems []models.SalesItems `json:"sales_items"`
}

// AddSale ..
func AddSale(c *gin.Context) {
	var addSales AddSalesType
	c.ShouldBindJSON(&addSales)

	sales := addSales.Sales
	salesItems := addSales.SalesItems

	config.DB.Create(&sales)

	// if billType = 'ذمم' make balance change
	if sales.BillType == 1 {
		var client models.Clients
		config.DB.Where("ID = ?", sales.ClientsID).First(&client)

		lastBalance := client.Balance
		newBalance := lastBalance + sales.Net

		config.DB.Model(&models.Clients{}).Where("ID = ?", sales.ClientsID).Updates(&models.Clients{
			Balance: newBalance,
		})
	}

	for _, item := range salesItems {
		// Add Item To Purchases Items
		salesItems := models.SalesItems{
			SalesID: sales.ID,
			ItemID:  item.ID,
			Qty:     item.Qty,
			Price:   item.Price,
			Total:   item.Total,
		}
		config.DB.Create(&salesItems) // Save And End

		// Now Store Item in Storages Items
		// Get The Last Value Of Qty
		var storageItems models.StoragesItems
		config.DB.Where(models.StoragesItems{StorageScope: "All", ItemID: item.ID}).
			First(&storageItems)
		// The Last value
		lastQty := storageItems.Qty
		newQty := lastQty - item.Qty
		config.DB.Model(models.StoragesItems{}).Where(models.StoragesItems{StorageScope: "All", ItemID: item.ID}).
			Updates(&models.StoragesItems{
				Qty: newQty,
			})
	}
}

// GetSalesLastID ..
func GetSalesLastID(c *gin.Context) {

	var sale models.Sales

	config.DB.Last(&sale)

	ID := sale.ID

	c.JSON(200, gin.H{
		"ID": ID,
	})
}

// ShowSale ..
func ShowSale(c *gin.Context) {
	ID := c.Param("id")

	var sale models.Sales

	config.DB.Preload("SalesItems", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Item")
	}).Preload("Clients").Where("id = ?", ID).Find(&sale)

	c.JSON(200, gin.H{
		"sale": sale,
	})

}
