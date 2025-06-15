import { ReactNode } from 'react'

interface CardProps {
  children: ReactNode
  className?: string
  padding?: 'none' | 'sm' | 'md' | 'lg'
}

export function Card({ children, className = '', padding = 'md' }: CardProps) {
  const paddingClasses = {
    none: '',
    sm: 'p-4',
    md: 'p-6',
    lg: 'p-8'
  }
  
  return (
    <div className={`bg-white rounded-lg shadow border border-gray-200 ${paddingClasses[padding]} ${className}`}>
      {children}
    </div>
  )
}

interface StatsCardProps {
  title: string
  value: string | number
  change?: string
  color?: 'blue' | 'green' | 'yellow' | 'purple' | 'red' | 'orange'
  icon?: ReactNode
}

export function StatsCard({ title, value, change, color = 'blue', icon }: StatsCardProps) {
  const colorClasses = {
    blue: 'bg-blue-50 text-blue-800 border-blue-200',
    green: 'bg-green-50 text-green-800 border-green-200', 
    yellow: 'bg-yellow-50 text-yellow-800 border-yellow-200',
    purple: 'bg-purple-50 text-purple-800 border-purple-200',
    red: 'bg-red-50 text-red-800 border-red-200',
    orange: 'bg-orange-50 text-orange-800 border-orange-200'
  }
  
  return (
    <div className={`rounded-lg border p-6 ${colorClasses[color]}`}>
      <div className="flex items-center justify-between">
        <div>
          <p className="text-sm font-medium opacity-75">{title}</p>
          <p className="text-3xl font-bold">{value}</p>
          {change && (
            <p className="text-sm mt-1 opacity-75">{change}</p>
          )}
        </div>
        {icon && (
          <div className="opacity-50">
            {icon}
          </div>
        )}
      </div>
    </div>
  )
}

interface GridProps {
  children: ReactNode
  cols?: 1 | 2 | 3 | 4 | 6
  gap?: 'sm' | 'md' | 'lg'
  className?: string
}

export function Grid({ children, cols = 1, gap = 'md', className = '' }: GridProps) {
  const colClasses = {
    1: 'grid-cols-1',
    2: 'grid-cols-1 md:grid-cols-2',
    3: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3',
    4: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-4',
    6: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-6'
  }
  
  const gapClasses = {
    sm: 'gap-4',
    md: 'gap-6',
    lg: 'gap-8'
  }
  
  return (
    <div className={`grid ${colClasses[cols]} ${gapClasses[gap]} ${className}`}>
      {children}
    </div>
  )
}