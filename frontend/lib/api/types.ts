/**
 * Type definitions for API entities
 */

export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  stock: number;
  created_at: string;
  updated_at: string;
}

export interface Customer {
  id: number;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface OrderItem {
  id?: number;
  order_id?: number;
  product_id: number;
  product?: Product;
  quantity: number;
  price?: number;
}

export interface Order {
  id: number;
  customer_id: number;
  customer?: Customer;
  items: OrderItem[];
  total: number;
  status: 'pending' | 'processing' | 'completed' | 'cancelled';
  created_at: string;
  updated_at: string;
}

export interface CreateOrderRequest {
  customer_id: number;
  items: {
    product_id: number;
    quantity: number;
  }[];
}

export interface ApiResponse<T> {
  data: T;
  message?: string;
}
