'use client';

import { useAuth, withAuth } from '../../../lib/auth';
import Link from 'next/link';

interface CRMPageProps {
  params: {
    slug: string
  }
}

function CRMDashboard({ params }: CRMPageProps) {
  const { slug } = params;
  const { user, logout } = useAuth();
  
  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <Link href={`/tenant/${slug}`} className="text-purple-600 hover:text-purple-800 mr-4">
                ← Back to Portal
              </Link>
              <h1 className="text-2xl font-bold text-purple-600">CRM Dashboard</h1>
              <nav className="ml-8 flex space-x-8">
                <a href="#" className="text-gray-600 hover:text-gray-900">Dashboard</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Leads</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Customers</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Opportunities</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Reports</a>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-2">
                <div className="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center text-white font-medium text-sm">
                  {user?.first_name?.charAt(0)}{user?.last_name?.charAt(0)}
                </div>
                <span className="text-sm font-medium text-gray-700">
                  {user?.first_name} {user?.last_name}
                </span>
                <span className="text-sm text-gray-500">({user?.role.name})</span>
              </div>
              <button
                onClick={logout}
                className="text-sm text-gray-600 hover:text-gray-900"
              >
                Logout
              </button>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {/* Page Header */}
        <div className="px-4 py-6 sm:px-0">
          <div className="flex justify-between items-center">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">CRM Dashboard</h1>
              <p className="mt-2 text-gray-600">Theo dõi leads, opportunities và hoạt động bán hàng</p>
            </div>
            <div className="flex space-x-3">
              <button className="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50">
                📊 Export Report
              </button>
              <button className="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-purple-600 hover:bg-purple-700">
                ➕ Add Lead
              </button>
            </div>
          </div>
        </div>

        {/* Stats Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div className="bg-white overflow-hidden shadow rounded-lg">
            <div className="p-5">
              <div className="flex items-center">
                <div className="flex-shrink-0">
                  <div className="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm font-medium">🎯</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Total Leads</dt>
                    <dd className="text-3xl font-bold text-gray-900">156</dd>
                    <dd className="text-sm text-blue-600">+12% từ tháng trước</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-white overflow-hidden shadow rounded-lg">
            <div className="p-5">
              <div className="flex items-center">
                <div className="flex-shrink-0">
                  <div className="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm font-medium">💼</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Opportunities</dt>
                    <dd className="text-3xl font-bold text-gray-900">34</dd>
                    <dd className="text-sm text-green-600">+8% từ tháng trước</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-white overflow-hidden shadow rounded-lg">
            <div className="p-5">
              <div className="flex items-center">
                <div className="flex-shrink-0">
                  <div className="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm font-medium">👤</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Customers</dt>
                    <dd className="text-3xl font-bold text-gray-900">1,247</dd>
                    <dd className="text-sm text-purple-600">+18 khách hàng mới</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-white overflow-hidden shadow rounded-lg">
            <div className="p-5">
              <div className="flex items-center">
                <div className="flex-shrink-0">
                  <div className="w-8 h-8 bg-yellow-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm font-medium">💰</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Monthly Revenue</dt>
                    <dd className="text-3xl font-bold text-gray-900">$45,230</dd>
                    <dd className="text-sm text-yellow-600">+25% từ tháng trước</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Content Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* Sales Pipeline */}
          <div className="bg-white shadow rounded-lg">
            <div className="px-4 py-5 sm:px-6 border-b border-gray-200">
              <h3 className="text-lg leading-6 font-medium text-gray-900">Sales Pipeline</h3>
            </div>
            <div className="px-4 py-5 sm:p-6">
              <div className="space-y-4">
                <div className="flex justify-between items-center p-3 bg-blue-50 rounded-lg">
                  <div>
                    <h4 className="font-medium text-gray-900">Qualified Leads</h4>
                    <p className="text-sm text-gray-500">42 leads trong giai đoạn này</p>
                  </div>
                  <div className="text-right">
                    <div className="text-lg font-bold text-blue-600">$128,400</div>
                    <div className="text-sm text-gray-500">Potential value</div>
                  </div>
                </div>
                
                <div className="flex justify-between items-center p-3 bg-yellow-50 rounded-lg">
                  <div>
                    <h4 className="font-medium text-gray-900">Proposal Sent</h4>
                    <p className="text-sm text-gray-500">18 opportunities</p>
                  </div>
                  <div className="text-right">
                    <div className="text-lg font-bold text-yellow-600">$89,200</div>
                    <div className="text-sm text-gray-500">Potential value</div>
                  </div>
                </div>
                
                <div className="flex justify-between items-center p-3 bg-green-50 rounded-lg">
                  <div>
                    <h4 className="font-medium text-gray-900">Negotiation</h4>
                    <p className="text-sm text-gray-500">12 deals closing soon</p>
                  </div>
                  <div className="text-right">
                    <div className="text-lg font-bold text-green-600">$156,800</div>
                    <div className="text-sm text-gray-500">Potential value</div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Recent Activities */}
          <div className="bg-white shadow rounded-lg">
            <div className="px-4 py-5 sm:px-6 border-b border-gray-200">
              <h3 className="text-lg leading-6 font-medium text-gray-900">Recent Activities</h3>
            </div>
            <div className="px-4 py-5 sm:p-6">
              <div className="space-y-4">
                <div className="flex items-start space-x-3">
                  <div className="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm">📞</span>
                  </div>
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-900">Called John Smith from ABC Corp</p>
                    <p className="text-sm text-gray-500">2 giờ trước • Lead conversion call</p>
                  </div>
                </div>
                
                <div className="flex items-start space-x-3">
                  <div className="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm">📧</span>
                  </div>
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-900">Sent proposal to XYZ Industries</p>
                    <p className="text-sm text-gray-500">4 giờ trước • $45,000 deal value</p>
                  </div>
                </div>
                
                <div className="flex items-start space-x-3">
                  <div className="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm">🤝</span>
                  </div>
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-900">Meeting scheduled with Tech Solutions</p>
                    <p className="text-sm text-gray-500">1 ngày trước • Product demo</p>
                  </div>
                </div>
                
                <div className="flex items-start space-x-3">
                  <div className="w-8 h-8 bg-yellow-500 rounded-full flex items-center justify-center">
                    <span className="text-white text-sm">➕</span>
                  </div>
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-900">New lead: Digital Marketing Co</p>
                    <p className="text-sm text-gray-500">2 ngày trước • Inbound lead from website</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default withAuth(CRMDashboard);