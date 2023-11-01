package controllers

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var newOrder models.ItemRecv

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newOrder,
		"message": "succeed order",
	})
}
