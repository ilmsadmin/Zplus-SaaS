'use client'

import { useState, useEffect } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import { AuthService } from '@/lib/auth'

interface LoginCredentials {
  email: string
  password: string
  tenant_slug: string
}

interface User {
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

interface LoginResponse {
  token: string
  refresh_token: string
  user: User
  expires_in: number
}

type TenantType = 'system' | 'demo-corp' | 'customer'

const DEMO_CREDENTIALS = {
  system: { email: 'admin@zplus.com', password: 'admin123' },
  'demo-corp': { email: 'admin@demo-corp.zplus.com', password: 'demo123' },
  customer: { email: 'john@demo-corp.zplus.com', password: 'user123' }
}

const TENANT_MAP = {
  system: 'system',
  'demo-corp': 'demo-corp',
  customer: 'demo-corp' // Customer users belong to demo-corp tenant
}

export default function LoginPage() {
  const [selectedTenant, setSelectedTenant] = useState<TenantType>('system')
  const [credentials, setCredentials] = useState<LoginCredentials>({
    email: DEMO_CREDENTIALS.system.email,
    password: DEMO_CREDENTIALS.system.password,
    tenant_slug: TENANT_MAP.system
  })
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')

  const router = useRouter()
  const searchParams = useSearchParams()

  useEffect(() => {
    // Check if user is already logged in
    const token = localStorage.getItem('zplus_token')
    const userStr = localStorage.getItem('zplus_user')
    
    if (token && userStr) {
      try {
        const user = JSON.parse(userStr)
        redirectUser(user)
      } catch (error) {
        // Clear invalid stored data
        localStorage.removeItem('zplus_token')
        localStorage.removeItem('zplus_refresh_token')
        localStorage.removeItem('zplus_user')
      }
    }
  }, [])

  const redirectUser = (user: User) => {
    const redirect = searchParams.get('redirect')
    
    if (redirect) {
      router.push(redirect)
      return
    }

    if (user.is_admin && user.tenant_id === 'system') {
      router.push('/admin')
    } else if (user.roles.includes('tenant_admin')) {
      router.push(`/tenant/${user.tenant_id}/admin`)
    } else {
      // Customer user - redirect to CRM module
      router.push(`/tenant/${user.tenant_id}/crm`)
    }
  }

  const handleTenantChange = (tenantType: TenantType) => {
    setSelectedTenant(tenantType)
    const demoCredentials = DEMO_CREDENTIALS[tenantType]
    const tenantSlug = TENANT_MAP[tenantType]
    
    setCredentials({
      email: demoCredentials.email,
      password: demoCredentials.password,
      tenant_slug: tenantSlug
    })
    clearMessages()
  }

  const clearMessages = () => {
    setError('')
    setSuccess('')
  }

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setCredentials(prev => ({
      ...prev,
      [name]: value
    }))
    clearMessages()
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    
    if (isLoading) return
    
    clearMessages()
    
    // Client-side validation
    if (!credentials.email || !credentials.password) {
      setError('Please fill in all fields')
      return
    }
    
    setIsLoading(true)
    
    try {
      const loginData = await AuthService.login({
        email: credentials.email.trim(),
        password: credentials.password,
        tenant_slug: credentials.tenant_slug
      })
      
      setSuccess('Login successful! Redirecting...')
      
      // Redirect after short delay
      setTimeout(() => {
        redirectUser(loginData.user)
      }, 1500)
    } catch (error) {
      console.error('Login error:', error)
      setError(error instanceof Error ? error.message : 'Network error. Please check if the auth service is running.')
    } finally {
      setIsLoading(false)
    }
  }

