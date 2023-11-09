package router

import (
	"go-jwt/controllers"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	bookRouter := router.Group("/books")
	{
		bookRouter.GET("/", controllers.GetBooks)

		bookRouter.Use(middlewares.Authentication())
		bookRouter.POST("/", controllers.CreateBook)
		bookRouter.PUT("/:bookUUID", middlewares.BookAuthorization(), controllers.UpdateBook)
	}

	return router
}
