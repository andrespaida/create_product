package models

import (
	"context"
	"os"
	"time"
	"strconv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
	ImageURL    string  `json:"image_url" bson:"image_url"`
}

var client *mongo.Client

func ProductCollection() *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	return client.Database(dbName).Collection("products")
}

func ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	return err
}

// ✅ Funciones auxiliares para convertir datos de string a float64 e int
func ParseFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func ParseInt(value string) (int, error) {
	return strconv.Atoi(value)
}