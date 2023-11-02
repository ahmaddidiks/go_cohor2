package controllers

import (
	"fmt"
	"net/http"
	"rest-api/database"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var newOrder models.ItemRecv

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	if db == nil {
		panic("Error: Database connection is nil")
	}

	// create order
	order := models.Order{
		Customername: newOrder.CustomerName,
		OrderAt:      newOrder.OrderAt,
	}
	err := db.Create(&order).Error
	if err != nil {
		panic(err)
	}

	// get order_id from last input data
	err = db.Last(&order, "customername = ?", order.Customername).Error
	if err != nil {
		panic(err)
	}

	// create order
	item := models.Item{
		Name:        newOrder.Items[0].Name,
		Description: newOrder.Items[0].Description,
		Quantity:    newOrder.Items[0].Quantity,
		OrderID:     order.ID,
	}
	err = db.Create(&item).Error
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newOrder,
		"message": "Order successfully created.",
	})
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Item

	ctx.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "All order data",
	})
}

func GetOrderbyID(ctx *gin.Context) {
	var order models.Item
	orderID := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": fmt.Sprintf("Order ID %v data", orderID),
	})
}

func UpdateOrder(ctx *gin.Context) {
	var newOrder models.ItemRecv
	var order models.Order

	orderID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	if db == nil {
		panic("Error: Database connection is nil")
	}

	// check if orderID exist
	row := db.Where("id = ?", orderID).Limit(1).Find(&order)
	err := row.Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v not found", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}
	exists := row.RowsAffected > 0
	if !exists {
		msg := fmt.Sprintf("Order ID %v not found", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	convertedOrderID, _ := strconv.Atoi(orderID)
	// update order , save order with desire ID it will update the desire ID
	err = db.Save(&models.Order{
		ID:           uint(convertedOrderID),
		Customername: newOrder.CustomerName,
		OrderAt:      newOrder.OrderAt,
	}).Error

	if err != nil {
		msg := fmt.Sprintf("Update Order ID %v Failed", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}
	// update item
	err = db.Save(&models.Item{
		ID:          uint(convertedOrderID),
		Name:        newOrder.Items[0].Name,
		Description: newOrder.Items[0].Description,
		Quantity:    newOrder.Items[0].Quantity,
	}).Error

	if err != nil {
		msg := fmt.Sprintf("Update Item of Order ID %v Failed", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": fmt.Sprintf("Order ID %v successfully updated.", orderID),
	})
}

func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")

	db := database.GetDB()
	if db == nil {
		panic("Error: Database connection is nil")
	}

	order := models.Order{}
	item := models.Item{}

	// check if order exist
	row := db.Where("id = ?", orderID).Limit(1).Find(&order)
	err := row.Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v not found", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}
	exists := row.RowsAffected > 0
	if !exists {
		msg := fmt.Sprintf("Order ID %v not found", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// delete items
	err = db.Where("order_id = ?", orderID).Delete(&item).Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v no item", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		// return
	}
	// delete order
	err = db.Where("id = ?", orderID).Delete(&order).Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v not found", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	msg := fmt.Sprintf("Order ID %v has been deleted.", orderID)
	ctx.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": msg,
	})
}
