package routers

import (
	"sesi8/controllers"
	"sesi8/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	book := router.Group("/books")
	{
		book.Use(middleware.Authentication())
		book.GET("/", controllers.GetBooks)
		book.POST("/", controllers.CreateBook)
		book.PUT("/:bookUUID", middleware.BookAuthorization(), controllers.UpdateBook)
	}

	return router
}
