/**
 * API client for shop orders and products
 * Demonstrates how to structure API calls with type safety
 */

import httpClient from './http-client';
import type { Product, Order, CreateOrderRequest, ApiResponse } from './types';

export const api = {
  // Product endpoints
  products: {
    getAll: async (limit = 10, offset = 0): Promise<Product[]> => {
      const response = await httpClient.get<ApiResponse<Product[]>>(
        `/products?limit=${limit}&offset=${offset}`
      );
      return response.data;
    },

    getById: async (id: number): Promise<Product> => {
      const response = await httpClient.get<ApiResponse<Product>>(`/products/${id}`);
      return response.data;
    },

    create: async (product: Omit<Product, 'id' | 'created_at' | 'updated_at'>): Promise<Product> => {
      const response = await httpClient.post<ApiResponse<Product>>('/products', product);
      return response.data;
    },

    update: async (id: number, product: Partial<Product>): Promise<Product> => {
      const response = await httpClient.put<ApiResponse<Product>>(`/products/${id}`, product);
      return response.data;
    },

    delete: async (id: number): Promise<void> => {
      await httpClient.delete(`/products/${id}`);
    },
  },

  // Order endpoints
  orders: {
    getAll: async (limit = 10, offset = 0): Promise<Order[]> => {
      const response = await httpClient.get<ApiResponse<Order[]>>(
        `/orders?limit=${limit}&offset=${offset}`
      );
      return response.data;
    },

    getById: async (id: number): Promise<Order> => {
      const response = await httpClient.get<ApiResponse<Order>>(`/orders/${id}`);
      return response.data;
    },

    create: async (order: CreateOrderRequest): Promise<Order> => {
      const response = await httpClient.post<ApiResponse<Order>>('/orders', order);
      return response.data;
    },

    updateStatus: async (id: number, status: string): Promise<void> => {
      await httpClient.put(`/orders/${id}/status`, { status });
    },

    delete: async (id: number): Promise<void> => {
      await httpClient.delete(`/orders/${id}`);
    },
  },
};

export default api;
