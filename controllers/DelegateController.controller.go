// Package controllers ...
package controllers

import (
	"fmt"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ShowDelegate ..
func ShowDelegate(c *gin.Context) {
	ID := c.Param("ID")
	var delegate models.Delegates
	config.DB.Preload("DelegatesStorage", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Item")
	}).Where("id = ?", ID).First(&delegate)
	c.JSON(200, gin.H{
		"delegate": delegate,
	})
}

// ExportItemFromDelegateToStorageType ..
type ExportItemFromDelegateToStorageType struct {
	ItemID     int     `json:"itemID"`
	Qty        float64 `json:"qty"`
	DelegateID uint    `json:"delegateID"`
}

// ExportItemFromDelegateToStorage ..
func ExportItemFromDelegateToStorage(c *gin.Context) {
	var data ExportItemFromDelegateToStorageType

	c.ShouldBindJSON(&data)

	ItemID := data.ItemID
	Qty := data.Qty
	DelegateID := data.DelegateID
	// Out Item From DelegateStorage
	var delegateStorage models.DelegatesStorage
	err := config.DB.Where("delegates_id = ?", DelegateID).Where("item_id = ?", ItemID).Find(&delegateStorage).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"Error": err,
		})
		return
	}

	delegateStorageQty := delegateStorage.Qty
	if Qty == delegateStorageQty {
		config.DB.Delete(&models.DelegatesStorage{}, delegateStorage.ID)
	} else {
		newQty := delegateStorageQty - Qty

		delegateStorage.Qty = newQty

		config.DB.Save(&delegateStorage)
	}

	// Insert Item To StorageItem
	var storageItem models.StoragesItems
	GetItemErr := config.DB.Where("item_id = ?", ItemID).First(&storageItem).Error
	if GetItemErr != nil {
		c.JSON(500, gin.H{
			"Error": "GetItemErr",
		})
		return
	}
	storageItemNewQty := storageItem.Qty + Qty
	storageItem.Qty = storageItemNewQty
	config.DB.Save(&storageItem)

	c.JSON(200, gin.H{
		"message": "Success",
	})

}
