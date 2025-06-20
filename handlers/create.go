package handlers

import (
	"create_product/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
	"os"
)

func CreateProduct(c *gin.Context) {
	// Parsear campos del formulario
	name := c.PostForm("name")
	description := c.PostForm("description")
	price := c.PostForm("price")
	stock := c.PostForm("stock")

	// Manejo del archivo de imagen
	file, err := c.FormFile("image")
	var imageURL string

	if err == nil {
		// Crear carpeta si no existe
		os.MkdirAll("uploads", os.ModePerm)

		// Nombre único basado en timestamp
		filename := time.Now().Format("20060102150405") + "_" + filepath.Base(file.Filename)
		filepath := "uploads/" + filename

		// Guardar la imagen localmente
		if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		imageURL = "/uploads/" + filename
	}

	// Convertir price y stock
	parsedPrice, parseErr1 := models.ParseFloat(price)
	parsedStock, parseErr2 := models.ParseInt(stock)
	if parseErr1 != nil || parseErr2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price or stock"})
		return
	}

	product := models.Product{
		Name:        name,
		Description: description,
		Price:       parsedPrice,
		Stock:       parsedStock,
		ImageURL:    imageURL,
	}

	collection := models.ProductCollection()
	_, err = collection.InsertOne(c, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}