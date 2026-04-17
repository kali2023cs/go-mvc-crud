package services

import (
	"go-mvc-crud/dto"
	"go-mvc-crud/models"
	"go-mvc-crud/repositories"
)

type ProductService interface {
	GetAllProducts(name string, minPrice, maxPrice string) ([]dto.ProductResponse, error)
	GetProductByID(id uint) (*dto.ProductResponse, error)
	CreateProduct(input dto.CreateProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(id uint, input dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(id uint) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts(name string, minPrice, maxPrice string) ([]dto.ProductResponse, error) {
	products, err := s.repo.FindAll(name, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}

	var response []dto.ProductResponse
	for _, p := range products {
		response = append(response, s.mapToResponse(p))
	}
	return response, nil
}

func (s *productService) GetProductByID(id uint) (*dto.ProductResponse, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	res := s.mapToResponse(*product)
	return &res, nil
}

func (s *productService) CreateProduct(input dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	if err := s.repo.Create(&product); err != nil {
		return nil, err
	}

	res := s.mapToResponse(product)
	return &res, nil
}

func (s *productService) UpdateProduct(id uint, input dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Update(product, input); err != nil {
		return nil, err
	}

	res := s.mapToResponse(*product)
	return &res, nil
}

func (s *productService) DeleteProduct(id uint) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(product)
}

func (s *productService) mapToResponse(p models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
