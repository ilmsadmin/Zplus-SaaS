import { TenantLayout, TenantHeader, TenantMain } from '@/components/layouts/TenantLayout'
import { Card, StatsCard, Grid } from '@/components/ui'

interface CRMPageProps {
  params: {
    slug: string
  }
}

export default function CRMDashboard({ params }: CRMPageProps) {
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
        title="CRM Dashboard"
        subtitle="Customer Relationship Management - Manage leads, customers, and sales pipeline"
      >
        {/* CRM Stats */}
        <Grid cols={4} className="mb-8">
          <StatsCard
            title="Total Leads"
            value="142"
            change="+12 this week"
            color="blue"
          />
          <StatsCard
            title="Active Customers"
            value="89"
            change="+5 this month"
            color="green"
          />
          <StatsCard
            title="Opportunities"
            value="23"
            change="$125K pipeline"
            color="purple"
          />
          <StatsCard
            title="Closed Deals"
            value="$45K"
            change="+18% vs last month"
            color="orange"
          />
        </Grid>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
          {/* Recent Leads */}
          <Card>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-lg font-semibold text-gray-900">Recent Leads</h2>
              <a href={`/tenant/${slug}/crm/leads`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                View all
              </a>
            </div>
            <div className="space-y-4">
              <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
                    <span className="text-sm font-medium text-blue-600">JD</span>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">John Doe</p>
                    <p className="text-sm text-gray-500">john@example.com</p>
                  </div>
                </div>
                <div className="text-right">
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-yellow-100 text-yellow-800">
                    New
                  </span>
                  <p className="text-sm text-gray-500 mt-1">2 hours ago</p>
                </div>
              </div>
              <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-green-100 rounded-full flex items-center justify-center">
                    <span className="text-sm font-medium text-green-600">SM</span>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Sarah Miller</p>
                    <p className="text-sm text-gray-500">sarah@company.com</p>
                  </div>
                </div>
                <div className="text-right">
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    Qualified
                  </span>
                  <p className="text-sm text-gray-500 mt-1">5 hours ago</p>
                </div>
              </div>
              <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-purple-100 rounded-full flex items-center justify-center">
                    <span className="text-sm font-medium text-purple-600">MW</span>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Mike Wilson</p>
                    <p className="text-sm text-gray-500">mike@techcorp.com</p>
                  </div>
                </div>
                <div className="text-right">
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                    Proposal
                  </span>
                  <p className="text-sm text-gray-500 mt-1">1 day ago</p>
                </div>
              </div>
            </div>
          </Card>

          {/* Sales Pipeline */}
          <Card>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-lg font-semibold text-gray-900">Sales Pipeline</h2>
              <a href={`/tenant/${slug}/crm/pipeline`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                View pipeline
              </a>
            </div>
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-3 h-3 bg-yellow-400 rounded-full"></div>
                  <span className="text-sm font-medium text-gray-900">Prospecting</span>
                </div>
                <div className="text-right">
                  <span className="text-sm font-medium text-gray-900">34 leads</span>
                  <p className="text-sm text-gray-500">$67K potential</p>
                </div>
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-3 h-3 bg-blue-400 rounded-full"></div>
                  <span className="text-sm font-medium text-gray-900">Qualified</span>
                </div>
                <div className="text-right">
                  <span className="text-sm font-medium text-gray-900">18 leads</span>
                  <p className="text-sm text-gray-500">$89K potential</p>
                </div>
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-3 h-3 bg-purple-400 rounded-full"></div>
                  <span className="text-sm font-medium text-gray-900">Proposal</span>
                </div>
                <div className="text-right">
                  <span className="text-sm font-medium text-gray-900">12 leads</span>
                  <p className="text-sm text-gray-500">$156K potential</p>
                </div>
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-3 h-3 bg-green-400 rounded-full"></div>
                  <span className="text-sm font-medium text-gray-900">Negotiation</span>
                </div>
                <div className="text-right">
                  <span className="text-sm font-medium text-gray-900">7 leads</span>
                  <p className="text-sm text-gray-500">$234K potential</p>
                </div>
              </div>
            </div>
          </Card>
        </div>

        {/* Quick Actions */}
        <Grid cols={4}>
          <Card>
            <div className="text-center p-4">
              <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900 mb-2">Add Lead</h3>
              <a 
                href={`/tenant/${slug}/crm/leads/new`}
                className="text-blue-600 hover:text-blue-700 text-sm font-medium"
              >
                Create new lead
              </a>
            </div>
          </Card>
          
          <Card>
            <div className="text-center p-4">
              <div className="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900 mb-2">Customers</h3>
              <a 
                href={`/tenant/${slug}/crm/customers`}
                className="text-green-600 hover:text-green-700 text-sm font-medium"
              >
                Manage customers
              </a>
            </div>
          </Card>
          
          <Card>
            <div className="text-center p-4">
              <div className="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900 mb-2">Reports</h3>
              <a 
                href={`/tenant/${slug}/crm/reports`}
                className="text-purple-600 hover:text-purple-700 text-sm font-medium"
              >
                View reports
              </a>
            </div>
          </Card>
          
          <Card>
            <div className="text-center p-4">
              <div className="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900 mb-2">Settings</h3>
              <a 
                href={`/tenant/${slug}/crm/settings`}
                className="text-orange-600 hover:text-orange-700 text-sm font-medium"
              >
                CRM settings
              </a>
            </div>
          </Card>
        </Grid>
      </TenantMain>
    </TenantLayout>
  )
}