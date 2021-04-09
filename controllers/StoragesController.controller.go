// Package controllers ..
package controllers

import (
	"fmt"
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// IndexStorageItems ..
func IndexStorageItems(c *gin.Context) {
	scope := c.Param("scope")

	var storageItems []models.StoragesItems

	config.DB.Preload("Item").Where("storage_scope = ?", scope).Find(&storageItems)

	c.JSON(200, gin.H{
		"storageItems": storageItems,
	})
}

// MoveItemsFromStorageToStorageModel ..
type MoveItemsFromStorageToStorageModel struct {
	ItemID           uint    `json:"item_id"`
	FromStorageScope string  `json:"from_storage_scope"`
	ToStorageScope   string  `json:"to_storage_scope"`
	Qty              float64 `json:"qty"`
	Note             string  `json:"note"`
}

// MoveFromStorageToStorage ..
func MoveFromStorageToStorage(c *gin.Context) {
	var data MoveItemsFromStorageToStorageModel

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Reduces The Qty From Last Storage ..
	var fromStorageRecord models.StoragesItems
	config.DB.Where("item_id = ?", data.ItemID).
		Where("storage_scope = ?", data.FromStorageScope).
		First(&fromStorageRecord)

	fromStorageRecord.Qty = fromStorageRecord.Qty - data.Qty
	config.DB.Save(&fromStorageRecord)
	// End Reduces The Qty From Last Storage ..

	// Incress Qty To The Next Storage ..
	var newStorageRecordCount int64
	config.DB.Model(&models.StoragesItems{}).Where("item_id = ?", data.ItemID).
		Where("storage_scope = ?", data.ToStorageScope).
		Count(&newStorageRecordCount)
	fmt.Println(newStorageRecordCount)
	if newStorageRecordCount == 0 {
		config.DB.Create(&models.StoragesItems{
			ItemID:       data.ItemID,
			StorageScope: data.ToStorageScope,
			Qty:          data.Qty,
		})
	} else {
		var toStorageRecord models.StoragesItems
		config.DB.Where("item_id = ?", data.ItemID).
			Where("storage_scope = ?", data.ToStorageScope).
			First(&toStorageRecord)
		toStorageRecord.Qty = toStorageRecord.Qty + data.Qty

		config.DB.Save(&toStorageRecord)
	}

	storageItemsLogs := models.StorageItemsLogs{
		ItemID:  data.ItemID,
		Actions: "تم النقل من مستودع إلى أخر",
		Qty:     data.Qty,
		Note:    data.Note,
	}
	config.DB.Create(&storageItemsLogs)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

// ChangeStorageItemQtyModel ..
type ChangeStorageItemQtyModel struct {
	ItemID       uint    `json:"item_id"`
	Qty          float64 `json:"qty"`
	StorageScope string  `json:"storage_scope"`
	Note         string  `json:"note"`
	Type         string  `json:"type"`
}

// ChangeStorageItemQty ..
func ChangeStorageItemQty(c *gin.Context) {
	var data ChangeStorageItemQtyModel

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storageItems models.StoragesItems
	config.DB.Where("item_id = ?", data.ItemID).
		Where("storage_scope = ?", data.StorageScope).
		First(&storageItems)

	var actions string
	if data.Type == "+" {
		storageItems.Qty = storageItems.Qty + data.Qty
		actions = "تم زيادة العدد يدويا"
	} else {
		storageItems.Qty = storageItems.Qty - data.Qty
		actions = "تم تقليل العدد يدويا"
	}

	config.DB.Save(&storageItems)

	storageItemsLogs := models.StorageItemsLogs{
		ItemID:  data.ItemID,
		Actions: actions,
		Qty:     data.Qty,
		Note:    data.Note,
	}
	config.DB.Create(&storageItemsLogs)

	c.JSON(200, gin.H{
		"message": "success",
	})

}
