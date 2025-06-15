import { TenantLayout, TenantHeader, TenantMain } from '@/components/layouts/TenantLayout'
import { Card, StatsCard, Grid } from '@/components/ui'

interface LMSPageProps {
  params: {
    slug: string
  }
}

export default function LMSDashboard({ params }: LMSPageProps) {
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
        title="Learning Management System"
        subtitle="Online courses, training programs, and certification tracking"
      >
        {/* Progress Overview */}
        <div className="bg-gradient-to-r from-blue-600 to-purple-600 rounded-xl p-8 text-white mb-8">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 items-center">
            <div>
              <h2 className="text-2xl font-bold mb-2">Welcome back to your learning journey!</h2>
              <p className="text-blue-100 mb-4">You're making great progress. Keep up the momentum!</p>
              <div className="flex items-center space-x-4">
                <div className="text-center">
                  <div className="text-3xl font-bold">78%</div>
                  <div className="text-sm text-blue-100">Overall Progress</div>
                </div>
                <div className="text-center">
                  <div className="text-3xl font-bold">12</div>
                  <div className="text-sm text-blue-100">Courses Completed</div>
                </div>
                <div className="text-center">
                  <div className="text-3xl font-bold">3</div>
                  <div className="text-sm text-blue-100">Certificates Earned</div>
                </div>
              </div>
            </div>
            <div className="text-center">
              <div className="w-32 h-32 mx-auto bg-white bg-opacity-20 rounded-full flex items-center justify-center mb-4">
                <div className="text-6xl">🎓</div>
              </div>
              <a 
                href={`/tenant/${slug}/lms/browse`}
                className="inline-flex items-center px-6 py-3 bg-white text-blue-600 font-medium rounded-lg hover:bg-blue-50 transition-colors"
              >
                Browse Courses
              </a>
            </div>
          </div>
        </div>

        {/* LMS Stats */}
        <Grid cols={4} className="mb-8">
          <StatsCard
            title="Enrolled Courses"
            value="8"
            change="2 in progress"
            color="blue"
          />
          <StatsCard
            title="Completed Courses"
            value="12"
            change="+3 this month"
            color="green"
          />
          <StatsCard
            title="Learning Hours"
            value="164"
            change="+12h this week"
            color="purple"
          />
          <StatsCard
            title="Certificates"
            value="3"
            change="Professional level"
            color="orange"
          />
        </Grid>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
          {/* Current Courses */}
          <Card>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-lg font-semibold text-gray-900">Continue Learning</h2>
              <a href={`/tenant/${slug}/lms/my-courses`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                View all courses
              </a>
            </div>
            <div className="space-y-4">
              <div className="border border-gray-200 rounded-lg p-4">
                <div className="flex items-start justify-between mb-3">
                  <div className="flex-1">
                    <h3 className="font-medium text-gray-900 mb-1">Advanced React Development</h3>
                    <p className="text-sm text-gray-500 mb-2">Learn advanced React patterns and optimization techniques</p>
                    <div className="flex items-center space-x-2">
                      <span className="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded-full">In Progress</span>
                      <span className="text-xs text-gray-500">8 lessons remaining</span>
                    </div>
                  </div>
                  <div className="ml-4 text-right">
                    <div className="text-sm font-medium text-gray-900">75%</div>
                    <div className="w-16 bg-gray-200 rounded-full h-2 mt-1">
                      <div className="bg-blue-600 h-2 rounded-full" style={{ width: '75%' }}></div>
                    </div>
                  </div>
                </div>
                <button className="w-full bg-blue-600 text-white text-sm font-medium py-2 rounded-md hover:bg-blue-700">
                  Continue Learning
                </button>
              </div>
              
              <div className="border border-gray-200 rounded-lg p-4">
                <div className="flex items-start justify-between mb-3">
                  <div className="flex-1">
                    <h3 className="font-medium text-gray-900 mb-1">Project Management Fundamentals</h3>
                    <p className="text-sm text-gray-500 mb-2">Master the basics of effective project management</p>
                    <div className="flex items-center space-x-2">
                      <span className="text-xs bg-yellow-100 text-yellow-800 px-2 py-1 rounded-full">Started</span>
                      <span className="text-xs text-gray-500">15 lessons remaining</span>
                    </div>
                  </div>
                  <div className="ml-4 text-right">
                    <div className="text-sm font-medium text-gray-900">25%</div>
                    <div className="w-16 bg-gray-200 rounded-full h-2 mt-1">
                      <div className="bg-yellow-600 h-2 rounded-full" style={{ width: '25%' }}></div>
                    </div>
                  </div>
                </div>
                <button className="w-full bg-yellow-600 text-white text-sm font-medium py-2 rounded-md hover:bg-yellow-700">
                  Continue Learning
                </button>
              </div>
            </div>
          </Card>

          {/* Recent Activity */}
          <Card>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-lg font-semibold text-gray-900">Recent Activity</h2>
              <a href={`/tenant/${slug}/lms/activity`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                View all activity
              </a>
            </div>
            <div className="space-y-4">
              <div className="flex items-center space-x-4">
                <div className="w-10 h-10 bg-green-100 rounded-full flex items-center justify-center">
                  <svg className="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div className="flex-1">
                  <p className="text-sm font-medium text-gray-900">Completed "API Integration Patterns"</p>
                  <p className="text-sm text-gray-500">2 hours ago • React Development Course</p>
                </div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg className="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.746 0 3.332.477 4.5 1.253v13C19.832 18.477 18.246 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                  </svg>
                </div>
                <div className="flex-1">
                  <p className="text-sm font-medium text-gray-900">Started "Team Leadership Skills"</p>
                  <p className="text-sm text-gray-500">1 day ago • Project Management Course</p>
                </div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-10 h-10 bg-purple-100 rounded-full flex items-center justify-center">
                  <svg className="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div className="flex-1">
                  <p className="text-sm font-medium text-gray-900">Earned "Frontend Developer" certificate</p>
                  <p className="text-sm text-gray-500">3 days ago • React Development Course</p>
                </div>
              </div>
            </div>
          </Card>
        </div>

        {/* Available Courses */}
        <Card>
          <div className="flex items-center justify-between mb-6">
            <h2 className="text-lg font-semibold text-gray-900">Recommended for You</h2>
            <a href={`/tenant/${slug}/lms/browse`} className="text-blue-600 hover:text-blue-700 text-sm font-medium">
              Browse all courses
            </a>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div className="border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow">
              <div className="h-32 bg-gradient-to-r from-blue-400 to-blue-600 flex items-center justify-center">
                <div className="text-4xl text-white">💻</div>
              </div>
              <div className="p-4">
                <h3 className="font-semibold text-gray-900 mb-2">Full Stack Development</h3>
                <p className="text-sm text-gray-600 mb-4">Complete guide to modern web development</p>
                <div className="flex items-center justify-between">
                  <span className="text-xs bg-gray-100 text-gray-800 px-2 py-1 rounded-full">24 lessons</span>
                  <button className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                    Enroll Now
                  </button>
                </div>
              </div>
            </div>
            
            <div className="border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow">
              <div className="h-32 bg-gradient-to-r from-green-400 to-green-600 flex items-center justify-center">
                <div className="text-4xl text-white">📊</div>
              </div>
              <div className="p-4">
                <h3 className="font-semibold text-gray-900 mb-2">Data Analytics</h3>
                <p className="text-sm text-gray-600 mb-4">Learn to analyze and visualize data effectively</p>
                <div className="flex items-center justify-between">
                  <span className="text-xs bg-gray-100 text-gray-800 px-2 py-1 rounded-full">18 lessons</span>
                  <button className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                    Enroll Now
                  </button>
                </div>
              </div>
            </div>
            
            <div className="border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow">
              <div className="h-32 bg-gradient-to-r from-purple-400 to-purple-600 flex items-center justify-center">
                <div className="text-4xl text-white">🚀</div>
              </div>
              <div className="p-4">
                <h3 className="font-semibold text-gray-900 mb-2">Digital Marketing</h3>
                <p className="text-sm text-gray-600 mb-4">Master modern digital marketing strategies</p>
                <div className="flex items-center justify-between">
                  <span className="text-xs bg-gray-100 text-gray-800 px-2 py-1 rounded-full">16 lessons</span>
                  <button className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                    Enroll Now
                  </button>
                </div>
              </div>
            </div>
          </div>
        </Card>
      </TenantMain>
    </TenantLayout>
  )
}