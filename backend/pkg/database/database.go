package database

import (
	"fmt"
	"log"

	"github.com/modmastei2/Go-next/backend/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config holds database configuration
type Config struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// NewDatabase creates a new database connection
func NewDatabase(config *Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// For demonstration, we'll use SQLite
	// In production, you'd use PostgreSQL or MySQL
	dsn := config.Database
	if dsn == "" {
		dsn = "shop.db"
	}

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}

// MigrateDatabase runs database migrations
func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.Customer{},
		&domain.Product{},
		&domain.Order{},
		&domain.OrderItem{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migrated successfully")
	return nil
}

// SeedDatabase seeds the database with initial data
func SeedDatabase(db *gorm.DB) error {
	// Check if data already exists
	var count int64
	db.Model(&domain.Product{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	// Seed customers
	customers := []domain.Customer{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Smith", Email: "jane@example.com"},
	}

	for _, customer := range customers {
		if err := db.Create(&customer).Error; err != nil {
			return err
		}
	}

	// Seed products
	products := []domain.Product{
		{Name: "Laptop", Description: "High-performance laptop", Price: 999.99, Stock: 10},
		{Name: "Mouse", Description: "Wireless mouse", Price: 29.99, Stock: 50},
		{Name: "Keyboard", Description: "Mechanical keyboard", Price: 79.99, Stock: 30},
		{Name: "Monitor", Description: "27-inch 4K monitor", Price: 399.99, Stock: 15},
		{Name: "Headphones", Description: "Noise-cancelling headphones", Price: 199.99, Stock: 25},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			return err
		}
	}

	log.Println("Database seeded successfully")
	return nil
}
