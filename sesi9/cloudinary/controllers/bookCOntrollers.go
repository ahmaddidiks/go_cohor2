package controllers

import (
	"sesi_9/cloudinary/request"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	var bookReq request.BookRequest
	if rtt := ctx.ShouldBind()

}
