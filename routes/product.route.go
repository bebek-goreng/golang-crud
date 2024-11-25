package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, controller *controllers.ProductController) {
	productRoutes := router.Group("/api/products")
	{
		productRoutes.POST("/create", controller.CreateProduct)
		productRoutes.GET("/:id", controller.GetProductById)
		productRoutes.GET("/", controller.GetAllProducts)
		productRoutes.PUT("/update-product/:id", controller.UpdateProduct)
		productRoutes.DELETE("/delete/:id", controller.DeleteProduct)
	}
}
