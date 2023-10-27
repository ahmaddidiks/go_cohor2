package router

import "github.com/gin-gonic/gin"

func StartApp() *gin.Engine {
	router := gin.Default()

	bookRouter := router.Group("/books")
	{

	}
}
