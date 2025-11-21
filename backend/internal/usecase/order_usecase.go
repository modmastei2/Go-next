package usecase

import (
	"errors"
	"github.com/modmastei2/Go-next/backend/internal/domain"
	"github.com/modmastei2/Go-next/backend/internal/repository"
	"time"
)

// OrderUsecase defines the interface for order business logic
type OrderUsecase interface {
	CreateOrder(req *domain.CreateOrderRequest) (*domain.Order, error)
	GetOrder(id uint) (*domain.Order, error)
	GetOrders(limit, offset int) ([]domain.Order, error)
	UpdateOrderStatus(id uint, status string) error
	DeleteOrder(id uint) error
}

// orderUsecase implements OrderUsecase interface
type orderUsecase struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

// NewOrderUsecase creates a new order usecase
func NewOrderUsecase(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

// CreateOrder creates a new order with validation
func (u *orderUsecase) CreateOrder(req *domain.CreateOrderRequest) (*domain.Order, error) {
	// Validate and calculate total
	var total float64
	var orderItems []domain.OrderItem

	for _, item := range req.Items {
		product, err := u.productRepo.GetByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		orderItem := domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
		orderItems = append(orderItems, orderItem)
		total += product.Price * float64(item.Quantity)

		// Update stock
		product.Stock -= item.Quantity
		u.productRepo.Update(product)
	}

	// Create order
	order := &domain.Order{
		CustomerID: req.CustomerID,
		Items:      orderItems,
		Total:      total,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := u.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetOrder retrieves an order by ID
func (u *orderUsecase) GetOrder(id uint) (*domain.Order, error) {
	return u.orderRepo.GetByID(id)
}

// GetOrders retrieves all orders with pagination
func (u *orderUsecase) GetOrders(limit, offset int) ([]domain.Order, error) {
	if limit <= 0 {
		limit = 10
	}
	return u.orderRepo.GetAll(limit, offset)
}

// UpdateOrderStatus updates the status of an order
func (u *orderUsecase) UpdateOrderStatus(id uint, status string) error {
	order, err := u.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	validStatuses := map[string]bool{
		"pending":    true,
		"processing": true,
		"completed":  true,
		"cancelled":  true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	order.Status = status
	order.UpdatedAt = time.Now()
	return u.orderRepo.Update(order)
}

// DeleteOrder deletes an order
func (u *orderUsecase) DeleteOrder(id uint) error {
	return u.orderRepo.Delete(id)
}
