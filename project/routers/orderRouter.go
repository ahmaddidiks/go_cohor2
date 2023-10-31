package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", test)
	router.GET("/orders", test)

	return router
}

func test(ctx *gin.Context) {

}
