import { ApiResponse, PaginatedResponse } from '@/lib/types'

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'

interface RequestConfig {
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'
  headers?: Record<string, string>
  body?: any
  tenantId?: string
}

class ApiClient {
  private baseURL: string
  private defaultHeaders: Record<string, string>

  constructor(baseURL: string = API_BASE_URL) {
    this.baseURL = baseURL
    this.defaultHeaders = {
      'Content-Type': 'application/json',
    }
  }

  setTenant(tenantId: string) {
    this.defaultHeaders['X-Tenant-ID'] = tenantId
  }

  setAuthToken(token: string) {
    this.defaultHeaders['Authorization'] = `Bearer ${token}`
  }

  async request<T>(endpoint: string, config: RequestConfig = {}): Promise<ApiResponse<T>> {
    const { method = 'GET', headers = {}, body, tenantId } = config
    
    const url = `${this.baseURL}${endpoint}`
    const requestHeaders = {
      ...this.defaultHeaders,
      ...headers,
    }

    // Add tenant ID to headers if provided
    if (tenantId) {
      requestHeaders['X-Tenant-ID'] = tenantId
    }

    const requestConfig: RequestInit = {
      method,
      headers: requestHeaders,
    }

    if (body && method !== 'GET') {
      requestConfig.body = JSON.stringify(body)
    }

    try {
      const response = await fetch(url, requestConfig)
      const data = await response.json()

      if (!response.ok) {
        throw new Error(data.message || `HTTP error! status: ${response.status}`)
      }

      return {
        data,
        success: true,
      }
    } catch (error) {
      console.error('API request failed:', error)
      return {
        data: null as unknown as T,
        success: false,
        error: error instanceof Error ? error.message : 'Unknown error occurred',
      }
    }
  }

  async get<T>(endpoint: string, tenantId?: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'GET', tenantId })
  }

  async post<T>(endpoint: string, body: any, tenantId?: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'POST', body, tenantId })
  }

  async put<T>(endpoint: string, body: any, tenantId?: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'PUT', body, tenantId })
  }

  async delete<T>(endpoint: string, tenantId?: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'DELETE', tenantId })
  }

  async getPaginated<T>(
    endpoint: string, 
    params: { page?: number; limit?: number } = {},
    tenantId?: string
  ): Promise<ApiResponse<PaginatedResponse<T>>> {
    const queryParams = new URLSearchParams()
    if (params.page) queryParams.append('page', params.page.toString())
    if (params.limit) queryParams.append('limit', params.limit.toString())
    
    const url = queryParams.toString() ? `${endpoint}?${queryParams}` : endpoint
    return this.request<PaginatedResponse<T>>(url, { method: 'GET', tenantId })
  }
}

// Create a singleton instance
export const apiClient = new ApiClient()

// Utility functions for common API patterns
export const withTenant = (tenantSlug: string) => {
  return {
    get: <T>(endpoint: string) => apiClient.get<T>(endpoint, tenantSlug),
    post: <T>(endpoint: string, body: any) => apiClient.post<T>(endpoint, body, tenantSlug),
    put: <T>(endpoint: string, body: any) => apiClient.put<T>(endpoint, body, tenantSlug),
    delete: <T>(endpoint: string) => apiClient.delete<T>(endpoint, tenantSlug),
    getPaginated: <T>(endpoint: string, params?: { page?: number; limit?: number }) => 
      apiClient.getPaginated<T>(endpoint, params, tenantSlug),
  }
}

export default apiClient