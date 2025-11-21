# Frontend - Next.js Shop Application

A modern, type-safe frontend built with Next.js 15, TypeScript, and TailwindCSS.

## Features

- **Next.js 15**: Latest version with App Router
- **TypeScript**: Full type safety throughout the application
- **TailwindCSS**: Utility-first CSS framework for rapid UI development
- **HTTP Interceptor**: Centralized API client with request/response interceptors
- **Middleware**: Next.js middleware for request processing
- **Clean Architecture**: Well-organized component structure

## Project Structure

```
frontend/
├── app/
│   ├── components/          # Reusable React components
│   │   ├── OrderForm.tsx
│   │   ├── OrderList.tsx
│   │   └── ProductList.tsx
│   ├── orders/              # Shop orders page
│   │   └── page.tsx
│   ├── layout.tsx           # Root layout
│   ├── page.tsx             # Home page
│   └── globals.css          # Global styles
├── lib/
│   └── api/                 # API client and types
│       ├── http-client.ts   # HTTP client with interceptors
│       ├── index.ts         # API methods
│       └── types.ts         # TypeScript types
├── middleware.ts            # Next.js middleware
└── public/                  # Static assets
```

## Getting Started

### Prerequisites

- Node.js 18 or higher
- npm or yarn
- Backend API running on port 3001 (see backend README)

### Installation

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Copy the environment file:
```bash
cp .env.local.example .env.local
```

4. Start the development server:
```bash
npm run dev
```

The application will start on `http://localhost:3000`

### Building for Production

```bash
npm run build
npm start
```

## Features Demonstrated

### 1. HTTP Client with Interceptors

The application includes a custom HTTP client (`lib/api/http-client.ts`) that demonstrates:
- Request interceptors for adding headers, logging
- Response interceptors for error handling, logging
- Centralized API configuration
- Type-safe API methods

Example usage:
```typescript
import { api } from '@/lib/api';

// Get all products
const products = await api.products.getAll();

// Create an order
const order = await api.orders.create({
  customer_id: 1,
  items: [
    { product_id: 1, quantity: 2 }
  ]
});
```

### 2. Next.js Middleware

The middleware (`middleware.ts`) demonstrates:
- Request logging
- Adding custom headers
- Request/response manipulation
- Path-based middleware execution

### 3. Type-Safe API Integration

All API calls are fully typed using TypeScript interfaces defined in `lib/api/types.ts`:
- Product
- Order
- OrderItem
- CreateOrderRequest
- ApiResponse

### 4. Shop Order Page

The `/orders` page demonstrates:
- Product listing
- Shopping cart functionality
- Order creation
- Order management
- Real-time updates
- State management with React hooks

## Pages

### Home Page (`/`)
- Project overview
- Feature highlights
- Links to shop and API health check

### Orders Page (`/orders`)
- Browse products
- Add items to cart
- Create orders
- View and manage orders
- Update order status

## API Integration

The frontend connects to the backend API at `http://localhost:3001/api` (configurable via environment variables).

### API Endpoints Used

- `GET /products` - Get all products
- `GET /products/:id` - Get a specific product
- `POST /orders` - Create a new order
- `GET /orders` - Get all orders
- `PUT /orders/:id/status` - Update order status

## Environment Variables

Create a `.env.local` file with the following variables:

```
NEXT_PUBLIC_API_URL=http://localhost:3001/api
```

## Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm start` - Start production server
- `npm run lint` - Run ESLint

## Styling

The application uses TailwindCSS for styling with a modern, clean design:
- Responsive layout
- Hover effects and transitions
- Color-coded status indicators
- Consistent spacing and typography

## Type Safety

All components are written in TypeScript with proper type definitions:
- Props are typed
- API responses are typed
- State is typed
- Event handlers are typed

## Best Practices Demonstrated

1. **Component Organization**: Reusable components in separate files
2. **API Abstraction**: API calls abstracted into a service layer
3. **Error Handling**: Proper error handling with user feedback
4. **Loading States**: Loading indicators for async operations
5. **Form Validation**: Input validation for forms
6. **Code Splitting**: Automatic code splitting with Next.js
7. **SEO**: Proper page metadata and structure

## License

MIT
