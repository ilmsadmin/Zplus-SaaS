import { TenantLayout, TenantHeader, TenantMain } from '@/components/layouts/TenantLayout'
import { Card, Grid } from '@/components/ui'

interface TenantPageProps {
  params: {
    slug: string
  }
}

export default function TenantHomePage({ params }: TenantPageProps) {
  const { slug } = params
  const tenantName = slug.charAt(0).toUpperCase() + slug.slice(1) + ' Corporation'
  
  return (
    <TenantLayout>
      <TenantHeader 
        tenantName={tenantName}
        tenantSlug={slug}
        isAdmin={false}
      />
      <TenantMain
        title={`Welcome to ${tenantName}`}
        subtitle="Access your organization's services and applications"
      >
        <Grid cols={2} className="mb-8">
          <Card>
            <div className="p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="w-12 h-12 bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
                  <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                  </svg>
                </div>
                <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full">Active</span>
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">CRM</h3>
              <p className="text-gray-600 mb-6">Customer Relationship Management - Manage leads, customers, and sales pipeline</p>
              <a 
                href={`/tenant/${slug}/crm`}
                className="inline-flex items-center justify-center w-full px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-md hover:bg-blue-700 transition-colors"
              >
                Access CRM
              </a>
            </div>
          </Card>

          <Card>
            <div className="p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="w-12 h-12 bg-gradient-to-r from-green-500 to-green-600 rounded-lg flex items-center justify-center">
                  <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.746 0 3.332.477 4.5 1.253v13C19.832 18.477 18.246 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                  </svg>
                </div>
                <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full">Active</span>
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">LMS</h3>
              <p className="text-gray-600 mb-6">Learning Management System - Online courses, training, and certificates</p>
              <a 
                href={`/tenant/${slug}/lms`}
                className="inline-flex items-center justify-center w-full px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-md hover:bg-green-700 transition-colors"
              >
                Access LMS
              </a>
            </div>
          </Card>

          <Card>
            <div className="p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="w-12 h-12 bg-gradient-to-r from-purple-500 to-purple-600 rounded-lg flex items-center justify-center">
                  <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2-2v2m8 0V6a2 2 0 012 2v6a2 2 0 01-2 2H6a2 2 0 01-2-2V8a2 2 0 012-2V4" />
                  </svg>
                </div>
                <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full">Active</span>
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">HRM</h3>
              <p className="text-gray-600 mb-6">Human Resource Management - Employee management, payroll, and HR processes</p>
              <a 
                href={`/tenant/${slug}/hrm`}
                className="inline-flex items-center justify-center w-full px-4 py-2 bg-purple-600 text-white text-sm font-medium rounded-md hover:bg-purple-700 transition-colors"
              >
                Access HRM
              </a>
            </div>
          </Card>

          <Card>
            <div className="p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="w-12 h-12 bg-gradient-to-r from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
                  <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full">Active</span>
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">POS</h3>
              <p className="text-gray-600 mb-6">Point of Sale System - Retail management, inventory, and sales tracking</p>
              <a 
                href={`/tenant/${slug}/pos`}
                className="inline-flex items-center justify-center w-full px-4 py-2 bg-orange-600 text-white text-sm font-medium rounded-md hover:bg-orange-700 transition-colors"
              >
                Access POS
              </a>
            </div>
          </Card>
        </Grid>

        {/* Quick Info */}
        <Card>
          <div className="p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Organization Information</h3>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <h4 className="text-sm font-medium text-gray-500 uppercase tracking-wide">Organization</h4>
                <p className="text-lg font-medium text-gray-900">{tenantName}</p>
              </div>
              <div>
                <h4 className="text-sm font-medium text-gray-500 uppercase tracking-wide">Plan</h4>
                <p className="text-lg font-medium text-gray-900">Professional</p>
              </div>
              <div>
                <h4 className="text-sm font-medium text-gray-500 uppercase tracking-wide">Active Modules</h4>
                <p className="text-lg font-medium text-gray-900">4 Modules</p>
              </div>
            </div>
          </div>
        </Card>
      </TenantMain>
    </TenantLayout>
  )
}