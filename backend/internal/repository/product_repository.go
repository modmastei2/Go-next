package repository

import (
	"github.com/modmastei2/Go-next/backend/internal/domain"
	"gorm.io/gorm"
)

// ProductRepository defines the interface for product data access
type ProductRepository interface {
	Create(product *domain.Product) error
	GetByID(id uint) (*domain.Product, error)
	GetAll(limit, offset int) ([]domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uint) error
}

// productRepository implements ProductRepository interface
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create creates a new product
func (r *productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

// GetByID retrieves a product by ID
func (r *productRepository) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAll retrieves all products with pagination
func (r *productRepository) GetAll(limit, offset int) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

// Update updates an existing product
func (r *productRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

// Delete deletes a product by ID
func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}
