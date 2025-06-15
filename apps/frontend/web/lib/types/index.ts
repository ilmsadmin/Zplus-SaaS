// System Admin Types
export interface SystemStats {
  totalTenants: number
  activePlans: number
  systemModules: number
  systemHealth: number
}

export interface Tenant {
  id: string
  slug: string
  name: string
  domain?: string
  plan: string
  status: 'active' | 'inactive' | 'suspended'
  createdAt: string
  updatedAt: string
}

export interface Plan {
  id: string
  name: string
  price: number
  features: string[]
  maxUsers: number
  maxModules: number
}

export interface SystemModule {
  id: string
  name: string
  slug: string
  description: string
  version: string
  enabled: boolean
}

// Tenant Admin Types
export interface TenantStats {
  activeUsers: number
  activeModules: number
  userRoles: number
  integrations: number
}

export interface User {
  id: string
  email: string
  firstName: string
  lastName: string
  role: string
  status: 'active' | 'inactive' | 'pending'
  lastLogin?: string
  createdAt: string
}

export interface Role {
  id: string
  name: string
  permissions: string[]
  userCount: number
}

// CRM Types
export interface CRMStats {
  totalLeads: number
  activeCustomers: number
  opportunities: number
  closedDeals: number
}

export interface Lead {
  id: string
  firstName: string
  lastName: string
  email: string
  phone?: string
  company?: string
  status: 'new' | 'qualified' | 'proposal' | 'negotiation' | 'closed-won' | 'closed-lost'
  source: string
  value?: number
  createdAt: string
  updatedAt: string
}

export interface Customer {
  id: string
  name: string
  email: string
  phone?: string
  company?: string
  address?: string
  status: 'active' | 'inactive'
  totalValue: number
  createdAt: string
}

export interface Opportunity {
  id: string
  title: string
  customerId: string
  value: number
  stage: string
  probability: number
  closeDate: string
  createdAt: string
}

// LMS Types
export interface LMSStats {
  enrolledCourses: number
  completedCourses: number
  learningHours: number
  certificates: number
}

export interface Course {
  id: string
  title: string
  description: string
  instructor: string
  duration: number // in minutes
  lessonsCount: number
  difficulty: 'beginner' | 'intermediate' | 'advanced'
  category: string
  thumbnail?: string
  price?: number
  enrolled: boolean
  progress?: number
  completedAt?: string
  createdAt: string
}

export interface Enrollment {
  id: string
  courseId: string
  userId: string
  progress: number
  completedLessons: number
  totalLessons: number
  startedAt: string
  completedAt?: string
}

export interface Certificate {
  id: string
  courseId: string
  userId: string
  issueDate: string
  expiryDate?: string
  certificateUrl: string
}

// HRM Types (placeholder for future implementation)
export interface HRMStats {
  totalEmployees: number
  activePositions: number
  leaveRequests: number
  payrollPending: number
}

// POS Types (placeholder for future implementation)
export interface POSStats {
  dailySales: number
  totalProducts: number
  lowStockItems: number
  activeSessions: number
}

// Common Types
export interface Activity {
  id: string
  type: string
  title: string
  description: string
  userId?: string
  userName?: string
  createdAt: string
}

export interface ApiResponse<T> {
  data: T
  success: boolean
  message?: string
  error?: string
}

export interface PaginatedResponse<T> {
  data: T[]
  pagination: {
    page: number
    limit: number
    total: number
    totalPages: number
  }
}