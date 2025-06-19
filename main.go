package main

import (
	"create_product/handlers"
	"create_product/models"
	"github.com/gin-contrib/cors"
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

	// ✅ Habilitar CORS globalmente
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // o especifica: []string{"http://98.85.86.231"}
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/products", handlers.CreateProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4002" // cambiamos a 4002 como definiste
	}

	log.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}