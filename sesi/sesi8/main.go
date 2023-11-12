package main

import (
	"log"
	"os"
	"sesi8/database"
	"sesi8/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .enf file")
	}

	port := os.Getenv("PORT")

	database.StartDB()
	r := routers.StartApp()
	r.Run(port)
}
