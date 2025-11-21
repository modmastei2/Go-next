package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/modmastei2/Go-next/backend/config"
	"github.com/modmastei2/Go-next/backend/internal/handler"
	"github.com/modmastei2/Go-next/backend/internal/middleware"
	"github.com/modmastei2/Go-next/backend/internal/repository"
	"github.com/modmastei2/Go-next/backend/internal/usecase"
	"github.com/modmastei2/Go-next/backend/pkg/database"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.NewDatabase(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.MigrateDatabase(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Seed database
	if err := database.SeedDatabase(db); err != nil {
		log.Fatal("Failed to seed database:", err)
	}

	// Dependency Injection - Initialize repositories
	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Dependency Injection - Initialize usecases
	orderUsecase := usecase.NewOrderUsecase(orderRepo, productRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)

	// Dependency Injection - Initialize handlers
	orderHandler := handler.NewOrderHandler(orderUsecase)
	productHandler := handler.NewProductHandler(productUsecase)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Shop Order API",
	})

	// Apply global middleware
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())
	app.Use(middleware.RequestID())

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	// API routes
	api := app.Group("/api")

	// Product routes
	products := api.Group("/products")
	products.Get("/", productHandler.GetProducts)
	products.Get("/:id", productHandler.GetProduct)
	products.Post("/", productHandler.CreateProduct)
	products.Put("/:id", productHandler.UpdateProduct)
	products.Delete("/:id", productHandler.DeleteProduct)

	// Order routes
	orders := api.Group("/orders")
	orders.Get("/", orderHandler.GetOrders)
	orders.Get("/:id", orderHandler.GetOrder)
	orders.Post("/", orderHandler.CreateOrder)
	orders.Put("/:id/status", orderHandler.UpdateOrderStatus)
	orders.Delete("/:id", orderHandler.DeleteOrder)

	// Start server
	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", serverAddr)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
