// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreUserRoles ..
func StoreUserRoles(c *gin.Context) {
	var role models.Roles
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"role": role,
	})
}

// IndexUserRoles ...
func IndexUserRoles(c *gin.Context) {
	var roles []models.Roles
	config.DB.Find(&roles)

	c.JSON(200, gin.H{
		"roles": roles,
	})
}

// UpdateUserRole ...
func UpdateUserRole(c *gin.Context) {
	var role models.Roles
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&role).Update(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var roles []models.Roles
	config.DB.Find(&roles)

	c.JSON(200, gin.H{
		"role":  role,
		"roles": roles,
	})
}

// DeleteUserRole ...
func DeleteUserRole(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Roles{}, ID)
	var roles []models.Roles
	config.DB.Find(&roles)
	c.JSON(200, gin.H{
		"roles": roles,
	})
}
