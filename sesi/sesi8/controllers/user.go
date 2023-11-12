package controllers

import (
	"net/http"
	"sesi8/database"
	"sesi8/helpers"
	"sesi8/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	appJSON = "application/json"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	contextType := helpers.GetContentType(ctx)
	user := models.User{}

	if contextType == appJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	// generate UUID
	newUUID := uuid.New()
	user.UUID = newUUID.String()

	err := db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	user := models.User{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	password := user.Pasword

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid emali",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Pasword), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)
	
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
