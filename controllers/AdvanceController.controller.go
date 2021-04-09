// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// IndexDiscountCodes ..
func IndexDiscountCodes(c *gin.Context) {
	var discountCodes []models.DiscountCode
	config.DB.Find(&discountCodes)

	c.JSON(200, gin.H{
		"discountCodes": discountCodes,
	})
}

// StoreDiscountCode ..
func StoreDiscountCode(c *gin.Context) {
	var discountCode models.DiscountCode
	c.ShouldBindJSON(&discountCode)
	config.DB.Create(&discountCode)

	c.JSON(200, gin.H{
		"discountCode": discountCode,
	})
}

// UpdateDiscountCode ..
func UpdateDiscountCode(c *gin.Context) {
	var discountCode models.DiscountCode
	c.ShouldBindJSON(&discountCode)
	config.DB.Model(&discountCode).Update(&discountCode)

	var discountCodes []models.DiscountCode
	config.DB.Find(&discountCodes)

	c.JSON(200, gin.H{
		"discountCodes": discountCodes,
	})
}

// DestroyDiscountCode ..
func DestroyDiscountCode(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.DiscountCode{}, id)

	var discountCodes []models.DiscountCode
	config.DB.Find(&discountCodes)
	c.JSON(200, gin.H{
		"discountCodes": discountCodes,
	})
}

// IndexAds ..
func IndexAds(c *gin.Context) {
	var ads []models.Ads
	config.DB.Find(&ads)
	c.JSON(200, gin.H{
		"ads": ads,
	})
}

// StoreAds ..
func StoreAds(c *gin.Context) {
	var ad models.Ads
	c.ShouldBindJSON(&ad)
	config.DB.Create(&ad)
	c.JSON(200, gin.H{
		"ad": ad,
	})
}

// UpdateAds ..
func UpdateAds(c *gin.Context) {
	var ad models.Ads
	c.ShouldBindJSON(&ad)
	config.DB.Model(&ad).Update(&ad)
	var ads []models.Ads
	config.DB.Find(&ads)
	c.JSON(200, gin.H{
		"ads": ads,
	})
}

// RemoveAds ..
func RemoveAds(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Ads{}, id)

	var ads []models.Ads
	config.DB.Find(&ads)
	c.JSON(200, gin.H{
		"ads": ads,
	})
}

// StoreAdItemModel ..
type StoreAdItemModel struct {
	AdID    uint   `json:"ad_id"`
	ItemIDS []uint `json:"item_ids"`
}

// StoreAdItem ..
func StoreAdItem(c *gin.Context) {
	var adItem StoreAdItemModel
	c.ShouldBindJSON(&adItem)

	itemIDS := adItem.ItemIDS
	for _, id := range itemIDS {
		config.DB.Create(&models.AdsItems{
			AdsID:  adItem.AdID,
			ItemID: id,
		})
	}

	var adItems []models.AdsItems

	config.DB.Preload("Item").Where("ad_id = ?", adItem.AdID).Find(&adItems)

	c.JSON(200, gin.H{
		"adItems": adItems,
	})
}

// IndexAdItems ..
func IndexAdItems(c *gin.Context) {
	id := c.Param("id")
	var adItems []models.AdsItems
	config.DB.Preload("Item", func(db *gorm.DB) *gorm.DB {
		return db.Scopes(models.WithTranslation("items"))
	}).Where("ads_id = ?", id).Find(&adItems)
	c.JSON(200, gin.H{
		"adItems": adItems,
	})
}

// RemoveAdItem ..
func RemoveAdItem(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.AdsItems{}, id)

	c.JSON(200, gin.H{
		"message": "Deleted",
	})
}

type GlobalSendNotificationType struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Type  string `json:"type"`
	ID    uint   `json:"ID"`
}

// GlobalSendNotification ..
func GlobalSendNotification(c *gin.Context) {
	var notification GlobalSendNotificationType
	c.ShouldBindJSON(&notification)

	var notificationTokens []models.NotificationToken
	config.DB.Find(&notificationTokens)

	var tokens []expo.ExponentPushToken

	for _, notificationToken := range notificationTokens {
		token, _ := expo.NewExponentPushToken(notificationToken.Token)
		tokens = append(tokens, token)
	}

	var IDString string = strconv.FormatUint(uint64(notification.ID), 10)
	data := vendors.NotificationData{
		Type: notification.Type,
		ID:   IDString,
	}

	message := vendors.NotificationMessage{
		Title: notification.Title,
		Body:  notification.Body,
	}

	vendors.SendNotification(tokens, message, data)

	c.JSON(200, gin.H{
		"message": "success",
	})

}
