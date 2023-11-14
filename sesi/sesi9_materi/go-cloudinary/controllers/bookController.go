package controllers

import (
	"github.com/gin-gonic/gin"
	"go-cloudinary/database"
	"go-cloudinary/helpers"
	models "go-cloudinary/models/entity"
	requests "go-cloudinary/models/request"

	"net/http"
)

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	var bookReq requests.BookRequest
	if err := ctx.ShouldBind(&bookReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the filename without extension
	fileName := helpers.RemoveExtension(bookReq.Image.Filename)

	uploadResult, err := helpers.UploadFile(bookReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Book := models.Book{
		Title:    bookReq.Title,
		Author:   bookReq.Author,
		Stock:    bookReq.Stock,
		ImageURL: uploadResult,
	}

	err = db.Debug().Create(&Book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Book,
	})
}
