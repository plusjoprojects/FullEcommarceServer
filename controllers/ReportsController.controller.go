// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IndexReport ..
func IndexReport(c *gin.Context) {
	reportType := c.Param("reportType")

	if reportType == "تقرير اليوم" {
		startOfToday, endOfToday := vendors.BetwenToday()

		var orders []models.Orders
		config.DB.
			Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Item")
			}).
			Where("created_at BETWEEN ? AND ?", startOfToday, endOfToday).
			Order("id desc").
			Find(&orders)

		c.JSON(200, gin.H{
			"orders": orders,
		})
		return
	}

	if reportType == "تقرير الأمس" {
		startOfYesterday, endOfYesterday := vendors.BetwenYesterDay()
		var orders []models.Orders
		config.DB.
			Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Item")
			}).
			Where("created_at BETWEEN ? AND ?", startOfYesterday, endOfYesterday).
			Order("id desc").
			Find(&orders)

		c.JSON(200, gin.H{
			"orders": orders,
		})
		return
	}

	if reportType == "تقرير أخر 7 ايام" {
		start, end := vendors.BetwenLastSevenDay()
		var orders []models.Orders
		config.DB.
			Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Item")
			}).
			Where("created_at BETWEEN ? AND ?", start, end).
			Order("id desc").
			Find(&orders)

		c.JSON(200, gin.H{
			"orders": orders,
		})
		return
	}

	if reportType == "تقرير الشهر" {
		start, end := vendors.BetweenThisMonth()
		var orders []models.Orders
		config.DB.
			Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Item")
			}).
			Where("created_at BETWEEN ? AND ?", start, end).
			Order("id desc").
			Find(&orders)

		c.JSON(200, gin.H{
			"orders": orders,
		})
		return
	}
}

// DatesInput ..
type DatesInput struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// ReportWithDates ..
func ReportWithDates(c *gin.Context) {
	var datesInput DatesInput
	if err := c.ShouldBindJSON(&datesInput); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	start, end := vendors.BetweenDates(datesInput.StartDate, datesInput.EndDate)

	var orders []models.Orders
	config.DB.
		Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Item")
		}).
		Where("created_at BETWEEN ? AND ?", start, end).
		Order("id desc").
		Find(&orders)

	c.JSON(200, gin.H{
		"orders": orders,
	})
	return
}
