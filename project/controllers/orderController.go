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
		msg := "Failed to connect database"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// create order
	order := models.Order{
		Customername: newOrder.CustomerName,
		OrderAt:      newOrder.OrderAt,
	}
	err := db.Create(&order).Error
	if err != nil {
		msg := "Failed to create the order"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// get order_id from last input data
	err = db.Last(&order, "customername = ?", order.Customername).Error
	if err != nil {
		msg := "Failed to create the order"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
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
		msg := "Failed to create the item order"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newOrder,
		"message": "Order successfully created.",
	})
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Order
	var items []models.Item
	var results []models.ItemRecv

	db := database.GetDB()
	if db == nil {
		msg := "Failed to to connect database"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// get order
	err := db.Find(&orders).Error
	if err != nil {
		msg := "No order data in database"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}
	// get item
	err = db.Find(&items).Error
	if err != nil {
		msg := "No items data in database"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// copy the data
	for i := 0; i < len(orders); i++ {
		itemResult := models.ItemOrderRecv{
			Name:        items[i].Name,
			Description: items[i].Description,
			Quantity:    items[i].Quantity,
		}

		result := models.ItemRecv{
			CustomerName: orders[i].Customername,
			OrderAt:      orders[i].OrderAt,
			Items:        []models.ItemOrderRecv{itemResult},
		}

		results = append(results, result)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    results,
		"message": "All order data",
	})
}

func GetOrderbyID(ctx *gin.Context) {
	var order models.Order
	var item models.Item

	orderID := ctx.Param("id")
	convertedOrderID, _ := strconv.Atoi(orderID)

	db := database.GetDB()
	if db == nil {
		msg := "Failed to connect database"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// get order and item
	err := db.First(&order, "id = ?", convertedOrderID).Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v not found", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	err = db.First(&item, "order_id = ?", convertedOrderID).Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v has no Item", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	itemResult := models.ItemOrderRecv{
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
	}
	result := models.ItemRecv{
		OrderAt:      order.OrderAt,
		CustomerName: order.Customername,
		Items:        []models.ItemOrderRecv{itemResult},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": fmt.Sprintf("Order ID %v data", orderID),
	})
}

func UpdateOrder(ctx *gin.Context) {
	var newOrder models.ItemRecv
	var order models.Order
	var item models.Item

	orderID := ctx.Param("id")
	convertedOrderID, _ := strconv.Atoi(orderID)

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	if db == nil {
		msg := "Failed to to connect database"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// check order is exist in db
	err := db.First(&order, "id = ?", convertedOrderID).Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v not found", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}
	// update order
	updateStatement := map[string]interface{}{
		"customername": newOrder.CustomerName,
		"order_at":     newOrder.OrderAt}
	err = db.Model(&order).Where("id = ?", convertedOrderID).Updates(updateStatement).Error
	if err != nil {
		msg := fmt.Sprintf("Update Order ID %v Failed", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	// get item_id from db
	err = db.First(&item, "order_id = ?", convertedOrderID).Error
	if err != nil {
		msg := fmt.Sprintf("Order ID %v Has no Item", orderID)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
	}

	updateStatement = map[string]interface{}{
		"name":        newOrder.Items[0].Name,
		"description": newOrder.Items[0].Description,
		"quantity":    newOrder.Items[0].Quantity,
	}
	err = db.Model(&item).Where("order_id = ?", convertedOrderID).Updates(updateStatement).Error
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
		msg := "Failed to connect database"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": msg,
		})
		return
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
