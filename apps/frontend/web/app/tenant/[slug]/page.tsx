'use client';

import { useAuth, withAuth } from '../../lib/auth';
import Link from 'next/link';

interface TenantPageProps {
  params: {
    slug: string
  }
}

function TenantHomePage({ params }: TenantPageProps) {
  const { slug } = params;
  const { user, logout } = useAuth();
  
  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-blue-600 text-white shadow-lg">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <h1 className="text-2xl font-bold">
                {user?.tenant?.name || slug}
              </h1>
            </div>
            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-2">
                <div className="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white font-medium text-sm">
                  {user?.first_name?.charAt(0)}{user?.last_name?.charAt(0)}
                </div>
                <span className="text-sm font-medium">
                  {user?.first_name} {user?.last_name}
                </span>
                <span className="text-sm text-blue-200">({user?.role.name})</span>
              </div>
              <button
                onClick={logout}
                className="text-sm text-white hover:text-blue-200"
              >
                Logout
              </button>
            </div>
          </div>
        </div>
      </header>
      
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {/* Welcome Section */}
        <div className="px-4 py-6 sm:px-0">
          <div className="mb-8">
            <h2 className="text-3xl font-bold text-gray-900 mb-2">
              Welcome back, {user?.first_name}!
            </h2>
            <p className="text-gray-600">
              Access your organization's services and applications
            </p>
          </div>

          {/* Module Grid */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
            <Link href={`/tenant/${slug}/crm`} className="block">
              <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border-l-4 border-purple-500">
                <div className="flex items-center mb-4">
                  <div className="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center">
                    <span className="text-2xl">👥</span>
                  </div>
                  <div className="ml-4">
                    <h3 className="text-lg font-semibold text-gray-900">CRM</h3>
                    <p className="text-sm text-gray-500">Customer Management</p>
                  </div>
                </div>
                <p className="text-gray-600 text-sm mb-4">
                  Manage customer relationships, track leads, and monitor sales pipeline
                </p>
                <button className="w-full bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700 transition-colors">
                  Access CRM
                </button>
              </div>
            </Link>

            <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border-l-4 border-green-500">
              <div className="flex items-center mb-4">
                <div className="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
                  <span className="text-2xl">📚</span>
                </div>
                <div className="ml-4">
                  <h3 className="text-lg font-semibold text-gray-900">LMS</h3>
                  <p className="text-sm text-gray-500">Learning Management</p>
                </div>
              </div>
              <p className="text-gray-600 text-sm mb-4">
                Access courses, track progress, and manage educational content
              </p>
              <button className="w-full bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition-colors">
                Access LMS
              </button>
            </div>

            <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border-l-4 border-blue-500">
              <div className="flex items-center mb-4">
                <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
                  <span className="text-2xl">👔</span>
                </div>
                <div className="ml-4">
                  <h3 className="text-lg font-semibold text-gray-900">HRM</h3>
                  <p className="text-sm text-gray-500">Human Resources</p>
                </div>
              </div>
              <p className="text-gray-600 text-sm mb-4">
                Manage employee data, attendance, and HR processes
              </p>
              <button className="w-full bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition-colors">
                Access HRM
              </button>
            </div>

            <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border-l-4 border-orange-500">
              <div className="flex items-center mb-4">
                <div className="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center">
                  <span className="text-2xl">🛒</span>
                </div>
                <div className="ml-4">
                  <h3 className="text-lg font-semibold text-gray-900">POS</h3>
                  <p className="text-sm text-gray-500">Point of Sale</p>
                </div>
              </div>
              <p className="text-gray-600 text-sm mb-4">
                Process sales, manage inventory, and handle transactions
              </p>
              <button className="w-full bg-orange-600 text-white px-4 py-2 rounded hover:bg-orange-700 transition-colors">
                Access POS
              </button>
            </div>

            <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border-l-4 border-indigo-500">
              <div className="flex items-center mb-4">
                <div className="w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center">
                  <span className="text-2xl">⏰</span>
                </div>
                <div className="ml-4">
                  <h3 className="text-lg font-semibold text-gray-900">Check-in</h3>
                  <p className="text-sm text-gray-500">Attendance</p>
                </div>
              </div>
              <p className="text-gray-600 text-sm mb-4">
                Track attendance, manage schedules, and monitor work hours
              </p>
              <button className="w-full bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700 transition-colors">
                Access Check-in
              </button>
            </div>
          </div>

          {/* Quick Stats */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Your Activity</h3>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div className="text-center p-4 bg-gray-50 rounded-lg">
                <div className="text-2xl font-bold text-blue-600 mb-1">12</div>
                <div className="text-sm text-gray-600">Tasks Completed</div>
              </div>
              <div className="text-center p-4 bg-gray-50 rounded-lg">
                <div className="text-2xl font-bold text-green-600 mb-1">3</div>
                <div className="text-sm text-gray-600">Projects Active</div>
              </div>
              <div className="text-center p-4 bg-gray-50 rounded-lg">
                <div className="text-2xl font-bold text-purple-600 mb-1">28</div>
                <div className="text-sm text-gray-600">Hours This Week</div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default withAuth(TenantHomePage);