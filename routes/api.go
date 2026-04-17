package routes

import (
	"go-mvc-crud/config"
	"go-mvc-crud/controllers"
	"go-mvc-crud/repositories"
	"go-mvc-crud/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize Dependencies
	productRepo := repositories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepo)
	productHandler := controllers.NewProductHandler(productService)

	api := r.Group("/api")
	{
		api.GET("/products", productHandler.GetProducts)
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products/:id", productHandler.GetProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	return r
}
