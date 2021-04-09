// Package controllers ..
package controllers

import (
	"image"
	"net/http"
	"os"
	"path/filepath"
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SaveTranslations ..
func SaveTranslations(tableName string, columnName string, foreignKey uint, value string) {
	config.DB.Create(&models.Translations{
		TableName:  tableName,
		ColumnName: columnName,
		Locale:     "en",
		ForeignKey: foreignKey,
		Value:      value,
	})
}

// UpdateTranslations ..
func UpdateTranslations(tableName string, columnName string, foreignKey uint, value string) {
	var translations models.Translations
	if err := config.DB.Where("table_name = ?", tableName).
		Where("column_name = ?", columnName).
		Where("foreign_key = ?", foreignKey).
		First(&translations).Error; err != nil {
		SaveTranslations(tableName, columnName, foreignKey, value)
		return
	}
	translations.Value = value

	config.DB.Save(&translations)
}

// StoreCategories ..
func StoreCategories(c *gin.Context) {
	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Translations
	SaveTranslations("categories", "title", category.ID, category.TranslationsLocale.Title)
	SaveTranslations("categories", "description", category.ID, category.TranslationsLocale.Description)

	var categoryLocale models.Categories
	config.DB.Where("id = ?", category.ID).Scopes(models.WithTranslation("categories")).First(&categoryLocale)
	c.JSON(200, gin.H{
		"category": categoryLocale,
	})
}

// IndexCategories ...
func IndexCategories(c *gin.Context) {
	var categories []models.Categories

	config.DB.Scopes(models.WithTranslation("categories")).Find(&categories)

	c.JSON(200, gin.H{
		"categories": categories,
	})
}

// UpdateCategory ..
func UpdateCategory(c *gin.Context) {
	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&category).Update(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateTranslations("categories", "title", category.ID, category.TranslationsLocale.Title)
	UpdateTranslations("categories", "description", category.ID, category.TranslationsLocale.Description)

	var categories []models.Categories

	config.DB.Scopes(models.WithTranslation("categories")).Find(&categories)

	c.JSON(200, gin.H{
		"categories": categories,
		"category":   category,
	})
}

// DestroyCategory ..
func DestroyCategory(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Categories{}, ID)

	var categories []models.Categories
	config.DB.Scopes(models.WithTranslation("categories")).Find(&categories)
	c.JSON(200, gin.H{
		"categories": categories,
	})
}

// StoreSubCategories ..
func StoreSubCategories(c *gin.Context) {
	var subCategory models.SubCategories
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&subCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Translations
	SaveTranslations("subCategories", "title", subCategory.ID, subCategory.TranslationsLocale.Title)
	SaveTranslations("subCategories", "description", subCategory.ID, subCategory.TranslationsLocale.Description)

	config.DB.Scopes(models.WithTranslation("subCategories")).Preload("Categories").Where("id = ?", subCategory.ID).First(&subCategory)

	c.JSON(200, gin.H{
		"subCategories": subCategory,
	})
}

// UpdateSubCategory ..
func UpdateSubCategory(c *gin.Context) {
	var subCategory models.SubCategories
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&subCategory).Update(&subCategory)

	UpdateTranslations("subCategories", "title", subCategory.ID, subCategory.TranslationsLocale.Title)
	UpdateTranslations("subCategories", "description", subCategory.ID, subCategory.TranslationsLocale.Description)

	var subCategories []models.SubCategories
	config.DB.Scopes(models.WithTranslation("subCategories")).Preload("Categories").Find(&subCategories)

	c.JSON(200, gin.H{
		"subCategories": subCategories,
	})
}

// DestroySubCategory ..
func DestroySubCategory(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.SubCategories{}, ID)

	var subCategories []models.SubCategories
	config.DB.Scopes(models.WithTranslation("subCategories")).Preload("Categories").Find(&subCategories)

	c.JSON(200, gin.H{
		"subCategories": subCategories,
	})
}

// IndexSubCategories ...
func IndexSubCategories(c *gin.Context) {
	var subCategories []models.SubCategories
	config.DB.Scopes(models.WithTranslation("subCategories")).Preload("Categories").Find(&subCategories)

	c.JSON(200, gin.H{
		"subCategories": subCategories,
	})
}

// IndexSubCategoriesWithID ...
func IndexSubCategoriesWithID(c *gin.Context) {
	ID := c.Param("categories_id")

	var subCategories []models.SubCategories
	config.DB.Scopes(models.WithTranslation("subCategories")).Preload("Categories").Where("categories_id = ?", ID).Find(&subCategories)

	c.JSON(200, gin.H{
		"subCategories": subCategories,
	})

}

// StoreColors ..
func StoreColors(c *gin.Context) {
	var color models.Colors
	if err := c.ShouldBindJSON(&color); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&color).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"color": color,
	})
}

// UpdateColor ..
func UpdateColor(c *gin.Context) {
	var color models.Colors
	if err := c.ShouldBindJSON(&color); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&color).Update(&color)

	var colors []models.Colors
	config.DB.Find(&colors)
	c.JSON(200, gin.H{
		"colors": colors,
		"color":  color,
	})
}

// DestroyColor ..
func DestroyColor(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Colors{}, id)

	var colors []models.Colors

	config.DB.Find(&colors)

	c.JSON(200, gin.H{
		"colors": colors,
	})
}

// IndexColors ..
func IndexColors(c *gin.Context) {
	var colors []models.Colors
	config.DB.Find(&colors)

	c.JSON(200, gin.H{
		"colors": colors,
	})
}

