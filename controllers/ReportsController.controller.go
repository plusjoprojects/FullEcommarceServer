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

// ReportPurchasesClients ..
func ReportPurchasesClients(c *gin.Context) {
	ID := c.Param("id")

	var clients []models.Clients
	if ID == "All" {
		config.DB.
			Where("type = ?", "مورد").
			Find(&clients)
	} else {
		config.DB.
			Where("ID = ?", ID).
			Where("type = ?", "مورد").
			Find(&clients)
	}

	c.JSON(200, gin.H{
		"clients": clients,
	})
}

// ReportSalesClients ..
func ReportSalesClients(c *gin.Context) {
	ID := c.Param("id")

	var clients []models.Clients
	if ID == "All" {
		config.DB.
			Where("type = ?", "زبون").
			Find(&clients)
	} else {
		config.DB.
			Where("ID = ?", ID).
			Where("type = ?", "زبون").
			Find(&clients)
	}

	c.JSON(200, gin.H{
		"clients": clients,
	})
}

// ReportsPurchases ..
func ReportsPurchases(c *gin.Context) {
	ID := c.Param("id")

	var fromToDate models.DateReport
	c.ShouldBindJSON(&fromToDate)

	start, end := vendors.BetweenDates(fromToDate.From, fromToDate.To)

	var receipts []models.Purchases
	if ID == "All" {
		config.DB.
			Where("created_at BETWEEN ? AND ?", start, end).
			Preload("Clients").
			Order("id desc").Find(&receipts)
	} else {
		config.DB.
			Where("created_at BETWEEN ? AND ?", start, end).
			Where("clients_id = ?", ID).
			Preload("Clients").
			Order("id desc").Find(&receipts)
	}

	c.JSON(200, gin.H{
		"receipts": receipts,
	})
}

// ReportsSales ..
func ReportsSales(c *gin.Context) {
	ID := c.Param("id")

	var fromToDate models.DateReport
	c.ShouldBindJSON(&fromToDate)

	start, end := vendors.BetweenDates(fromToDate.From, fromToDate.To)

	var receipts []models.Sales
	if ID == "All" {
		config.DB.
			Where("created_at BETWEEN ? AND ?", start, end).
			Preload("Clients").
			Preload("Delegates").
			Order("id desc").Find(&receipts)
	} else {
		config.DB.
			Where("created_at BETWEEN ? AND ?", start, end).
			Where("clients_id = ?", ID).
			Preload("Delegates").
			Preload("Clients").
			Order("id desc").Find(&receipts)
	}

	c.JSON(200, gin.H{
		"receipts": receipts,
	})
}