  const getEmailPlaceholder = () => {
    const placeholders = {
      system: 'admin@zplus.com',
      'demo-corp': 'admin@demo-corp.zplus.com',
      customer: 'john@demo-corp.zplus.com'
    }
    return placeholders[selectedTenant]
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-600 to-indigo-700 flex items-center justify-center px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full bg-white rounded-xl shadow-2xl p-8">
        {/* Logo */}
        <div className="text-center mb-8">
          <div className="flex items-center justify-center mb-4">
            <div className="bg-blue-600 text-white p-3 rounded-lg">
              <svg className="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
          </div>
          <h1 className="text-3xl font-bold text-gray-900">Zplus SaaS</h1>
          <p className="text-gray-600 mt-2">Multi-tenant Business Platform</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Tenant Selection */}
          <div className="space-y-3">
            <label className="block text-sm font-medium text-gray-700">
              Select Account Type
            </label>
            <div className="grid grid-cols-1 gap-2 sm:grid-cols-3">
              <button
                type="button"
                onClick={() => handleTenantChange('system')}
                className={`p-3 text-sm font-medium rounded-lg border-2 transition-all ${
                  selectedTenant === 'system'
                    ? 'border-blue-600 bg-blue-600 text-white'
                    : 'border-gray-200 text-gray-700 hover:border-blue-300'
                }`}
              >
                🏗️ System Admin
              </button>
              <button
                type="button"
                onClick={() => handleTenantChange('demo-corp')}
                className={`p-3 text-sm font-medium rounded-lg border-2 transition-all ${
                  selectedTenant === 'demo-corp'
                    ? 'border-blue-600 bg-blue-600 text-white'
                    : 'border-gray-200 text-gray-700 hover:border-blue-300'
                }`}
              >
                🏢 Tenant Admin
              </button>
              <button
                type="button"
                onClick={() => handleTenantChange('customer')}
                className={`p-3 text-sm font-medium rounded-lg border-2 transition-all ${
                  selectedTenant === 'customer'
                    ? 'border-blue-600 bg-blue-600 text-white'
                    : 'border-gray-200 text-gray-700 hover:border-blue-300'
                }`}
              >
                👤 Customer
              </button>
            </div>
          </div>

          {/* Email Input */}
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-2">
              Email Address
            </label>
            <input
              type="email"
              id="email"
              name="email"
              value={credentials.email}
              onChange={handleInputChange}
              placeholder={getEmailPlaceholder()}
              required
              className={`w-full px-4 py-3 border rounded-lg text-gray-900 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors ${
                error ? 'border-red-500' : 'border-gray-300'
              }`}
            />
          </div>

          {/* Password Input */}
          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-2">
              Password
            </label>
            <input
              type="password"
              id="password"
              name="password"
              value={credentials.password}
              onChange={handleInputChange}
              placeholder="Enter your password"
              required
              className={`w-full px-4 py-3 border rounded-lg text-gray-900 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors ${
                error ? 'border-red-500' : 'border-gray-300'
              }`}
            />
          </div>

          {/* Error Message */}
          {error && (
            <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm">
              {error}
            </div>
          )}

          {/* Success Message */}
          {success && (
            <div className="bg-green-50 border border-green-200 text-green-700 px-4 py-3 rounded-lg text-sm">
              {success}
            </div>
          )}

          {/* Submit Button */}
          <button
            type="submit"
            disabled={isLoading}
            className="w-full bg-blue-600 text-white py-3 px-4 rounded-lg font-medium hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
          >
            {isLoading ? (
              <div className="flex items-center justify-center">
                <div className="animate-spin rounded-full h-5 w-5 border-b-2 border-white mr-2"></div>
                Logging in...
              </div>
            ) : (
              'Login'
            )}
          </button>
        </form>

        {/* Demo Credentials */}
        <div className="mt-8 bg-gray-50 rounded-lg p-4">
          <h4 className="text-sm font-medium text-gray-700 mb-3">Demo Credentials:</h4>
          <div className="space-y-2 text-xs text-gray-600">
            <p><strong>System Admin:</strong> admin@zplus.com / admin123</p>
            <p><strong>Tenant Admin:</strong> admin@demo-corp.zplus.com / demo123</p>
            <p><strong>Customer:</strong> john@demo-corp.zplus.com / user123</p>
          </div>
        </div>

        {/* Back to Home */}
        <div className="mt-6 text-center">
          <button
            onClick={() => router.push('/')}
            className="text-sm text-blue-600 hover:text-blue-700 transition-colors"
          >
            ← Back to Home
          </button>
        </div>
      </div>
    </div>
  )
}
