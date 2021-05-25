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

// StoreDelegate ..
func StoreDelegate(c *gin.Context) {
	var delegate models.Delegates
	c.ShouldBindJSON(&delegate)

	config.DB.Create(&delegate)

	c.JSON(200, gin.H{
		"delegate": delegate,
	})
}

// IndexDelegates ..
func IndexDelegates(c *gin.Context) {
	var delegates []models.Delegates

	config.DB.Find(&delegates)

	c.JSON(200, gin.H{
		"delegates": delegates,
	})
}

// DeleteDelegate ..
func DeleteDelegate(c *gin.Context) {
	ID := c.Param("id")

	config.DB.Delete(&models.Delegates{}, ID)

	var delegates []models.Delegates
	config.DB.Find(&delegates)

	c.JSON(200, gin.H{
		"delegates": delegates,
	})
}

// UpdateDelegate ..
func UpdateDelegate(c *gin.Context) {
	var delegate models.Delegates

	c.ShouldBindJSON(&delegate)

	config.DB.Model(&delegate).Update(&delegate)

	var delegates []models.Delegates

	config.DB.Find(&delegates)

	c.JSON(200, gin.H{
		"delegates": delegates,
		"delegate":  delegate,
	})

}

// GetDriversList ..
func GetDriversList(c *gin.Context) {
	var driverRole models.Roles

	config.DB.Where("scopes LIKE ?", "%driver%").Find(&driverRole)

	roleID := driverRole.ID

	var drivers []models.User

	config.DB.Preload("Roles").Preload("Employee").Where("roles_id = ?", roleID).Find(&drivers)

	c.JSON(200, gin.H{
		"drivers": drivers,
	})
}
