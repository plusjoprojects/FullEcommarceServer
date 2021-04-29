// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"
	"strconv"

	"github.com/gin-gonic/gin"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// IndexNewOrders ..
func IndexNewOrders(c *gin.Context) {
	status := c.Param("status")
	var orders []models.Orders
	config.DB.Where("type = ?", "Application").Where("status = ?", status).Find(&orders)

	c.JSON(200, gin.H{
		"orders": orders,
	})
}

// ShowOrder ..
func ShowOrder(c *gin.Context) {
	orderID := c.Param("id")

	var order models.Orders
	config.DB.Where("id = ?", orderID).First(&order)

	var orderItems []models.OrderItems
	config.DB.Preload("Item").Where("order_id = ?", orderID).Find(&orderItems)

	var user models.User
	config.DB.Where("id = ?", order.UserID).First(&user)

	var address models.Addresses
	config.DB.Where("id = ?", order.AddressID).First(&address)

	c.JSON(200, gin.H{
		"order":      order,
		"orderItems": orderItems,
		"user":       user,
		"address":    address,
	})
}

func SendNotificationWithUserID(userID uint, orderID uint, title string, body string) {
	var notificationToken models.NotificationToken
	config.DB.Where("user_id = ?", userID).First(&notificationToken)
	var orderIDString string = strconv.FormatUint(uint64(orderID), 10)
	var tokens []expo.ExponentPushToken
	token, _ := expo.NewExponentPushToken(notificationToken.Token)
	tokens = append(tokens, token)

	data := vendors.NotificationData{
		Type: "Order",
		ID:   orderIDString,
	}

	message := vendors.NotificationMessage{
		Title: title,
		Body:  body,
		Data:  data,
	}

	vendors.SendNotification(tokens, message, data)
}

// OrderActionsType ..
type OrderActionsType struct {
	OrderID    uint   `json:"order_id"`
	ActionType string `json:"action_Type"`
	EmployeeID uint   `json:"employee_id"`
}

// OrderActions ..
func OrderActions(c *gin.Context) {

	var data OrderActionsType
	c.ShouldBindJSON(&data)

	orderID := data.OrderID
	actionType := data.ActionType
	employeeID := data.EmployeeID

	var order models.Orders
	config.DB.Where("id = ?", orderID).First(&order)

	if actionType == "approve" {
		// Here Function For approve
		order.Status = 1
		order.EmployeeID = employeeID
		config.DB.Save(&order)
		var title string = "تم تجهيز الطلبية"
		var body string = "تم تجهيز طلبيتك وسوف يتم ارسالها"
		SendNotificationWithUserID(order.UserID, order.ID, title, body)
	}

	if actionType == "sendToDelivery" {
		// Here Function For sendToDelivery
		order.Status = 2
		config.DB.Save(&order)
		var title string = "تم ارسال طلبك"
		var body string = "تم ارسال طلبك وطلبك حاليا قيد التوصيل"
		SendNotificationWithUserID(order.UserID, order.ID, title, body)
	}

	if actionType == "endOrder" {
		// Here Function For endOrder
		order.Status = 3
		config.DB.Save(&order)
		var title string = "تم انهاء الطلب"
		var body string = "تم انهاء طلبك بنجاح, شكرا لتعاملك معنا"
		SendNotificationWithUserID(order.UserID, order.ID, title, body)
	}

	if actionType == "cancelOrder" {
		// Here Function For CancelOrder
		order.Status = 4
		config.DB.Save(&order)
		var title string = "تم الغاء الطلب"
		var body string = "تم الغاء الطلب من قبل الادارة"
		SendNotificationWithUserID(order.UserID, order.ID, title, body)
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
