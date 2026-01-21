package main

import (
	"log"
	"os"
	"task-manager/config"
	"task-manager/routes"
	"task-manager/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load .env locally if not in Docker

	// 1. Connect to DB
	config.ConnectDB()

	// 2. Start Background Worker
	services.StartWorker()

	// 3. Init Router
	r := gin.Default()
	routes.SetupRoutes(r)

	// 4. Run Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(r.Run(":" + port))
}
