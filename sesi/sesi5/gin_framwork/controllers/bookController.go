package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"id"`
	Title  string `json:"title"`
	Stock  string `json:"stock"`
	Author string `json:"author"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	// var newBook Book

	// if err := ctx.ShouldBindJSON(&newBook); err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// }

}

func GetBook(ctx *gin.Context) {
	// bookID := ctx.Params("BookID")
	// condition := false
	// var bookData
}
