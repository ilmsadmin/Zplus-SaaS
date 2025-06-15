'use client'

import { useEffect, ReactNode } from 'react'
import { useRouter, usePathname } from 'next/navigation'
import { useAuth } from '@/hooks/useAuth'

interface AuthGuardProps {
  children: ReactNode
  requireAuth?: boolean
  requiredRoles?: string[]
  requiredPermissions?: string[]
  redirectTo?: string
  allowUnauthenticated?: boolean
}

export function AuthGuard({
  children,
  requireAuth = true,
  requiredRoles = [],
  requiredPermissions = [],
  redirectTo,
  allowUnauthenticated = false
}: AuthGuardProps) {
  const { user, isAuthenticated, isLoading } = useAuth()
  const router = useRouter()
  const pathname = usePathname()

  useEffect(() => {
    if (isLoading) return

    // If authentication is required but user is not authenticated
    if (requireAuth && !isAuthenticated) {
      const loginUrl = redirectTo || `/login?redirect=${encodeURIComponent(pathname)}`
      router.push(loginUrl)
      return
    }

    // If unauthenticated access is not allowed and user is not authenticated
    if (!allowUnauthenticated && !isAuthenticated) {
      const loginUrl = redirectTo || `/login?redirect=${encodeURIComponent(pathname)}`
      router.push(loginUrl)
      return
    }

    // Check role requirements
    if (isAuthenticated && user && requiredRoles.length > 0) {
      const hasRequiredRole = requiredRoles.some(role => user.roles?.includes(role))
      if (!hasRequiredRole) {
        router.push('/unauthorized')
        return
      }
    }

    // Check permission requirements
    if (isAuthenticated && user && requiredPermissions.length > 0) {
      const hasRequiredPermission = requiredPermissions.some(permission => 
        user.permissions?.includes(permission)
      )
      if (!hasRequiredPermission) {
        router.push('/unauthorized')
        return
      }
    }
  }, [
    isLoading,
    isAuthenticated,
    user,
    requireAuth,
    requiredRoles,
    requiredPermissions,
    allowUnauthenticated,
    redirectTo,
    router,
    pathname
  ])

  // Show loading spinner while checking authentication
  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    )
  }

  // If authentication is required but user is not authenticated, don't render children
  if (requireAuth && !isAuthenticated) {
    return null
  }

  // If unauthenticated access is not allowed and user is not authenticated
  if (!allowUnauthenticated && !isAuthenticated) {
    return null
  }

  // If user doesn't have required roles, don't render children
  if (isAuthenticated && user && requiredRoles.length > 0) {
    const hasRequiredRole = requiredRoles.some(role => user.roles?.includes(role))
    if (!hasRequiredRole) {
      return null
    }
  }

  // If user doesn't have required permissions, don't render children
  if (isAuthenticated && user && requiredPermissions.length > 0) {
    const hasRequiredPermission = requiredPermissions.some(permission => 
      user.permissions?.includes(permission)
    )
    if (!hasRequiredPermission) {
      return null
    }
  }

  return <>{children}</>
}

// Specific guards for common use cases
export function SystemAdminGuard({ children }: { children: ReactNode }) {
  return (
    <AuthGuard requiredRoles={['system_admin']}>
      {children}
    </AuthGuard>
  )
}

export function TenantAdminGuard({ children }: { children: ReactNode }) {
  return (
    <AuthGuard requiredRoles={['tenant_admin']}>
      {children}
    </AuthGuard>
  )
}

export function AuthenticatedGuard({ children }: { children: ReactNode }) {
  return (
    <AuthGuard requireAuth={true}>
      {children}
    </AuthGuard>
  )
}

export function PublicGuard({ children }: { children: ReactNode }) {
  return (
    <AuthGuard requireAuth={false} allowUnauthenticated={true}>
      {children}
    </AuthGuard>
  )
}

export default AuthGuard
