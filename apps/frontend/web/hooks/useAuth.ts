'use client'

import { useState, useEffect, useCallback } from 'react'
import { useRouter } from 'next/navigation'
import { AuthService, User, LoginResponse } from '@/lib/auth'

export interface UseAuthReturn {
  user: User | null
  isAuthenticated: boolean
  isLoading: boolean
  login: (credentials: {
    email: string
    password: string
    tenant_slug: string
  }) => Promise<LoginResponse>
  logout: () => Promise<void>
  refreshToken: () => Promise<void>
  hasRole: (role: string) => boolean
  hasPermission: (permission: string) => boolean
  isSystemAdmin: boolean
  isTenantAdmin: boolean
  userDisplayName: string
  userInitials: string
}

export function useAuth(): UseAuthReturn {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const router = useRouter()

  // Load user data on mount
  useEffect(() => {
    const loadUser = () => {
      try {
        const storedUser = AuthService.getUser()
        setUser(storedUser)
      } catch (error) {
        console.error('Failed to load user data:', error)
        setUser(null)
      } finally {
        setIsLoading(false)
      }
    }

    loadUser()
  }, [])

  const login = useCallback(async (credentials: {
    email: string
    password: string
    tenant_slug: string
  }): Promise<LoginResponse> => {
    setIsLoading(true)
    try {
      const response = await AuthService.login(credentials)
      setUser(response.user)
      return response
    } catch (error) {
      console.error('Login failed:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }, [])

  const logout = useCallback(async (): Promise<void> => {
    setIsLoading(true)
    try {
      await AuthService.logout()
      setUser(null)
      router.push('/login')
    } catch (error) {
      console.error('Logout failed:', error)
      // Still clear local state even if API call fails
      setUser(null)
      router.push('/login')
    } finally {
      setIsLoading(false)
    }
  }, [router])

  const refreshToken = useCallback(async (): Promise<void> => {
    try {
      const response = await AuthService.refreshToken()
      setUser(response.user)
    } catch (error) {
      console.error('Token refresh failed:', error)
      setUser(null)
      router.push('/login')
      throw error
    }
  }, [router])

  const hasRole = useCallback((role: string): boolean => {
    return user?.roles?.includes(role) ?? false
  }, [user])

  const hasPermission = useCallback((permission: string): boolean => {
    return user?.permissions?.includes(permission) ?? false
  }, [user])

  const isSystemAdmin = user?.is_admin === true && user?.tenant_id === 'system'
  const isTenantAdmin = user?.roles?.includes('tenant_admin') ?? false
  const isAuthenticated = !!user && AuthService.isAuthenticated()

  const userDisplayName = user 
    ? `${user.first_name} ${user.last_name}`.trim() || user.email
    : 'Unknown User'

  const userInitials = user
    ? (user.first_name?.[0] || '') + (user.last_name?.[0] || '') || user.email[0]
    : 'U'

  return {
    user,
    isAuthenticated,
    isLoading,
    login,
    logout,
    refreshToken,
    hasRole,
    hasPermission,
    isSystemAdmin,
    isTenantAdmin,
    userDisplayName,
    userInitials: userInitials.toUpperCase()
  }
}

export default useAuth
