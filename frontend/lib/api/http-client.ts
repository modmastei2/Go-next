/**
 * HTTP Client with interceptor functionality
 * Demonstrates how to create a centralized API client with request/response interceptors
 */

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001/api';

export interface ApiError {
  message: string;
  status?: number;
}

export interface RequestInterceptor {
  onRequest?: (config: RequestInit) => RequestInit | Promise<RequestInit>;
  onError?: (error: ApiError) => void | Promise<void>;
}

export interface ResponseInterceptor {
  onResponse?: (response: Response) => Response | Promise<Response>;
  onError?: (error: ApiError) => void | Promise<void>;
}

class HttpClient {
  private baseURL: string;
  private requestInterceptors: RequestInterceptor[] = [];
  private responseInterceptors: ResponseInterceptor[] = [];

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  // Add request interceptor
  addRequestInterceptor(interceptor: RequestInterceptor) {
    this.requestInterceptors.push(interceptor);
  }

  // Add response interceptor
  addResponseInterceptor(interceptor: ResponseInterceptor) {
    this.responseInterceptors.push(interceptor);
  }

  // Process request through interceptors
  private async processRequest(url: string, config: RequestInit): Promise<RequestInit> {
    let processedConfig: RequestInit = {
      ...config,
      headers: {
        'Content-Type': 'application/json',
        ...(config.headers as Record<string, string>),
      },
    };

    for (const interceptor of this.requestInterceptors) {
      if (interceptor.onRequest) {
        try {
          processedConfig = await interceptor.onRequest(processedConfig);
        } catch (error) {
          if (interceptor.onError) {
            await interceptor.onError(error as ApiError);
          }
          throw error;
        }
      }
    }

    return processedConfig;
  }

  // Process response through interceptors
  private async processResponse(response: Response): Promise<Response> {
    let processedResponse = response;

    for (const interceptor of this.responseInterceptors) {
      if (interceptor.onResponse) {
        try {
          processedResponse = await interceptor.onResponse(processedResponse);
        } catch (error) {
          if (interceptor.onError) {
            await interceptor.onError(error as ApiError);
          }
          throw error;
        }
      }
    }

    return processedResponse;
  }

  // Generic request method
  private async request<T>(endpoint: string, config: RequestInit = {}): Promise<T> {
    const url = `${this.baseURL}${endpoint}`;

    try {
      // Process request through interceptors
      const processedConfig = await this.processRequest(url, config);

      // Make the request
      let response = await fetch(url, processedConfig);

      // Process response through interceptors
      response = await this.processResponse(response);

      // Handle errors
      if (!response.ok) {
        const error: ApiError = {
          message: `HTTP error! status: ${response.status}`,
          status: response.status,
        };
        throw error;
      }

      // Parse response
      const data = await response.json();
      return data;
    } catch (error) {
      // Call error interceptors
      const apiError = error as ApiError;
      for (const interceptor of this.responseInterceptors) {
        if (interceptor.onError) {
          await interceptor.onError(apiError);
        }
      }
      throw error;
    }
  }

  // HTTP methods
  get<T>(endpoint: string): Promise<T> {
    return this.request<T>(endpoint, { method: 'GET' });
  }

  post<T>(endpoint: string, data?: unknown): Promise<T> {
    return this.request<T>(endpoint, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  put<T>(endpoint: string, data?: unknown): Promise<T> {
    return this.request<T>(endpoint, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  delete<T>(endpoint: string): Promise<T> {
    return this.request<T>(endpoint, { method: 'DELETE' });
  }
}

// Create a singleton instance
export const httpClient = new HttpClient(API_BASE_URL);

// Add default request interceptor for logging
httpClient.addRequestInterceptor({
  onRequest: (config) => {
    console.log(`[HTTP] Request:`, config);
    return config;
  },
  onError: (error) => {
    console.error('[HTTP] Request Error:', error);
  },
});

// Add default response interceptor for logging and error handling
httpClient.addResponseInterceptor({
  onResponse: (response) => {
    console.log(`[HTTP] Response:`, response.status, response.statusText);
    return response;
  },
  onError: (error) => {
    console.error('[HTTP] Response Error:', error);
  },
});

// Example: Add authentication interceptor (can be uncommented when auth is implemented)
// httpClient.addRequestInterceptor({
//   onRequest: (config) => {
//     const token = localStorage.getItem('token');
//     if (token) {
//       config.headers = {
//         ...config.headers,
//         Authorization: `Bearer ${token}`,
//       };
//     }
//     return config;
//   },
// });

export default httpClient;
