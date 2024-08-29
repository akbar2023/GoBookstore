package main

import (
	"example/web-service-gin/database"
	"example/web-service-gin/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	router := gin.Default()
	routes.RegisterBookRoutes(router)
	router.Run("localhost:8080")
}
