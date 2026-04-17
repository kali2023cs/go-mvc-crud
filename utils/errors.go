package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppError represents a standardized error structure
type AppError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

// NewNotFoundError returns a 404 error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

// NewBadRequestError returns a 400 error
func NewBadRequestError(message string, errors interface{}) *AppError {
	return &AppError{
		Status:  http.StatusBadRequest,
		Message: message,
		Errors:  errors,
	}
}

// NewInternalServerError returns a 500 error
func NewInternalServerError(message string) *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

// HandleError is a helper to send standardized error responses
func HandleError(c *gin.Context, err interface{}) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Status, appErr)
		return
	}

	// Default to 500
	c.JSON(http.StatusInternalServerError, AppError{
		Status:  http.StatusInternalServerError,
		Message: "An unexpected error occurred",
	})
}
