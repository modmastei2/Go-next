# Backend - Shop Order API

A RESTful API built with Go, Fiber framework, and GORM using Clean Architecture principles.

## Features

- **Clean Architecture**: Separation of concerns with domain, repository, usecase, and handler layers
- **Dependency Injection**: Proper DI implementation for better testability
- **Middleware**: Custom middleware for logging, CORS, request ID, and panic recovery
- **GORM**: Database ORM with migrations and seeding
- **Fiber**: Fast and lightweight web framework

## Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── domain/                  # Domain entities and models
│   │   └── order.go
│   ├── repository/              # Data access layer
│   │   ├── order_repository.go
│   │   └── product_repository.go
│   ├── usecase/                 # Business logic layer
│   │   ├── order_usecase.go
│   │   └── product_usecase.go
│   ├── handler/                 # HTTP handlers
│   │   ├── order_handler.go
│   │   └── product_handler.go
│   └── middleware/              # Custom middleware
│       └── middleware.go
├── pkg/
│   └── database/                # Database utilities
│       └── database.go
├── config/                      # Configuration management
│   └── config.go
└── go.mod
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- SQLite (or PostgreSQL/MySQL for production)

### Installation

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Copy the environment file:
```bash
cp .env.example .env
```

4. Run the application:
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:3001`

### Building

```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

## API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### Products
- `GET /api/products` - Get all products (with pagination)
- `GET /api/products/:id` - Get a product by ID
- `POST /api/products` - Create a new product
- `PUT /api/products/:id` - Update a product
- `DELETE /api/products/:id` - Delete a product

### Orders
- `GET /api/orders` - Get all orders (with pagination)
- `GET /api/orders/:id` - Get an order by ID
- `POST /api/orders` - Create a new order
- `PUT /api/orders/:id/status` - Update order status
- `DELETE /api/orders/:id` - Delete an order

## Example Requests

### Create a Product
```bash
curl -X POST http://localhost:3001/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "price": 999.99,
    "stock": 10
  }'
```

### Create an Order
```bash
curl -X POST http://localhost:3001/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "items": [
      {
        "product_id": 1,
        "quantity": 2
      }
    ]
  }'
```

### Get All Orders
```bash
curl http://localhost:3001/api/orders?limit=10&offset=0
```

## Clean Architecture Layers

### Domain Layer (`internal/domain`)
- Contains business entities and models
- No dependencies on other layers
- Represents the core business logic

### Repository Layer (`internal/repository`)
- Handles data persistence
- Abstracts database operations
- Implements repository interfaces

### Usecase Layer (`internal/usecase`)
- Contains business logic
- Orchestrates data flow between layers
- Implements business rules and validations

### Handler Layer (`internal/handler`)
- HTTP request handlers
- Input validation
- Response formatting

## Middleware

The application includes several middleware components:

1. **Logger**: Logs all HTTP requests with method, path, status, and duration
2. **CORS**: Handles Cross-Origin Resource Sharing
3. **RequestID**: Adds unique request ID to each request
4. **Recover**: Recovers from panics and returns proper error responses

## Database

The application uses SQLite by default for easy setup. The database is automatically:
- Created on first run
- Migrated with the latest schema
- Seeded with sample data

To use PostgreSQL or MySQL, update the configuration in `.env` file.

## Configuration

Configuration is managed through environment variables. See `.env.example` for available options.

## License

MIT
