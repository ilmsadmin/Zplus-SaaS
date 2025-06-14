'use client';

import { useAuth, withAuth } from '../../../lib/auth';

interface TenantAdminPageProps {
  params: {
    slug: string
  }
}

function TenantAdminDashboard({ params }: TenantAdminPageProps) {
  const { slug } = params;
  const { user, logout } = useAuth();

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <h1 className="text-2xl font-bold text-orange-600">
                {user?.tenant?.name || slug} Admin
              </h1>
              <nav className="ml-8 flex space-x-8">
                <a href="#" className="text-gray-600 hover:text-gray-900">Dashboard</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Team</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Customers</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Modules</a>
                <a href="#" className="text-gray-600 hover:text-gray-900">Settings</a>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-2">
                <div className="w-8 h-8 bg-orange-500 rounded-full flex items-center justify-center text-white font-medium text-sm">
                  TA
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
              <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
              <p className="mt-2 text-gray-600">Chào mừng trở lại! Đây là tổng quan về tổ chức của bạn.</p>
            </div>
            <div className="flex space-x-3">
              <button className="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50">
                📊 Reports
              </button>
              <button className="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700">
                ➕ Add User
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
                    <span className="text-white text-sm font-medium">👥</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Total Users</dt>
                    <dd className="text-3xl font-bold text-gray-900">24</dd>
                    <dd className="text-sm text-gray-600">Active users trong tổ chức</dd>
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
                    <span className="text-white text-sm font-medium">👤</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Customers</dt>
                    <dd className="text-3xl font-bold text-gray-900">1,247</dd>
                    <dd className="text-sm text-green-600">+18 khách hàng mới tuần này</dd>
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
                    <span className="text-white text-sm font-medium">📦</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Active Modules</dt>
                    <dd className="text-3xl font-bold text-gray-900">5</dd>
                    <dd className="text-sm text-gray-600">CRM, LMS, POS, HRM, Checkin</dd>
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
                    <span className="text-white text-sm font-medium">💾</span>
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">Storage Used</dt>
                    <dd className="text-3xl font-bold text-gray-900">2.4 GB</dd>
                    <dd className="text-sm text-gray-600">của 10 GB limit</dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Content Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* Team Members */}
          <div className="bg-white shadow rounded-lg">
            <div className="px-4 py-5 sm:px-6 border-b border-gray-200">
              <div className="flex justify-between items-center">
                <h3 className="text-lg leading-6 font-medium text-gray-900">Team Members</h3>
                <button className="text-sm text-blue-600 hover:text-blue-900">View All</button>
              </div>
            </div>
            <div className="px-4 py-5 sm:p-6">
              <div className="space-y-4">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-blue-500 rounded-full flex items-center justify-center text-white font-medium">
                    {user?.first_name?.charAt(0)}{user?.last_name?.charAt(0)}
                  </div>
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-900">{user?.first_name} {user?.last_name}</p>
                    <p className="text-sm text-gray-500">{user?.role.name}</p>
                  </div>
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                    Active
                  </span>
                </div>
                
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-green-500 rounded-full flex items-center justify-center text-white font-medium">
                    JS
                  </div>
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-900">Jane Smith</p>
                    <p className="text-sm text-gray-500">Manager</p>
                  </div>
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                    Active
                  </span>
                </div>
                
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-purple-500 rounded-full flex items-center justify-center text-white font-medium">
                    MB
                  </div>
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-900">Mike Brown</p>
                    <p className="text-sm text-gray-500">User</p>
                  </div>
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-yellow-100 text-yellow-800">
                    Pending
                  </span>
                </div>
              </div>
            </div>
          </div>

          {/* Active Modules */}
          <div className="bg-white shadow rounded-lg">
            <div className="px-4 py-5 sm:px-6 border-b border-gray-200">
              <h3 className="text-lg leading-6 font-medium text-gray-900">Active Modules</h3>
            </div>
            <div className="px-4 py-5 sm:p-6">
              <div className="grid grid-cols-2 gap-4">
                <div className="text-center p-4 border border-gray-200 rounded-lg">
                  <div className="text-2xl mb-2">👥</div>
                  <h4 className="font-medium text-gray-900">CRM</h4>
                  <p className="text-sm text-gray-500">Customer Management</p>
                </div>
                
                <div className="text-center p-4 border border-gray-200 rounded-lg">
                  <div className="text-2xl mb-2">📚</div>
                  <h4 className="font-medium text-gray-900">LMS</h4>
                  <p className="text-sm text-gray-500">Learning Management</p>
                </div>
                
                <div className="text-center p-4 border border-gray-200 rounded-lg">
                  <div className="text-2xl mb-2">🛒</div>
                  <h4 className="font-medium text-gray-900">POS</h4>
                  <p className="text-sm text-gray-500">Point of Sale</p>
                </div>
                
                <div className="text-center p-4 border border-gray-200 rounded-lg">
                  <div className="text-2xl mb-2">👔</div>
                  <h4 className="font-medium text-gray-900">HRM</h4>
                  <p className="text-sm text-gray-500">Human Resources</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default withAuth(TenantAdminDashboard);