// StoreItems ..
func StoreItems(c *gin.Context) {
	var item models.Items
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Translations
	SaveTranslations("items", "title", item.ID, item.TranslationsLocale.Title)
	SaveTranslations("items", "description", item.ID, item.TranslationsLocale.Description)

	config.DB.Scopes(models.WithTranslation("items")).Preload("Categories").Preload("SubCategories").Where("id = ?", item.ID).First(&item)

	c.JSON(200, gin.H{
		"item": item,
	})
}

// IndexItems ..
func IndexItems(c *gin.Context) {
	var items []models.Items
	config.DB.Scopes(models.WithTranslation("items")).Preload("Categories").Preload("SubCategories").Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

// UpdateItem ..
func UpdateItem(c *gin.Context) {
	var item models.Items
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Update(&item)

	UpdateTranslations("items", "title", item.ID, item.TranslationsLocale.Title)
	UpdateTranslations("items", "description", item.ID, item.TranslationsLocale.Description)

	var items []models.Items
	config.DB.Scopes(models.WithTranslation("items")).Preload("Categories").Preload("SubCategories").Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

// SearchItemsModel ..
type SearchItemsModel struct {
	ItemTitle   string `json:"item_title"`
	ItemID      uint   `json:"item_id"`
	ItemBarcode string `json:"item_barcode"`
	Type        string `json:"type"`
}

// SearchItems ..
func SearchItems(c *gin.Context) {
	var searchInput SearchItemsModel
	if err := c.ShouldBindJSON(&searchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var items []models.Items
	if searchInput.Type == "title" {
		config.DB.Scopes(models.WithTranslation("items")).Where("title Like ?", "%"+searchInput.ItemTitle+"%").
			Limit(15).
			Find(&items)
	} else if searchInput.Type == "barcode" {
		config.DB.Scopes(models.WithTranslation("items")).Where("barcode Like ?", "%"+searchInput.ItemBarcode+"%").
			Limit(15).
			Find(&items)
	} else if searchInput.Type == "ID" {
		config.DB.Scopes(models.WithTranslation("items")).Where("id = ?", searchInput.ItemID).
			Limit(15).
			Find(&items)
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}

// ShowItem ..
func ShowItem(c *gin.Context) {
	id := c.Param("id")

	var item models.Items
	config.DB.Scopes(models.WithTranslation("items")).Preload("Categories").Preload("SubCategories").Where("id = ?", id).First(&item)

	var itemStorages []models.StoragesItems
	config.DB.Where("item_id = ?", id).Find(&itemStorages)

	var storageItemsLogs []models.StorageItemsLogs
	config.DB.Where("item_id = ?", id).Find(&storageItemsLogs)

	c.JSON(200, gin.H{
		"item":             item,
		"itemStorages":     itemStorages,
		"storageItemsLogs": storageItemsLogs,
	})
}

// DestroyItem ..
func DestroyItem(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Items{}, id)

	var items []models.Items
	config.DB.Scopes(models.WithTranslation("items")).Preload("Categories").Preload("SubCategories").Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

// StoreCompany ..
func StoreCompany(c *gin.Context) {
	var company models.Companies
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"company": company,
	})
}

// IndexCompanies ..
func IndexCompanies(c *gin.Context) {
	var companies []models.Companies

	config.DB.Find(&companies)

	c.JSON(200, gin.H{
		"companies": companies,
	})
}

// UpdateCompany ..
func UpdateCompany(c *gin.Context) {
	var company models.Companies
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&company).Update(&company)

	var companies []models.Companies

	config.DB.Find(&companies)

	c.JSON(200, gin.H{
		"companies": companies,
	})
}

// DestroyCompany ..
func DestroyCompany(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Companies{}, id)
	var companies []models.Companies

	config.DB.Find(&companies)

	c.JSON(200, gin.H{
		"companies": companies,
	})
}

// UpdateImage ..
func UpdateImage(c *gin.Context) {
	file, _ := c.FormFile("image")
	imageType := c.Param("imageType")

	filename := filepath.Base(file.Filename)
	id := uuid.NewString()

	if err := c.SaveUploadedFile(file, config.ServerInfo.PublicPath+"public/"+imageType+"/"+filename); err != nil {
		c.JSON(500, gin.H{
			"error":   err.Error(),
			"message": "error",
		})
		return
	}

	afterUploadImage, _ := os.Open(config.ServerInfo.PublicPath + "public/" + imageType + "/" + filename)
	_, format, _ := image.DecodeConfig(afterUploadImage)
	afterUploadImage.Close()

	id = id + "." + format
	vendors.ResizeImage(filename, id, config.ServerInfo.PublicPath+"public/"+imageType+"/", format)
	c.JSON(200, gin.H{
		"image":   id,
		"message": "success",
	})
}

// StoreStorages ..
func StoreStorages() {
	var StorageCounts int64
	config.DB.Model(&models.Storages{}).Count(&StorageCounts)
	if StorageCounts == 0 {
		config.DB.Create(&models.Storages{
			Title: "الكل",
			Scope: "All",
		})
		config.DB.Create(&models.Storages{
			Title: "جملة الجملة",
			Scope: "GGlobal",
		})
		config.DB.Create(&models.Storages{
			Title: "الجملة",
			Scope: "Global",
		})
		config.DB.Create(&models.Storages{
			Title: "المفرق",
			Scope: "Parts",
		})
		config.DB.Create(&models.Storages{
			Title: "متجر الكتروني",
			Scope: "Ecommarce",
		})
	}

}

// IndexStorages ..
func IndexStorages(c *gin.Context) {
	var storages []models.Storages

	config.DB.Order("id asc").Find(&storages)

	c.JSON(200, gin.H{
		"storages": storages,
	})
}
