package main

import (
	"go-cloudinary/database"
	"go-cloudinary/router"
)

var (
	PORT = ":5050"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
