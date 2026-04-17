package repositories

import (
	"go-mvc-crud/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(name string, minPrice, maxPrice string) ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product, data interface{}) error
	Delete(product *models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll(name string, minPrice, maxPrice string) ([]models.Product, error) {
	var products []models.Product
	query := r.db

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	err := query.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product, data interface{}) error {
	return r.db.Model(product).Updates(data).Error
}

func (r *productRepository) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}
