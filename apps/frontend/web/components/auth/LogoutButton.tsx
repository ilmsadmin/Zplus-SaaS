'use client'

import { useAuth } from '@/hooks/useAuth'

interface LogoutButtonProps {
  className?: string
  children?: React.ReactNode
}

export function LogoutButton({ className = '', children }: LogoutButtonProps) {
  const { logout } = useAuth()

  const handleLogout = async () => {
    try {
      await logout()
    } catch (error) {
      console.error('Logout failed:', error)
    }
  }

  return (
    <button
      onClick={handleLogout}
      className={`text-red-600 hover:text-red-700 ${className}`}
    >
      {children || 'Logout'}
    </button>
  )
}

export default LogoutButton
