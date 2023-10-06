package routers

import (
	"go_cohort_2/sesi5/gin_framwork/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)

	return router
}
