package router

import (
	"github.com/gin-gonic/gin"
	"go-cloudinary/controllers"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	bookRouter := router.Group("/books")
	{
		bookRouter.POST("/", controllers.CreateBook)
	}

	return router
}
