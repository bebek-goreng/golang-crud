package main

import (
	"go-rest-api/controllers"
	"go-rest-api/models"
	"go-rest-api/routes"
	"go-rest-api/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set in .env file")
	}

	db = models.DB

	r := gin.Default()

	productService := &services.ProductService{DB: db}
	productController := &controllers.ProductController{Service: productService}

	routes.SetupProductRoutes(r, productController)

	r.Run(":" + port)
}
