import { TenantLayout, TenantHeader, TenantMain } from '@/components/layouts/TenantLayout'
import { Card, StatsCard, Grid } from '@/components/ui'

interface POSPageProps {
  params: {
    slug: string
  }
}

export default function POSDashboard({ params }: POSPageProps) {
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
        title="Point of Sale System"
        subtitle="Retail management, inventory tracking, sales analytics, and customer transactions"
      >
        {/* POS Stats */}
        <Grid cols={4} className="mb-8">
          <StatsCard
            title="Today's Sales"
            value="$12,450"
            change="+15% vs yesterday"
            color="green"
          />
          <StatsCard
            title="Transactions"
            value="186"
            change="43 this hour"
            color="blue"
          />
          <StatsCard
            title="Inventory Items"
            value="1,247"
            change="28 low stock"
            color="yellow"
          />
          <StatsCard
            title="Active Registers"
            value="5"
            change="All operational"
            color="purple"
          />
        </Grid>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
          {/* Recent Transactions */}
          <Card>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-lg font-semibold text-gray-900">Recent Transactions</h2>
              <a href={`/tenant/${slug}/pos/transactions`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                View all
              </a>
            </div>
            <div className="space-y-4">
              <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-green-100 rounded-full flex items-center justify-center">
                    <svg className="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v2a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
                    </svg>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Transaction #TXN001</p>
                    <p className="text-sm text-gray-500">Register 1 • Cash payment</p>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-medium text-gray-900">$67.45</p>
                  <p className="text-sm text-gray-500">2 min ago</p>
                </div>
              </div>
              
              <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
                    <svg className="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                    </svg>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Transaction #TXN002</p>
                    <p className="text-sm text-gray-500">Register 2 • Card payment</p>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-medium text-gray-900">$124.99</p>
                  <p className="text-sm text-gray-500">5 min ago</p>
                </div>
              </div>
              
              <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-purple-100 rounded-full flex items-center justify-center">
                    <svg className="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />
                    </svg>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Transaction #TXN003</p>
                    <p className="text-sm text-gray-500">Register 3 • Mobile payment</p>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-medium text-gray-900">$89.50</p>
                  <p className="text-sm text-gray-500">8 min ago</p>
                </div>
              </div>
            </div>
          </Card>

          {/* Top Products */}
          <Card>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-lg font-semibold text-gray-900">Top Selling Products</h2>
              <a href={`/tenant/${slug}/pos/products`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                View all products
              </a>
            </div>
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-gradient-to-r from-blue-400 to-blue-600 rounded-lg flex items-center justify-center text-white text-sm font-medium">
                    1
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Premium Coffee Blend</p>
                    <p className="text-sm text-gray-500">SKU: COF001</p>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-medium text-gray-900">47 sold</p>
                  <p className="text-sm text-gray-500">$14.99 each</p>
                </div>
              </div>
              
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-gradient-to-r from-green-400 to-green-600 rounded-lg flex items-center justify-center text-white text-sm font-medium">
                    2
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Wireless Headphones</p>
                    <p className="text-sm text-gray-500">SKU: ELE001</p>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-medium text-gray-900">23 sold</p>
                  <p className="text-sm text-gray-500">$79.99 each</p>
                </div>
              </div>
              
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-gradient-to-r from-purple-400 to-purple-600 rounded-lg flex items-center justify-center text-white text-sm font-medium">
                    3
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Organic T-Shirt</p>
                    <p className="text-sm text-gray-500">SKU: CLO001</p>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-medium text-gray-900">31 sold</p>
                  <p className="text-sm text-gray-500">$24.99 each</p>
                </div>
              </div>
              
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-gradient-to-r from-orange-400 to-orange-600 rounded-lg flex items-center justify-center text-white text-sm font-medium">
                    4
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-900">Smartphone Case</p>
                    <p className="text-sm text-gray-500">SKU: ACC001</p>
                  </div>
                </div>
                <div className="text-right">
                  <p className="text-sm font-medium text-gray-900">19 sold</p>
                  <p className="text-sm text-gray-500">$19.99 each</p>
                </div>
              </div>
            </div>
          </Card>
        </div>

        {/* Quick Actions */}
        <Card className="mb-8">
          <h2 className="text-lg font-semibold text-gray-900 mb-6">Quick Actions</h2>
          <Grid cols={2} className="sm:grid-cols-4">
            <button className="text-center p-4 border border-gray-200 rounded-lg hover:border-blue-300 hover:bg-blue-50 transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500">
              <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v2a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900">New Sale</h3>
            </button>
            
            <button className="text-center p-4 border border-gray-200 rounded-lg hover:border-green-300 hover:bg-green-50 transition-colors focus:outline-none focus:ring-2 focus:ring-green-500">
              <div className="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900">Inventory</h3>
            </button>
            
            <button className="text-center p-4 border border-gray-200 rounded-lg hover:border-purple-300 hover:bg-purple-50 transition-colors focus:outline-none focus:ring-2 focus:ring-purple-500">
              <div className="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900">Reports</h3>
            </button>
            
            <button className="text-center p-4 border border-gray-200 rounded-lg hover:border-orange-300 hover:bg-orange-50 transition-colors focus:outline-none focus:ring-2 focus:ring-orange-500">
              <div className="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center mx-auto mb-3">
                <svg className="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <h3 className="text-sm font-semibold text-gray-900">Settings</h3>
            </button>
          </Grid>
        </Card>

        {/* Low Stock Alert */}
        <Card>
          <div className="flex items-center justify-between mb-6">
            <h2 className="text-lg font-semibold text-gray-900">Low Stock Alerts</h2>
            <span className="bg-red-100 text-red-800 text-xs font-medium px-2.5 py-0.5 rounded-full">
              28 items
            </span>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div className="border border-red-200 bg-red-50 rounded-lg p-4">
              <div className="flex items-center justify-between mb-2">
                <h3 className="font-medium text-gray-900">Premium Coffee Blend</h3>
                <span className="text-red-600 text-sm font-medium">5 left</span>
              </div>
              <p className="text-sm text-gray-600 mb-3">SKU: COF001 • Min. stock: 20</p>
              <button className="w-full bg-red-600 text-white text-sm font-medium py-2 rounded-md hover:bg-red-700">
                Reorder Now
              </button>
            </div>
            
            <div className="border border-orange-200 bg-orange-50 rounded-lg p-4">
              <div className="flex items-center justify-between mb-2">
                <h3 className="font-medium text-gray-900">Wireless Mouse</h3>
                <span className="text-orange-600 text-sm font-medium">12 left</span>
              </div>
              <p className="text-sm text-gray-600 mb-3">SKU: ELE002 • Min. stock: 15</p>
              <button className="w-full bg-orange-600 text-white text-sm font-medium py-2 rounded-md hover:bg-orange-700">
                Reorder Now
              </button>
            </div>
            
            <div className="border border-yellow-200 bg-yellow-50 rounded-lg p-4">
              <div className="flex items-center justify-between mb-2">
                <h3 className="font-medium text-gray-900">Notebook Set</h3>
                <span className="text-yellow-600 text-sm font-medium">18 left</span>
              </div>
              <p className="text-sm text-gray-600 mb-3">SKU: STA001 • Min. stock: 25</p>
              <button className="w-full bg-yellow-600 text-white text-sm font-medium py-2 rounded-md hover:bg-yellow-700">
                Reorder Now
              </button>
            </div>
          </div>
        </Card>
      </TenantMain>
    </TenantLayout>
  )
}