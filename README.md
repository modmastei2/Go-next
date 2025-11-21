# Go-next

Full-stack application demonstrating Next.js frontend with Golang backend integration.

## Project Structure

This project is split into two main parts:

### Backend (Golang + Fiber + GORM)
- **Framework**: Go with Fiber web framework
- **ORM**: GORM for database operations
- **Architecture**: Clean Architecture with dependency injection
- **Database**: SQLite (easily configurable for PostgreSQL/MySQL)
- **Features**: RESTful API, middleware, CORS, request logging

[See backend README](./backend/README.md) for detailed documentation.

### Frontend (Next.js + TypeScript + TailwindCSS)
- **Framework**: Next.js 15 with App Router
- **Language**: TypeScript for type safety
- **Styling**: TailwindCSS
- **Features**: HTTP interceptor, middleware, shopping cart, order management

[See frontend README](./frontend/README.md) for detailed documentation.

## Quick Start

### Prerequisites
- Go 1.21 or higher
- Node.js 18 or higher
- npm or yarn

### Backend Setup

1. Navigate to backend directory:
```bash
cd backend
```

2. Install Go dependencies:
```bash
go mod download
```

3. Run the backend server:
```bash
go run cmd/api/main.go
```

The backend API will start on `http://localhost:3001`

### Frontend Setup

1. Navigate to frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Create environment file:
```bash
cp .env.local.example .env.local
```

4. Run the development server:
```bash
npm run dev
```

The frontend will start on `http://localhost:3000`

## Features Demonstrated

### Backend (Clean Architecture)

1. **Domain Layer**: Business entities (Order, Product, Customer, OrderItem)
2. **Repository Layer**: Data access with GORM
3. **Usecase Layer**: Business logic and validation
4. **Handler Layer**: HTTP request handlers
5. **Middleware**: Custom middleware for logging, CORS, recovery, request ID
6. **Dependency Injection**: Proper DI pattern implementation

### Frontend (Modern React/Next.js)

1. **HTTP Client**: Custom client with request/response interceptors
2. **Middleware**: Next.js middleware for request processing
3. **Type Safety**: Full TypeScript implementation
4. **State Management**: React hooks (useState, useEffect)
5. **API Integration**: Type-safe API calls
6. **UI Components**: Reusable components with TailwindCSS

## API Endpoints

### Products
- `GET /api/products` - List all products
- `GET /api/products/:id` - Get product by ID
- `POST /api/products` - Create new product
- `PUT /api/products/:id` - Update product
- `DELETE /api/products/:id` - Delete product

### Orders
- `GET /api/orders` - List all orders
- `GET /api/orders/:id` - Get order by ID
- `POST /api/orders` - Create new order
- `PUT /api/orders/:id/status` - Update order status
- `DELETE /api/orders/:id` - Delete order

## Technologies

### Backend
- Go
- Fiber (Web Framework)
- GORM (ORM)
- SQLite/PostgreSQL/MySQL

### Frontend
- Next.js 15
- TypeScript
- TailwindCSS
- React 19

## Project Goals

This project demonstrates:
- How to structure a Next.js + Golang application
- Clean architecture implementation in Go
- Dependency injection patterns
- Middleware creation (both backend and frontend)
- HTTP interceptors for API calls
- Type-safe API integration
- Modern UI with TailwindCSS
- Full CRUD operations

## License

MIT
