package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"go-mvc-crud/config"
	"go-mvc-crud/dto"
	"go-mvc-crud/models"
	"go-mvc-crud/repositories"
	"go-mvc-crud/services"
	"go-mvc-crud/utils"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testHandler *ProductHandler

func setupTestDB() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{})
	config.DB = db
	utils.InitCustomValidators()

	// Initialize dependencies for testing
	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)
	testHandler = NewProductHandler(service)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetProducts(t *testing.T) {
	// Seed data
	config.DB.Create(&models.Product{Name: "Laptop", Price: 1000})
	config.DB.Create(&models.Product{Name: "Mouse", Price: 50})

	r := gin.Default()
	r.GET("/products", testHandler.GetProducts)

	t.Run("Get all products", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string][]dto.ProductResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.GreaterOrEqual(t, len(response["data"]), 2)
	})

	t.Run("Filter by name", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products?name=Lap", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var response map[string][]dto.ProductResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, 1, len(response["data"]))
		assert.Equal(t, "Laptop", response["data"][0].Name)
	})
}

func TestCreateProduct(t *testing.T) {
	r := gin.Default()
	r.POST("/products", testHandler.CreateProduct)

	t.Run("Success", func(t *testing.T) {
		body := map[string]interface{}{
			"name":  "Keyboard",
			"price": 75.5,
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Validation Failure - Missing Name", func(t *testing.T) {
		body := map[string]interface{}{
			"price": 75.5,
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "errors")
		assert.Contains(t, w.Body.String(), "This field is required")
	})

	t.Run("Validation Failure - Negative Price", func(t *testing.T) {
		body := map[string]interface{}{
			"name":  "Invalid",
			"price": -10,
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "must be greater than 0")
	})

	t.Run("Validation Failure - Reserved Name", func(t *testing.T) {
		body := map[string]interface{}{
			"name":  "Admin",
			"price": 100,
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "This name is reserved")
	})
}

func TestUpdateProduct(t *testing.T) {
	product := models.Product{Name: "Old Name", Price: 10}
	config.DB.Create(&product)

	r := gin.Default()
	r.PUT("/products/:id", testHandler.UpdateProduct)

	t.Run("Success", func(t *testing.T) {
		body := map[string]interface{}{
			"name":  "New Name",
			"price": 10.0,
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("PUT", "/products/"+strconv.Itoa(int(product.ID)), bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestDeleteProduct(t *testing.T) {
	product := models.Product{Name: "Delete Me", Price: 10}
	config.DB.Create(&product)

	r := gin.Default()
	r.DELETE("/products/:id", testHandler.DeleteProduct)

	req, _ := http.NewRequest("DELETE", "/products/"+strconv.Itoa(int(product.ID)), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
