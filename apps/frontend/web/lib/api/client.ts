import { ApiResponse, PaginatedResponse } from '@/lib/types'
import { AuthService } from '@/lib/auth'

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'
const AUTH_SERVICE_URL = process.env.NEXT_PUBLIC_AUTH_URL || 'http://localhost:8081'

interface RequestConfig {
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'
  headers?: Record<string, string>
  body?: any
  tenantId?: string
  requireAuth?: boolean
}

class ApiClient {
  private baseURL: string
  private authURL: string
  private defaultHeaders: Record<string, string>

  constructor(baseURL: string = API_BASE_URL, authURL: string = AUTH_SERVICE_URL) {
    this.baseURL = baseURL
    this.authURL = authURL
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

  private getAuthHeaders(): Record<string, string> {
    return AuthService.getAuthHeaders()
  }

  async request<T>(endpoint: string, config: RequestConfig = {}): Promise<ApiResponse<T>> {
    const { method = 'GET', headers = {}, body, tenantId, requireAuth = true } = config
    
    // Use auth service URL for auth endpoints
    const isAuthEndpoint = endpoint.startsWith('/login') || endpoint.startsWith('/logout') || endpoint.startsWith('/refresh')
    const url = isAuthEndpoint ? `${this.authURL}${endpoint}` : `${this.baseURL}${endpoint}`
    
    const requestHeaders = {
      ...this.defaultHeaders,
      ...headers,
    }

    // Add authentication headers if required and available
    if (requireAuth && !isAuthEndpoint) {
      const authHeaders = this.getAuthHeaders()
      Object.assign(requestHeaders, authHeaders)
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
      
      // Handle 401 responses by attempting token refresh
      if (response.status === 401 && requireAuth && !isAuthEndpoint) {
        try {
          await AuthService.refreshToken()
          
          // Retry the request with new token
          const retryHeaders = {
            ...requestHeaders,
            ...this.getAuthHeaders()
          }
          
          const retryResponse = await fetch(url, {
            ...requestConfig,
            headers: retryHeaders
          })
          
          const retryData = await retryResponse.json()
          
          if (!retryResponse.ok) {
            throw new Error(retryData.message || `HTTP error! status: ${retryResponse.status}`)
          }
          
          return {
            data: retryData,
            success: true,
          }
        } catch (refreshError) {
          // Token refresh failed, redirect to login
          if (typeof window !== 'undefined') {
            window.location.href = '/login'
          }
          throw new Error('Authentication failed')
        }
      }
      
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

  async get<T>(endpoint: string, tenantId?: string, requireAuth: boolean = true): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'GET', tenantId, requireAuth })
  }

  async post<T>(endpoint: string, body: any, tenantId?: string, requireAuth: boolean = true): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'POST', body, tenantId, requireAuth })
  }

  async put<T>(endpoint: string, body: any, tenantId?: string, requireAuth: boolean = true): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'PUT', body, tenantId, requireAuth })
  }

  async delete<T>(endpoint: string, tenantId?: string, requireAuth: boolean = true): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'DELETE', tenantId, requireAuth })
  }

  async getPaginated<T>(
    endpoint: string, 
    params: { page?: number; limit?: number } = {},
    tenantId?: string,
    requireAuth: boolean = true
  ): Promise<ApiResponse<PaginatedResponse<T>>> {
    const queryParams = new URLSearchParams()
    if (params.page) queryParams.append('page', params.page.toString())
    if (params.limit) queryParams.append('limit', params.limit.toString())
    
    const url = queryParams.toString() ? `${endpoint}?${queryParams}` : endpoint
    return this.request<PaginatedResponse<T>>(url, { method: 'GET', tenantId, requireAuth })
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