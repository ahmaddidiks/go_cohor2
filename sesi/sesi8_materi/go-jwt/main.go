package main

import (
	"go-jwt/database"
	"go-jwt/router"
)

var (
	PORT = ":7070"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
