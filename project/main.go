package main

import (
	"log"
	"os"
	"rest-api/database"
	"rest-api/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("WEB_SERVER_PORT")

	database.StartDB()
	routers.StartServer().Run(PORT)
}