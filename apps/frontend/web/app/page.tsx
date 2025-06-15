'use client'

import { useAuth } from '@/hooks/useAuth'
import UserMenu from '@/components/auth/UserMenu'
import Link from 'next/link'

export default function Home() {
  const { isAuthenticated, isLoading } = useAuth()

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      {/* Header */}
      <header className="bg-white shadow-lg">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div className="flex items-center space-x-4">
              <div className="bg-blue-600 text-white p-3 rounded-lg">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
              </div>
              <h1 className="text-2xl font-bold text-gray-900">Zplus SaaS</h1>
            </div>
            <div className="flex items-center space-x-4">
              <span className="text-sm text-gray-600">Multi-Tenant Platform</span>
              <div className="bg-green-100 text-green-800 px-3 py-1 rounded-full text-sm font-medium">
                Online
              </div>
              {!isLoading && (
                isAuthenticated ? (
                  <UserMenu />
                ) : (
                  <Link
                    href="/login"
                    className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors text-sm font-medium"
                  >
                    Login
                  </Link>
                )
              )}
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        {/* Hero Section */}
        <div className="text-center mb-16">
          <h2 className="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
            Welcome to
            <span className="text-blue-600 ml-4">Zplus SaaS</span>
          </h2>
          <p className="text-xl text-gray-600 max-w-2xl mx-auto mb-8">
            Modern multi-tenant platform with modular architecture supporting CRM, LMS, POS, and HRM modules
          </p>
          
          {/* Status Indicators */}
          <div className="flex justify-center space-x-4 mb-8">
            <div className="bg-white rounded-lg p-4 shadow-md">
              <div className="flex items-center space-x-2">
                <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                <span className="text-sm font-medium text-gray-700">Backend Services</span>
              </div>
            </div>
            <div className="bg-white rounded-lg p-4 shadow-md">
              <div className="flex items-center space-x-2">
                <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                <span className="text-sm font-medium text-gray-700">Database</span>
              </div>
            </div>
            <div className="bg-white rounded-lg p-4 shadow-md">
              <div className="flex items-center space-x-2">
                <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                <span className="text-sm font-medium text-gray-700">Frontend</span>
              </div>
            </div>
          </div>
        </div>

        {/* Navigation Cards */}
        <div className="grid md:grid-cols-3 gap-8 mb-16">
          <a
            href="/admin"
            className="group bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-8 border-2 border-transparent hover:border-blue-200"
          >
            <div className="flex items-center space-x-4 mb-4">
              <div className="bg-blue-100 text-blue-600 p-3 rounded-lg group-hover:bg-blue-600 group-hover:text-white transition-colors">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <div>
                <h3 className="text-xl font-semibold text-gray-900 group-hover:text-blue-600 transition-colors">
                  System Admin
                </h3>
                <div className="flex items-center text-blue-600 group-hover:translate-x-2 transition-transform">
                  <span className="text-sm">Access →</span>
                </div>
              </div>
            </div>
            <p className="text-gray-600">
              Global tenant and system management interface. Manage all tenants, users, and system configurations.
            </p>
          </a>

          <a
            href="/tenant/demo"
            className="group bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-8 border-2 border-transparent hover:border-purple-200"
          >
            <div className="flex items-center space-x-4 mb-4">
              <div className="bg-purple-100 text-purple-600 p-3 rounded-lg group-hover:bg-purple-600 group-hover:text-white transition-colors">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                </svg>
              </div>
              <div>
                <h3 className="text-xl font-semibold text-gray-900 group-hover:text-purple-600 transition-colors">
                  Demo Tenant
                </h3>
                <div className="flex items-center text-purple-600 group-hover:translate-x-2 transition-transform">
                  <span className="text-sm">Access →</span>
                </div>
              </div>
            </div>
            <p className="text-gray-600">
              Demo tenant customer portal with sample data. Explore CRM, LMS, POS, and HRM modules.
            </p>
          </a>

          <a
            href="/tenant/demo/admin"
            className="group bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 p-8 border-2 border-transparent hover:border-green-200"
          >
            <div className="flex items-center space-x-4 mb-4">
              <div className="bg-green-100 text-green-600 p-3 rounded-lg group-hover:bg-green-600 group-hover:text-white transition-colors">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                </svg>
              </div>
              <div>
                <h3 className="text-xl font-semibold text-gray-900 group-hover:text-green-600 transition-colors">
                  Tenant Admin
                </h3>
                <div className="flex items-center text-green-600 group-hover:translate-x-2 transition-transform">
                  <span className="text-sm">Access →</span>
                </div>
              </div>
            </div>
            <p className="text-gray-600">
              Demo tenant administration interface. Manage tenant-specific settings, users, and modules.
            </p>
          </a>
        </div>

        {/* Features Section */}
        <div className="bg-white rounded-2xl shadow-lg p-8">
          <h3 className="text-2xl font-bold text-gray-900 mb-6 text-center">Platform Features</h3>
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div className="text-center">
              <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg className="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
              <h4 className="font-semibold text-gray-900 mb-2">CRM</h4>
              <p className="text-sm text-gray-600">Customer Relationship Management</p>
            </div>
            
            <div className="text-center">
              <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg className="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                </svg>
              </div>
              <h4 className="font-semibold text-gray-900 mb-2">LMS</h4>
              <p className="text-sm text-gray-600">Learning Management System</p>
            </div>
            
            <div className="text-center">
              <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg className="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 3h2l.4 2M7 13h10l4-8H5.4m-.4-2H1m6 16a2 2 0 100-4 2 2 0 000 4zm10 0a2 2 0 100-4 2 2 0 000 4z" />
                </svg>
              </div>
              <h4 className="font-semibold text-gray-900 mb-2">POS</h4>
              <p className="text-sm text-gray-600">Point of Sale System</p>
            </div>
            
            <div className="text-center">
              <div className="bg-yellow-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg className="w-8 h-8 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2-2v2m8 0V6a2 2 0 012 2v6a2 2 0 01-2 2H8a2 2 0 01-2-2V8a2 2 0 012-2V6m8 0H8" />
                </svg>
              </div>
              <h4 className="font-semibold text-gray-900 mb-2">HRM</h4>
              <p className="text-sm text-gray-600">Human Resource Management</p>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="bg-gray-800 text-white py-8 mt-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <p className="text-gray-400">© 2024 Zplus SaaS. Modern multi-tenant platform.</p>
            <div className="mt-4 flex justify-center space-x-6">
              <span className="text-sm text-gray-400">Version 1.0.0</span>
              <span className="text-sm text-gray-400">•</span>
              <span className="text-sm text-gray-400">All Services Online</span>
            </div>
          </div>
        </div>
      </footer>
    </div>
  )
}
