// Auth utility functions for managing authentication state
export interface User {
  id: string
  tenant_id: string
  email: string
  first_name: string
  last_name: string
  roles: string[]
  is_admin: boolean
  status: string
  permissions: string[]
}

export interface AuthTokens {
  token: string
  refresh_token: string
  expires_in: number
}

export interface LoginResponse {
  token: string
  refresh_token: string
  user: User
  expires_in: number
}

const AUTH_SERVICE_URL = process.env.NEXT_PUBLIC_AUTH_URL || 'http://localhost:8081'

export class AuthService {
  // Local storage keys
  private static readonly TOKEN_KEY = 'zplus_token'
  private static readonly REFRESH_TOKEN_KEY = 'zplus_refresh_token'
  private static readonly USER_KEY = 'zplus_user'

  // Get stored authentication token
  static getToken(): string | null {
    if (typeof window === 'undefined') return null
    return localStorage.getItem(this.TOKEN_KEY)
  }

  // Get stored refresh token
  static getRefreshToken(): string | null {
    if (typeof window === 'undefined') return null
    return localStorage.getItem(this.REFRESH_TOKEN_KEY)
  }

  // Get stored user data
  static getUser(): User | null {
    if (typeof window === 'undefined') return null
    
    try {
      const userStr = localStorage.getItem(this.USER_KEY)
      return userStr ? JSON.parse(userStr) : null
    } catch (error) {
      console.error('Error parsing stored user data:', error)
      return null
    }
  }

  // Check if user is authenticated
  static isAuthenticated(): boolean {
    const token = this.getToken()
    const user = this.getUser()
    return !!(token && user)
  }

  // Store authentication data
  static setAuthData(data: LoginResponse): void {
    if (typeof window === 'undefined') return
    
    localStorage.setItem(this.TOKEN_KEY, data.token)
    localStorage.setItem(this.REFRESH_TOKEN_KEY, data.refresh_token)
    localStorage.setItem(this.USER_KEY, JSON.stringify(data.user))
  }

  // Clear authentication data
  static clearAuthData(): void {
    if (typeof window === 'undefined') return
    
    localStorage.removeItem(this.TOKEN_KEY)
    localStorage.removeItem(this.REFRESH_TOKEN_KEY)
    localStorage.removeItem(this.USER_KEY)
  }

  // Login
  static async login(credentials: {
    email: string
    password: string
    tenant_slug: string
  }): Promise<LoginResponse> {
    const response = await fetch(`${AUTH_SERVICE_URL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials)
    })

    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.message || data.error || 'Login failed')
    }

    // Store authentication data
    this.setAuthData(data)

    return data
  }

  // Logout
  static async logout(): Promise<void> {
    const token = this.getToken()
    
    if (token) {
      try {
        await fetch(`${AUTH_SERVICE_URL}/logout`, {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
          }
        })
      } catch (error) {
        console.error('Logout API call failed:', error)
      }
    }

    // Always clear local data regardless of API success
    this.clearAuthData()
  }

  // Refresh token
  static async refreshToken(): Promise<LoginResponse> {
    const refreshToken = this.getRefreshToken()
    
    if (!refreshToken) {
      throw new Error('No refresh token available')
    }

    const response = await fetch(`${AUTH_SERVICE_URL}/refresh`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        refresh_token: refreshToken
      })
    })

    const data = await response.json()

    if (!response.ok) {
      // Clear auth data if refresh fails
      this.clearAuthData()
      throw new Error(data.message || data.error || 'Token refresh failed')
    }

    // Update stored authentication data
    this.setAuthData(data)

    return data
  }

  // Get authorization headers for API requests
  static getAuthHeaders(): Record<string, string> {
    const token = this.getToken()
    
    if (!token) {
      return {}
    }

    return {
      'Authorization': `Bearer ${token}`
    }
  }

  // Check if user has specific role
  static hasRole(role: string): boolean {
    const user = this.getUser()
    return user?.roles?.includes(role) ?? false
  }

  // Check if user has specific permission
  static hasPermission(permission: string): boolean {
    const user = this.getUser()
    return user?.permissions?.includes(permission) ?? false
  }

  // Check if user is system admin
  static isSystemAdmin(): boolean {
    const user = this.getUser()
    return user?.is_admin === true && user?.tenant_id === 'system'
  }

  // Check if user is tenant admin
  static isTenantAdmin(): boolean {
    const user = this.getUser()
    return user?.roles?.includes('tenant_admin') ?? false
  }

  // Get user's full name
  static getUserDisplayName(): string {
    const user = this.getUser()
    if (!user) return 'Unknown User'
    
    return `${user.first_name} ${user.last_name}`.trim() || user.email
  }

  // Get user's initials for avatar
  static getUserInitials(): string {
    const user = this.getUser()
    if (!user) return 'U'
    
    const firstInitial = user.first_name?.[0] || ''
    const lastInitial = user.last_name?.[0] || ''
    
    return (firstInitial + lastInitial).toUpperCase() || user.email[0].toUpperCase()
  }

  // Setup automatic token refresh
  static setupTokenRefresh(): void {
    if (typeof window === 'undefined') return

    // Refresh token every 30 minutes
    setInterval(async () => {
      if (this.isAuthenticated()) {
        try {
          await this.refreshToken()
        } catch (error) {
          console.error('Auto token refresh failed:', error)
          // Redirect to login on refresh failure
          window.location.href = '/login'
        }
      }
    }, 30 * 60 * 1000) // 30 minutes
  }
}

// Initialize token refresh on module load
if (typeof window !== 'undefined') {
  AuthService.setupTokenRefresh()
}

export default AuthService
