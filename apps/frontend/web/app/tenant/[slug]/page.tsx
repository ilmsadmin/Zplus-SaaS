interface TenantPageProps {
  params: {
    slug: string
  }
}

export default function TenantHomePage({ params }: TenantPageProps) {
  const { slug } = params
  
  return (
    <div className="min-h-screen bg-white">
      <header className="bg-blue-600 text-white">
        <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
          <h1 className="text-3xl font-bold">Welcome to {slug}</h1>
        </div>
      </header>
      
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <div className="bg-white shadow overflow-hidden sm:rounded-lg">
            <div className="px-4 py-5 sm:px-6">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                {slug} Services
              </h3>
              <p className="mt-1 max-w-2xl text-sm text-gray-500">
                Access your organization's services and applications
              </p>
            </div>
            <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                <div className="bg-gradient-to-r from-blue-50 to-blue-100 p-6 rounded-lg">
                  <h4 className="text-lg font-semibold text-blue-800 mb-2">CRM</h4>
                  <p className="text-sm text-blue-600 mb-4">Customer Relationship Management</p>
                  <button className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
                    Access CRM
                  </button>
                </div>
                <div className="bg-gradient-to-r from-green-50 to-green-100 p-6 rounded-lg">
                  <h4 className="text-lg font-semibold text-green-800 mb-2">LMS</h4>
                  <p className="text-sm text-green-600 mb-4">Learning Management System</p>
                  <button className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700">
                    Access LMS
                  </button>
                </div>
                <div className="bg-gradient-to-r from-purple-50 to-purple-100 p-6 rounded-lg">
                  <h4 className="text-lg font-semibold text-purple-800 mb-2">HRM</h4>
                  <p className="text-sm text-purple-600 mb-4">Human Resource Management</p>
                  <button className="bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700">
                    Access HRM
                  </button>
                </div>
                <div className="bg-gradient-to-r from-orange-50 to-orange-100 p-6 rounded-lg">
                  <h4 className="text-lg font-semibold text-orange-800 mb-2">POS</h4>
                  <p className="text-sm text-orange-600 mb-4">Point of Sale System</p>
                  <button className="bg-orange-600 text-white px-4 py-2 rounded hover:bg-orange-700">
                    Access POS
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}