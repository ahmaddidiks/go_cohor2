package routers

import (
	"log"
	"os"
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	trustedIP := os.Getenv("TRUSTED_IP")

	var trustedIPs []string
	trustedIPs = append(trustedIPs, trustedIP)

	router := gin.Default()
	// fix trusted all proxies this is not safe
	router.ForwardedByClientIP = true
	router.SetTrustedProxies(trustedIPs)

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:id", controllers.GetOrderbyID) // ini tambahan saja
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)

	return router
}
