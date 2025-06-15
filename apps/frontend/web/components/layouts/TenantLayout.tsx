import { ReactNode } from 'react'

interface TenantLayoutProps {
  children: ReactNode
}

export function TenantLayout({ children }: TenantLayoutProps) {
  return (
    <div className="min-h-screen bg-gray-50">
      {children}
    </div>
  )
}

interface TenantHeaderProps {
  tenantName: string
  tenantSlug: string
  isAdmin?: boolean
  userMenu?: ReactNode
}

export function TenantHeader({ tenantName, tenantSlug, isAdmin = false, userMenu }: TenantHeaderProps) {
  return (
    <header className="bg-white border-b border-gray-200">
      <div className="px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          {/* Tenant Branding */}
          <div className="flex items-center space-x-4">
            <div className="w-8 h-8 sm:w-10 sm:h-10 bg-blue-600 rounded-lg flex items-center justify-center text-white font-bold text-sm sm:text-lg">
              {tenantName.charAt(0)}
            </div>
            <div>
              <h1 className="text-sm sm:text-lg font-bold text-gray-900">{tenantName}</h1>
              <p className="text-xs sm:text-sm text-gray-600">{isAdmin ? 'Administration' : 'Customer Portal'}</p>
            </div>
          </div>

          {/* Mobile menu button */}
          <div className="md:hidden">
            <button
              type="button"
              className="p-2 rounded-md text-gray-600 hover:text-gray-900 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500"
              aria-expanded="false"
              aria-label="Toggle navigation menu"
            >
              <svg className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </button>
          </div>

          {/* Navigation */}
          <nav className="hidden md:flex items-center space-x-4 lg:space-x-8" role="navigation">
            {isAdmin ? (
              <>
                <a 
                  href={`/tenant/${tenantSlug}/admin`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  Dashboard
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/admin/users`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  Users & Roles
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/admin/customers`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  Customers
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/admin/modules`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  Modules
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/admin/integrations`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  Integrations
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/admin/settings`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  Settings
                </a>
              </>
            ) : (
              <>
                <a 
                  href={`/tenant/${tenantSlug}`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  Home
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/crm`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  CRM
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/lms`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  LMS
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/hrm`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  HRM
                </a>
                <a 
                  href={`/tenant/${tenantSlug}/pos`} 
                  className="text-gray-600 hover:text-gray-900 px-2 lg:px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  POS
                </a>
              </>
            )}
          </nav>

          {/* User Menu */}
          <div className="flex items-center">
            {userMenu || (
              <div className="flex items-center space-x-2 cursor-pointer hover:bg-gray-100 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500" tabIndex={0}>
                <div className="w-8 h-8 bg-gray-600 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  {isAdmin ? 'TA' : 'U'}
                </div>
                <span className="text-sm font-medium text-gray-700 hidden sm:block">
                  {isAdmin ? 'Tenant Admin' : 'User'}
                </span>
                <svg className="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                </svg>
              </div>
            )}
          </div>
        </div>
      </div>
    </header>
  )
}

interface TenantMainProps {
  children: ReactNode
  title?: string
  subtitle?: string
}

export function TenantMain({ children, title, subtitle }: TenantMainProps) {
  return (
    <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8" role="main">
      {(title || subtitle) && (
        <div className="px-4 py-6 sm:px-0">
          <div className="mb-8">
            {title && <h1 className="text-2xl sm:text-3xl font-bold text-gray-900">{title}</h1>}
            {subtitle && <p className="mt-2 text-gray-600">{subtitle}</p>}
          </div>
        </div>
      )}
      <div className="px-4 py-6 sm:px-0">
        {children}
      </div>
    </main>
  )
}