package controllers

import (
	"net/http"
	"strconv"

	"go-mvc-crud/dto"
	"go-mvc-crud/services"
	"go-mvc-crud/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// Get all products (with optional filters)
func (h *ProductHandler) GetProducts(c *gin.Context) {
	name := c.Query("name")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")

	products, err := h.service.GetAllProducts(name, minPrice, maxPrice)
	if err != nil {
		utils.HandleError(c, utils.NewInternalServerError("Failed to fetch products"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// Create a product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input dto.CreateProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleError(c, utils.NewBadRequestError("Validation failed", utils.FormatError(err)))
		return
	}

	product, err := h.service.CreateProduct(input)
	if err != nil {
		utils.HandleError(c, utils.NewInternalServerError("Failed to create product"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

// Get a single product
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		utils.HandleError(c, utils.NewNotFoundError("Product not found"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// Update a product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleError(c, utils.NewBadRequestError("Validation failed", utils.FormatError(err)))
		return
	}

	product, err := h.service.UpdateProduct(uint(id), input)
	if err != nil {
		utils.HandleError(c, utils.NewNotFoundError("Product not found or update failed"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// Delete a product
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteProduct(uint(id)); err != nil {
		utils.HandleError(c, utils.NewNotFoundError("Product not found"))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
