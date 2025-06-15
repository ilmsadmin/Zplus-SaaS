interface ModulesPageProps {
  params: Promise<{
    slug: string
  }>
}

export default async function TenantAdminModules({ params }: ModulesPageProps) {
  const resolvedParams = await params
  const { slug } = resolvedParams
  const tenantName = slug.charAt(0).toUpperCase() + slug.slice(1) + ' Corporation'

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-4">
            <div className="flex items-center space-x-4">
              <div className="bg-indigo-600 text-white p-2 rounded-lg">
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                </svg>
              </div>
              <div>
                <h1 className="text-xl font-semibold text-gray-900">Module Management</h1>
                <p className="text-sm text-gray-600">{tenantName} - Configure available modules</p>
              </div>
            </div>
            <div className="flex items-center space-x-4">
              <button className="bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700 transition-colors">
                Save Changes
              </button>
              <a href={`/tenant/${slug}/admin`} className="text-sm text-blue-600 hover:text-blue-700">
                ← Back to Dashboard
              </a>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="mb-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-2">Available Modules</h2>
          <p className="text-gray-600">Enable or disable modules for your organization</p>
        </div>

        {/* Module Status Overview */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Total Modules</p>
                <p className="text-2xl font-bold text-gray-900">4</p>
              </div>
              <div className="bg-blue-100 text-blue-600 p-3 rounded-lg">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                </svg>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Active Modules</p>
                <p className="text-2xl font-bold text-gray-900">3</p>
              </div>
              <div className="bg-green-100 text-green-600 p-3 rounded-lg">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Inactive Modules</p>
                <p className="text-2xl font-bold text-gray-900">1</p>
              </div>
              <div className="bg-red-100 text-red-600 p-3 rounded-lg">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-xl shadow-sm border p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Usage Rate</p>
                <p className="text-2xl font-bold text-gray-900">75%</p>
              </div>
              <div className="bg-purple-100 text-purple-600 p-3 rounded-lg">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
            </div>
          </div>
        </div>

        {/* Module Configuration */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* CRM Module */}
          <div className="bg-white rounded-xl shadow-sm border overflow-hidden">
            <div className="p-6 border-b bg-gradient-to-r from-blue-50 to-indigo-50">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-4">
                  <div className="bg-blue-600 text-white p-3 rounded-lg">
                    <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 515.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 919.288 0M15 7a3 3 0 11-6 0 3 3 0 616 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                    </svg>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-gray-900">CRM Module</h3>
                    <p className="text-sm text-gray-600">Customer Relationship Management</p>
                  </div>
                </div>
                <div className="flex items-center">
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" defaultChecked />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                </div>
              </div>
            </div>
            <div className="p-6">
              <div className="grid grid-cols-2 gap-4 mb-4">
                <div>
                  <span className="text-sm text-gray-500">Active Users</span>
                  <p className="text-lg font-semibold text-gray-900">45</p>
                </div>
                <div>
                  <span className="text-sm text-gray-500">Last Activity</span>
                  <p className="text-lg font-semibold text-gray-900">2 min ago</p>
                </div>
              </div>
              <div className="flex space-x-2">
                <a href={`/tenant/${slug}/crm`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                  Open Module
                </a>
                <span className="text-gray-300">|</span>
                <button className="text-gray-600 hover:text-gray-700 text-sm font-medium">
                  Configure
                </button>
              </div>
            </div>
          </div>

          {/* LMS Module */}
          <div className="bg-white rounded-xl shadow-sm border overflow-hidden">
            <div className="p-6 border-b bg-gradient-to-r from-green-50 to-emerald-50">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-4">
                  <div className="bg-green-600 text-white p-3 rounded-lg">
                    <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                    </svg>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-gray-900">LMS Module</h3>
                    <p className="text-sm text-gray-600">Learning Management System</p>
                  </div>
                </div>
                <div className="flex items-center">
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" defaultChecked />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
                  </label>
                </div>
              </div>
            </div>
            <div className="p-6">
              <div className="grid grid-cols-2 gap-4 mb-4">
                <div>
                  <span className="text-sm text-gray-500">Active Users</span>
                  <p className="text-lg font-semibold text-gray-900">32</p>
                </div>
                <div>
                  <span className="text-sm text-gray-500">Last Activity</span>
                  <p className="text-lg font-semibold text-gray-900">5 min ago</p>
                </div>
              </div>
              <div className="flex space-x-2">
                <a href={`/tenant/${slug}/lms`} className="text-green-600 hover:text-green-700 text-sm font-medium">
                  Open Module
                </a>
                <span className="text-gray-300">|</span>
                <button className="text-gray-600 hover:text-gray-700 text-sm font-medium">
                  Configure
                </button>
              </div>
            </div>
          </div>

          {/* HRM Module */}
          <div className="bg-white rounded-xl shadow-sm border overflow-hidden">
            <div className="p-6 border-b bg-gradient-to-r from-purple-50 to-pink-50">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-4">
                  <div className="bg-purple-600 text-white p-3 rounded-lg">
                    <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                    </svg>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-gray-900">HRM Module</h3>
                    <p className="text-sm text-gray-600">Human Resource Management</p>
                  </div>
                </div>
                <div className="flex items-center">
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" defaultChecked />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
                  </label>
                </div>
              </div>
            </div>
            <div className="p-6">
              <div className="grid grid-cols-2 gap-4 mb-4">
                <div>
                  <span className="text-sm text-gray-500">Active Users</span>
                  <p className="text-lg font-semibold text-gray-900">28</p>
                </div>
                <div>
                  <span className="text-sm text-gray-500">Last Activity</span>
                  <p className="text-lg font-semibold text-gray-900">1 hour ago</p>
                </div>
              </div>
              <div className="flex space-x-2">
                <a href={`/tenant/${slug}/hrm`} className="text-purple-600 hover:text-purple-700 text-sm font-medium">
                  Open Module
                </a>
                <span className="text-gray-300">|</span>
                <button className="text-gray-600 hover:text-gray-700 text-sm font-medium">
                  Configure
                </button>
              </div>
            </div>
          </div>

          {/* POS Module */}
          <div className="bg-white rounded-xl shadow-sm border overflow-hidden opacity-60">
            <div className="p-6 border-b bg-gradient-to-r from-orange-50 to-red-50">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-4">
                  <div className="bg-orange-600 text-white p-3 rounded-lg">
                    <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                    </svg>
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-gray-900">POS Module</h3>
                    <p className="text-sm text-gray-600">Point of Sale System</p>
                  </div>
                </div>
                <div className="flex items-center">
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-orange-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-orange-600"></div>
                  </label>
                </div>
              </div>
            </div>
            <div className="p-6">
              <div className="grid grid-cols-2 gap-4 mb-4">
                <div>
                  <span className="text-sm text-gray-500">Active Users</span>
                  <p className="text-lg font-semibold text-gray-900">0</p>
                </div>
                <div>
                  <span className="text-sm text-gray-500">Status</span>
                  <p className="text-lg font-semibold text-red-600">Disabled</p>
                </div>
              </div>
              <div className="flex space-x-2">
                <button className="text-gray-400 text-sm font-medium cursor-not-allowed">
                  Open Module
                </button>
                <span className="text-gray-300">|</span>
                <button className="text-gray-600 hover:text-gray-700 text-sm font-medium">
                  Configure
                </button>
              </div>
            </div>
          </div>
        </div>

        {/* Module Permissions */}
        <div className="mt-8 bg-white rounded-xl shadow-sm border overflow-hidden">
          <div className="p-6 border-b">
            <h3 className="text-lg font-semibold text-gray-900">Module Permissions</h3>
            <p className="text-gray-600 mt-1">Configure role-based access to modules</p>
          </div>
          
          <div className="overflow-x-auto">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Role
                  </th>
                  <th className="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
                    CRM
                  </th>
                  <th className="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
                    LMS
                  </th>
                  <th className="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
                    HRM
                  </th>
                  <th className="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
                    POS
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                <tr>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    Administrator
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-blue-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-green-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-purple-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-orange-600 rounded" />
                  </td>
                </tr>
                <tr>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    Manager
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-blue-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-green-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-purple-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-orange-600 rounded" />
                  </td>
                </tr>
                <tr>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    Employee
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-blue-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-green-600 rounded" defaultChecked />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-purple-600 rounded" />
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-center">
                    <input type="checkbox" className="h-4 w-4 text-orange-600 rounded" />
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </main>
    </div>
  )
}