// Package controllers ..
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Paginate(pageParam string, page_size string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page, _ := strconv.Atoi(pageParam)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(page_size)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// AppIndex
func AppIndex(c *gin.Context) {
	// Index Categories
	var categories []models.Categories
	config.DB.Scopes(models.WithTranslation("categories")).Find(&categories)

	// Index Select For You Items
	var selectForYouItemsFirst []models.Items
	config.DB.Preload("Categories").Scopes(models.WithTranslation("items")).Preload("StoragesItems").Order("RAND()").Limit(20).Find(&selectForYouItemsFirst)
	var selectForYouItemsSecond []models.Items
	config.DB.Preload("Categories").Preload("StoragesItems").Scopes(models.WithTranslation("items")).Order("RAND()").Limit(20).Find(&selectForYouItemsSecond)

	// Index New Items
	var newItems []models.Items
	config.DB.Preload("Categories").Preload("StoragesItems").Scopes(models.WithTranslation("items")).Order("id desc").Limit(20).Find(&newItems)

	// Index Categories With Limit Items
	var categoriesWithItems []models.CategoriesWithItems
	// Loop the categories
	for _, category := range categories {
		var categoriesItems []models.Items
		config.DB.Preload("Categories").Preload("StoragesItems").Scopes(models.WithTranslation("items")).Where("categories_id = ?", category.ID).Limit(30).Find(&categoriesItems)
		categoriesWithItems = append(categoriesWithItems, models.CategoriesWithItems{
			Category: category,
			Items:    categoriesItems,
		})
	}

	var ads []models.Ads
	config.DB.Find(&ads)

	c.JSON(200, gin.H{
		"categories":              categories,
		"selectForYouItemsFirst":  selectForYouItemsFirst,
		"selectForYouItemsSecond": selectForYouItemsSecond,
		"newItems":                newItems,
		"categoriesWithItems":     categoriesWithItems,
		"ads":                     ads,
	})
}

// IndexItemsAndSubCategoriesWithCategoriesID ..
func IndexItemsAndSubCategoriesWithCategoriesID(c *gin.Context) {
	categories_id := c.Param("id")
	var category models.Categories
	config.DB.Scopes(models.WithTranslation("categories")).Where("id = ?", categories_id).First(&category)

	var subCategories []models.SubCategories
	config.DB.Scopes(models.WithTranslation("subCategories")).Where("categories_id = ?", categories_id).Find(&subCategories)

	var selectForYouItems []models.Items
	config.DB.Preload("Categories").Scopes(models.WithTranslation("items"), models.WithStorageCounts()).Where("categories_id = ?", categories_id).Order("RAND()").Limit(15).Find(&selectForYouItems)

	var items []models.Items
	err := config.DB.Preload("Categories").Preload("StoragesItems").Where("categories_id = ?", categories_id).Scopes(Paginate(c.Param("page"), "14"), models.WithTranslation("items")).Find(&items).Error
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"category":          category,
		"selectForYouItems": selectForYouItems,
		"items":             items,
		"subCategories":     subCategories,
	})
}

// CategoriesItemsLoadMore ..
func CategoriesItemsLoadMore(c *gin.Context) {
	categoryID := c.Param("id")
	page := c.Param("page")

	var items []models.Items
	err := config.DB.Preload("Categories").Preload("StoragesItems").Where("categories_id = ?", categoryID).Scopes(Paginate(page, "14"), models.WithTranslation("items")).Find(&items).Error
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}

// ShowItem ..
func AppShowItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Items
	config.DB.Preload("Categories").Preload("StoragesItems").Scopes(models.WithTranslation("items")).Where("id = ?", id).First(&item)

	var relatedItems []models.Items
	config.DB.Preload("Categories").Preload("StoragesItems").Scopes(models.WithTranslation("items")).Order("RAND()").Limit(14).Find(&relatedItems)

	c.JSON(200, gin.H{
		"item":         item,
		"relatedItems": relatedItems,
	})
}

// AppIndexItemsWithSubCategories ..
func AppIndexItemsWithSubCategories(c *gin.Context) {
	subCategoryID := c.Param("id")
	var subCategory models.SubCategories
	config.DB.Scopes(models.WithTranslation("subCategories")).Where("id = ?", subCategoryID).Find(&subCategory)

	var items []models.Items
	err := config.DB.Preload("Categories").Preload("StoragesItems").Scopes(models.WithTranslation("items")).Where("sub_categories_id = ?", subCategoryID).Scopes(Paginate(c.Param("page"), "14")).Find(&items).Error
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"items":       items,
		"subCategory": subCategory,
	})
}

// UserTest ..
type UserTest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// TestExternalCall ..
func TestExternalCall(c *gin.Context) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		fmt.Println("error call:" + err.Error())
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data UserTest
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	fmt.Println(data)

	c.JSON(200, gin.H{
		"user": data,
	})
}

// IndexShowAllItem ..
func IndexShowAllItem(c *gin.Context) {
	indexType := c.Param("type")

	var items []models.Items
	if indexType == "selectForYou" {
		config.DB.Order("RAND()").
			Preload("Categories").
			Preload("StoragesItems").
			Scopes(Paginate(c.Param("page"), "14"), models.WithTranslation("items")).
			Find(&items)
	}

	if indexType == "newItems" {
		config.DB.Order("id desc").
			Preload("Categories").
			Preload("StoragesItems").
			Scopes(Paginate(c.Param("page"), "14"), models.WithTranslation("items")).
			Find(&items)
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}

// TestItemsType

// TestItemsWithStorageCount ..
func TestItemsWithStorageCount(c *gin.Context) {
	var items []models.Items
	config.DB.Preload("Categories").Preload("StoragesItems").Scopes(models.WithTranslation("items")).Order("RAND()").Limit(15).Find(&items)

	// err := config.DB.Scopes(models.WithStorageCounts()).Find(&items).Error
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"error": err,
	// 	})
	// 	return
	// }

	c.JSON(200, gin.H{
		"items": items,
	})

}
