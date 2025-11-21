package repository

import (
	"github.com/modmastei2/Go-next/backend/internal/domain"
	"gorm.io/gorm"
)

// OrderRepository defines the interface for order data access
type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id uint) (*domain.Order, error)
	GetAll(limit, offset int) ([]domain.Order, error)
	Update(order *domain.Order) error
	Delete(id uint) error
}

// orderRepository implements OrderRepository interface
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// Create creates a new order
func (r *orderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

// GetByID retrieves an order by ID
func (r *orderRepository) GetByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Customer").Preload("Items.Product").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// GetAll retrieves all orders with pagination
func (r *orderRepository) GetAll(limit, offset int) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("Customer").Preload("Items.Product").
		Limit(limit).Offset(offset).Find(&orders).Error
	return orders, err
}

// Update updates an existing order
func (r *orderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

// Delete deletes an order by ID
func (r *orderRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Order{}, id).Error
}
