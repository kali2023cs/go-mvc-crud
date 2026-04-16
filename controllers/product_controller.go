package controllers

import (
	"net/http"
	"go-mvc-crud/config"
	"go-mvc-crud/models"
	"go-mvc-crud/utils"

	"github.com/gin-gonic/gin"
)

// Get all products (with optional filters)
func GetProducts(c *gin.Context) {
	var products []models.Product
	query := config.DB

	// Filtering by name (partial match)
	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// Filtering by min price
	if minPrice := c.Query("min_price"); minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}

	// Filtering by max price
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	query.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// Create a product
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": utils.FormatError(err)})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
	config.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// Get a single product
func GetProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// Update a product
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": utils.FormatError(err)})
		return
	}

	config.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
