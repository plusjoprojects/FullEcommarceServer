// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
)

// StoreMajorExpenses ..
func StoreMajorExpenses(c *gin.Context) {
	var majorExpenses models.MajorExpenses

	c.ShouldBindJSON(&majorExpenses)

	config.DB.Create(&majorExpenses)

	c.JSON(200, gin.H{
		"majorExpenses": majorExpenses,
	})
}

// IndexMajorExpenses ..
func IndexMajorExpenses(c *gin.Context) {
	var majorExpenses []models.MajorExpenses

	config.DB.Find(&majorExpenses)

	c.JSON(200, gin.H{
		"majorExpenses": majorExpenses,
	})
}

// UpdateMajorExpenses ..
func UpdateMajorExpenses(c *gin.Context) {
	var majorExpenses models.MajorExpenses

	c.ShouldBindJSON(&majorExpenses)

	config.DB.Model(&majorExpenses).Update(&majorExpenses)

	var majorExpensesList []models.MajorExpenses
	config.DB.Find(&majorExpenses)

	c.JSON(200, gin.H{
		"majorExpenses": majorExpensesList,
	})
}

// DestroyMajorExpenses ..
func DestroyMajorExpenses(c *gin.Context) {
	ID := c.Param("id")

	config.DB.Delete(&models.MajorExpenses{}, ID)

	var majorExpenses []models.MajorExpenses

	config.DB.Find(&majorExpenses)

	c.JSON(200, gin.H{
		"majorExpenses": majorExpenses,
	})
}

// MajorExpensesReportDaily ..
func MajorExpensesReport(c *gin.Context) {
	ID := c.Param("id")
	var fromToDate models.DateReport
	c.ShouldBindJSON(&fromToDate)

	var exchangeReceipt []models.ExchangeReceipt
	start, end := vendors.BetweenDates(fromToDate.From, fromToDate.To)
	config.DB.Where("major_expenses_id = ?", ID).
		Where("created_at BETWEEN ? AND ?", start, end).
		Order("id desc").
		Find(&exchangeReceipt)

	c.JSON(200, gin.H{
		"exchangeReceipt": exchangeReceipt,
	})
}
