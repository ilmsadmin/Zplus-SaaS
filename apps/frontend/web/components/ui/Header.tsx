import { ReactNode } from 'react'

interface HeaderProps {
  children: ReactNode
  className?: string
}

export function Header({ children, className = '' }: HeaderProps) {
  return (
    <header className={`bg-white border-b border-gray-200 ${className}`}>
      <div className="px-4 sm:px-6 lg:px-8 py-4">
        {children}
      </div>
    </header>
  )
}

interface HeaderContentProps {
  logo?: ReactNode
  navigation?: ReactNode
  userMenu?: ReactNode
}

export function HeaderContent({ logo, navigation, userMenu }: HeaderContentProps) {
  return (
    <div className="flex items-center justify-between">
      {logo && <div className="flex-shrink-0">{logo}</div>}
      {navigation && (
        <nav className="hidden md:flex items-center space-x-8">
          {navigation}
        </nav>  
      )}
      {userMenu && <div className="flex items-center">{userMenu}</div>}
    </div>
  )
}

interface LogoProps {
  title: string
  subtitle?: string
  className?: string
}

export function Logo({ title, subtitle, className = '' }: LogoProps) {
  return (
    <div className={`flex items-center space-x-3 ${className}`}>
      <div className="w-10 h-10 bg-blue-600 rounded-lg flex items-center justify-center text-white font-bold text-lg">
        {title.charAt(0)}
      </div>
      <div>
        <h1 className="text-xl font-bold text-gray-900">{title}</h1>
        {subtitle && <p className="text-sm text-gray-600">{subtitle}</p>}
      </div>
    </div>
  )
}

interface NavLinkProps {
  href: string
  children: ReactNode
  active?: boolean
  className?: string
}

export function NavLink({ href, children, active = false, className = '' }: NavLinkProps) {
  const baseClasses = "text-gray-600 hover:text-gray-900 hover:bg-gray-100 px-3 py-2 rounded-md text-sm font-medium transition-colors"
  const activeClasses = active ? "bg-gray-100 text-gray-900" : ""
  
  return (
    <a href={href} className={`${baseClasses} ${activeClasses} ${className}`}>
      {children}
    </a>
  )
}