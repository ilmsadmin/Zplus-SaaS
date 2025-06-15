import { ReactNode } from 'react'

interface SystemLayoutProps {
  children: ReactNode
}

export function SystemLayout({ children }: SystemLayoutProps) {
  return (
    <div className="min-h-screen bg-gray-50">
      {children}
    </div>
  )
}

interface SystemHeaderProps {
  userMenu?: ReactNode
}

export function SystemHeader({ userMenu }: SystemHeaderProps) {
  return (
    <header className="bg-white border-b border-gray-200">
      <div className="px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <h1 className="text-xl font-bold text-blue-600">Zplus SaaS</h1>
            </div>
          </div>

          {/* Navigation */}
          <nav className="hidden md:flex items-center space-x-8">
            <a href="/admin" className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium">
              Dashboard
            </a>
            <a href="/admin/tenants" className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium">
              Tenants
            </a>
            <a href="/admin/plans" className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium">
              Plans
            </a>
            <a href="/admin/modules" className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium">
              Modules
            </a>
            <a href="/admin/billing" className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium">
              Billing
            </a>
            <a href="/admin/settings" className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium">
              Settings
            </a>
          </nav>

          {/* User Menu */}
          <div className="flex items-center">
            {userMenu || (
              <div className="flex items-center space-x-2 cursor-pointer hover:bg-gray-100 rounded-md p-2">
                <div className="w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  SA
                </div>
                <span className="text-sm font-medium text-gray-700">System Admin</span>
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

interface SystemMainProps {
  children: ReactNode
  title?: string
  subtitle?: string
}

export function SystemMain({ children, title, subtitle }: SystemMainProps) {
  return (
    <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      {(title || subtitle) && (
        <div className="px-4 py-6 sm:px-0">
          <div className="mb-8">
            {title && <h1 className="text-3xl font-bold text-gray-900">{title}</h1>}
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