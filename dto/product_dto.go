package dto

import "time"

// CreateProductRequest defines the structure for product creation
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,not-reserved"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,numeric,gt=0"`
}

// UpdateProductRequest defines the structure for product update
type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"omitempty,not-reserved"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"omitempty,numeric,gt=0"`
}

// ProductResponse defines the structure for returning product data
type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
