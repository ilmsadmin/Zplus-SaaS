export default function SystemAdminDashboard() {
  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow">
        <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
          <h1 className="text-3xl font-bold text-gray-900">System Administration</h1>
        </div>
      </header>
      
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <div className="bg-white shadow overflow-hidden sm:rounded-lg">
            <div className="px-4 py-5 sm:px-6">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                System Management
              </h3>
              <p className="mt-1 max-w-2xl text-sm text-gray-500">
                Global system administration, tenant management, and system configuration
              </p>
            </div>
            <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                <div className="bg-blue-50 p-4 rounded-lg">
                  <h4 className="text-sm font-medium text-blue-800">Total Tenants</h4>
                  <p className="text-2xl font-bold text-blue-900">0</p>
                </div>
                <div className="bg-green-50 p-4 rounded-lg">
                  <h4 className="text-sm font-medium text-green-800">Active Plans</h4>
                  <p className="text-2xl font-bold text-green-900">3</p>
                </div>
                <div className="bg-yellow-50 p-4 rounded-lg">
                  <h4 className="text-sm font-medium text-yellow-800">System Modules</h4>
                  <p className="text-2xl font-bold text-yellow-900">6</p>
                </div>
                <div className="bg-purple-50 p-4 rounded-lg">
                  <h4 className="text-sm font-medium text-purple-800">System Health</h4>
                  <p className="text-2xl font-bold text-purple-900">100%</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}