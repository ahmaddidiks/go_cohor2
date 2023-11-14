package controllers

import (
	"net/http"
	"sesi8/database"
	"sesi8/helpers"
	"sesi8/models"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Book := models.Book{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Book)
	} else {
		ctx.ShouldBind(&Book)
	}

	Book.UserID = userID
	Book.UUID = uuid.New().String()

	err := db.Debug().Create(&Book).Error
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

func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Book := models.Book{}

	bookUUID := ctx.Param("bookUUID")
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Book)
	} else {
		ctx.ShouldBind(&Book)
	}

	// Retrieve existing book from the database
	var getBook models.Book
	if err := db.Model(&getBook).Where("uuid = ?", bookUUID).First(&getBook).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Update the Book struct with retrieved data
	Book.ID = uint(getBook.ID)
	Book.UserID = userID

	// Update the book record in the database
	updateData := models.Book{
		Title:  Book.Title,
		Author: Book.Author,
		Stock:  Book.Stock,
	}

	if err := db.Model(&Book).Where("uuid = ?", bookUUID).Updates(updateData).Error; err != nil {
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

func GetBooks(ctx *gin.Context) {
	db := database.GetDB()

	results := []models.Book{}

	err := db.Debug().Preload("User").Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}
