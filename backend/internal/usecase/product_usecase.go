package usecase

import (
	"github.com/modmastei2/Go-next/backend/internal/domain"
	"github.com/modmastei2/Go-next/backend/internal/repository"
)

// ProductUsecase defines the interface for product business logic
type ProductUsecase interface {
	CreateProduct(product *domain.Product) error
	GetProduct(id uint) (*domain.Product, error)
	GetProducts(limit, offset int) ([]domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uint) error
}

// productUsecase implements ProductUsecase interface
type productUsecase struct {
	productRepo repository.ProductRepository
}

// NewProductUsecase creates a new product usecase
func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

// CreateProduct creates a new product
func (u *productUsecase) CreateProduct(product *domain.Product) error {
	return u.productRepo.Create(product)
}

// GetProduct retrieves a product by ID
func (u *productUsecase) GetProduct(id uint) (*domain.Product, error) {
	return u.productRepo.GetByID(id)
}

// GetProducts retrieves all products with pagination
func (u *productUsecase) GetProducts(limit, offset int) ([]domain.Product, error) {
	if limit <= 0 {
		limit = 10
	}
	return u.productRepo.GetAll(limit, offset)
}

// UpdateProduct updates an existing product
func (u *productUsecase) UpdateProduct(product *domain.Product) error {
	return u.productRepo.Update(product)
}

// DeleteProduct deletes a product
func (u *productUsecase) DeleteProduct(id uint) error {
	return u.productRepo.Delete(id)
}
