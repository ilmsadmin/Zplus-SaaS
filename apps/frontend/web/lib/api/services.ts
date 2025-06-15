import { apiClient, withTenant } from './client'
import { 
  SystemStats, 
  Tenant, 
  Plan, 
  SystemModule, 
  Activity,
  TenantStats,
  User,
  Role,
  CRMStats,
  Lead,
  Customer,
  Opportunity,
  LMSStats,
  Course,
  Enrollment,
  Certificate
} from '@/lib/types'

// System Admin APIs
export const systemApi = {
  getStats: () => apiClient.get<SystemStats>('/system/stats'),
  getTenants: (params?: { page?: number; limit?: number }) => 
    apiClient.getPaginated<Tenant>('/system/tenants', params),
  getTenant: (id: string) => apiClient.get<Tenant>(`/system/tenants/${id}`),
  createTenant: (data: Partial<Tenant>) => apiClient.post<Tenant>('/system/tenants', data),
  updateTenant: (id: string, data: Partial<Tenant>) => 
    apiClient.put<Tenant>(`/system/tenants/${id}`, data),
  deleteTenant: (id: string) => apiClient.delete(`/system/tenants/${id}`),
  
  getPlans: () => apiClient.get<Plan[]>('/system/plans'),
  getModules: () => apiClient.get<SystemModule[]>('/system/modules'),
  getActivity: (params?: { page?: number; limit?: number }) => 
    apiClient.getPaginated<Activity>('/system/activity', params),
}

// Tenant Admin APIs
export const tenantAdminApi = (tenantSlug: string) => {
  const api = withTenant(tenantSlug)
  
  return {
    getStats: () => api.get<TenantStats>('/tenant/stats'),
    
    // User management
    getUsers: (params?: { page?: number; limit?: number }) => 
      api.getPaginated<User>('/tenant/users', params),
    getUser: (id: string) => api.get<User>(`/tenant/users/${id}`),
    createUser: (data: Partial<User>) => api.post<User>('/tenant/users', data),
    updateUser: (id: string, data: Partial<User>) => 
      api.put<User>(`/tenant/users/${id}`, data),
    deleteUser: (id: string) => api.delete(`/tenant/users/${id}`),
    
    // Role management
    getRoles: () => api.get<Role[]>('/tenant/roles'),
    createRole: (data: Partial<Role>) => api.post<Role>('/tenant/roles', data),
    updateRole: (id: string, data: Partial<Role>) => 
      api.put<Role>(`/tenant/roles/${id}`, data),
    deleteRole: (id: string) => api.delete(`/tenant/roles/${id}`),
    
    getActivity: (params?: { page?: number; limit?: number }) => 
      api.getPaginated<Activity>('/tenant/activity', params),
  }
}

// CRM APIs
export const crmApi = (tenantSlug: string) => {
  const api = withTenant(tenantSlug)
  
  return {
    getStats: () => api.get<CRMStats>('/crm/stats'),
    
    // Lead management
    getLeads: (params?: { page?: number; limit?: number; status?: string }) => 
      api.getPaginated<Lead>('/crm/leads', params),
    getLead: (id: string) => api.get<Lead>(`/crm/leads/${id}`),
    createLead: (data: Partial<Lead>) => api.post<Lead>('/crm/leads', data),
    updateLead: (id: string, data: Partial<Lead>) => 
      api.put<Lead>(`/crm/leads/${id}`, data),
    deleteLead: (id: string) => api.delete(`/crm/leads/${id}`),
    
    // Customer management
    getCustomers: (params?: { page?: number; limit?: number }) => 
      api.getPaginated<Customer>('/crm/customers', params),
    getCustomer: (id: string) => api.get<Customer>(`/crm/customers/${id}`),
    createCustomer: (data: Partial<Customer>) => api.post<Customer>('/crm/customers', data),
    updateCustomer: (id: string, data: Partial<Customer>) => 
      api.put<Customer>(`/crm/customers/${id}`, data),
    deleteCustomer: (id: string) => api.delete(`/crm/customers/${id}`),
    
    // Opportunity management
    getOpportunities: (params?: { page?: number; limit?: number }) => 
      api.getPaginated<Opportunity>('/crm/opportunities', params),
    getOpportunity: (id: string) => api.get<Opportunity>(`/crm/opportunities/${id}`),
    createOpportunity: (data: Partial<Opportunity>) => 
      api.post<Opportunity>('/crm/opportunities', data),
    updateOpportunity: (id: string, data: Partial<Opportunity>) => 
      api.put<Opportunity>(`/crm/opportunities/${id}`, data),
    deleteOpportunity: (id: string) => api.delete(`/crm/opportunities/${id}`),
  }
}

// LMS APIs
export const lmsApi = (tenantSlug: string) => {
  const api = withTenant(tenantSlug)
  
  return {
    getStats: () => api.get<LMSStats>('/lms/stats'),
    
    // Course management
    getCourses: (params?: { page?: number; limit?: number; category?: string }) => 
      api.getPaginated<Course>('/lms/courses', params),
    getCourse: (id: string) => api.get<Course>(`/lms/courses/${id}`),
    
    // Enrollment management
    getEnrollments: (params?: { page?: number; limit?: number }) => 
      api.getPaginated<Enrollment>('/lms/enrollments', params),
    enrollCourse: (courseId: string) => 
      api.post<Enrollment>('/lms/enrollments', { courseId }),
    updateProgress: (enrollmentId: string, progress: number) => 
      api.put<Enrollment>(`/lms/enrollments/${enrollmentId}/progress`, { progress }),
    
    // Certificate management
    getCertificates: () => api.get<Certificate[]>('/lms/certificates'),
    getCertificate: (id: string) => api.get<Certificate>(`/lms/certificates/${id}`),
    
    getActivity: (params?: { page?: number; limit?: number }) => 
      api.getPaginated<Activity>('/lms/activity', params),
  }
}

// Common utility functions
export const mockApiDelay = (ms: number = 1000) => 
  new Promise(resolve => setTimeout(resolve, ms))

// Mock data generators (for development/testing)
export const generateMockStats = {
  system: (): SystemStats => ({
    totalTenants: Math.floor(Math.random() * 50) + 10,
    activePlans: 3,
    systemModules: 6,
    systemHealth: 100,
  }),
  
  tenant: (): TenantStats => ({
    activeUsers: Math.floor(Math.random() * 100) + 20,
    activeModules: Math.floor(Math.random() * 4) + 2,
    userRoles: Math.floor(Math.random() * 3) + 3,
    integrations: Math.floor(Math.random() * 5) + 1,
  }),
  
  crm: (): CRMStats => ({
    totalLeads: Math.floor(Math.random() * 200) + 50,
    activeCustomers: Math.floor(Math.random() * 100) + 30,
    opportunities: Math.floor(Math.random() * 50) + 10,
    closedDeals: Math.floor(Math.random() * 100000) + 25000,
  }),
  
  lms: (): LMSStats => ({
    enrolledCourses: Math.floor(Math.random() * 10) + 5,
    completedCourses: Math.floor(Math.random() * 15) + 8,
    learningHours: Math.floor(Math.random() * 200) + 100,
    certificates: Math.floor(Math.random() * 5) + 2,
  }),
}