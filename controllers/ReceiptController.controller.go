// Package controllers ..
package controllers

import (
	"fmt"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreCatchReceipt ..
func StoreCatchReceipt(c *gin.Context) {
	var catchReceipt models.CatchReceipt

	c.ShouldBindJSON(&catchReceipt)

	config.DB.Create(&catchReceipt)

	// Change The Client Balance
	var client models.Clients
	config.DB.Where("ID = ?", catchReceipt.ClientsID).First(&client)

	oldBalance := client.Balance
	newBalance := oldBalance - catchReceipt.Price

	client.Balance = newBalance
	fmt.Println(client.ID)

	config.DB.Save(&client)

	c.JSON(200, gin.H{
		"catchReceipt": catchReceipt,
	})
}

// StorePaymentReceipt ..
func StorePaymentReceipt(c *gin.Context) {
	var paymentReceipt models.PaymentReceipt

	c.ShouldBindJSON(&paymentReceipt)

	config.DB.Create(&paymentReceipt)

	// Change The Client Balance
	var client models.Clients
	config.DB.Where("ID = ?", paymentReceipt.ClientsID).First(&client)

	oldBalance := client.Balance
	newBalance := oldBalance + paymentReceipt.Price

	client.Balance = newBalance

	config.DB.Save(&client)

	c.JSON(200, gin.H{
		"paymentReceipt": paymentReceipt,
	})
}

// StoreExchangeReceipt ..
func StoreExchangeReceipt(c *gin.Context) {
	var exchangeReceipt models.ExchangeReceipt

	c.ShouldBindJSON(&exchangeReceipt)

	config.DB.Create(&exchangeReceipt)

	c.JSON(200, gin.H{
		"exchangeReceipt": exchangeReceipt,
	})
}

// StoreItemsToDelegateType ..
type StoreItemsToDelegateType struct {
	ItemsToDelegate  models.ItemsToDelegate    `json:"itemsToDelegate"`
	DelegatesStorage []models.DelegatesStorage `json:"delegatesStorages"`
}

// StoreItemsToDelegate ..
func StoreItemsToDelegate(c *gin.Context) {
	var data StoreItemsToDelegateType

	c.ShouldBindJSON(&data)

	itemsToDelegate := data.ItemsToDelegate
	delegatesStorage := data.DelegatesStorage

	config.DB.Create(&itemsToDelegate)

	for _, item := range delegatesStorage {
		// Check if DelegateStorageHasHhe item before

		var itemInDelegateStorageCount int64
		config.DB.Model(&models.DelegatesStorage{}).
			Where("delegates_id = ?", itemsToDelegate.DelegatesID).
			Where("item_id = ?", item.ItemID).
			Count(&itemInDelegateStorageCount)
		fmt.Println(itemInDelegateStorageCount)
		if itemInDelegateStorageCount == 0 {
			delegateStorage := models.DelegatesStorage{
				DelegatesID: itemsToDelegate.DelegatesID,
				ItemID:      item.ItemID,
				Qty:         item.Qty,
			}
			config.DB.Create(&delegateStorage)
		} else {
			var delegateStorage models.DelegatesStorage
			config.DB.Where("delegates_id = ?", itemsToDelegate.DelegatesID).
				Where("item_id = ?", item.ItemID).
				First(&delegateStorage)

			newQty := delegateStorage.Qty + item.Qty

			delegateStorage.Qty = newQty

			config.DB.Save(&delegateStorage)
		}

		// Remove Items From Storages
		var itemsStorages models.StoragesItems
		config.DB.Where("item_id = ?", item.ItemID).First(&itemsStorages)

		newQty := itemsStorages.Qty - item.Qty

		itemsStorages.Qty = newQty

		config.DB.Save(&itemsStorages)

	}

}
