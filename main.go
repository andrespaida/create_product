package main

import (
	"create_product/handlers"
	"create_product/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️  Error loading .env")
	}

	if err := models.ConnectDB(); err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	r := gin.Default()
	r.POST("/products", handlers.CreateProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}