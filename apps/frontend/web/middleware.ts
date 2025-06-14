import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl
  const hostname = request.headers.get('host') || ''
  
  // Extract subdomain
  const subdomain = hostname.split('.')[0]
  const isSubdomain = hostname.includes('.') && subdomain !== 'www' && subdomain !== 'localhost'
  
  // Handle subdomain routing for tenant pages
  if (isSubdomain && !pathname.startsWith('/admin')) {
    // Tenant customer pages: tenant-slug.domain.com -> /tenant/[slug]
    const url = request.nextUrl.clone()
    url.pathname = `/tenant/${subdomain}${pathname}`
    return NextResponse.rewrite(url)
  }
  
  // Handle subdomain admin routing
  if (isSubdomain && pathname.startsWith('/admin')) {
    // Tenant admin pages: tenant-slug.domain.com/admin -> /tenant/[slug]/admin
    const url = request.nextUrl.clone()
    url.pathname = `/tenant/${subdomain}/admin${pathname.replace('/admin', '')}`
    return NextResponse.rewrite(url)
  }
  
  return NextResponse.next()
}

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    '/((?!api|_next/static|_next/image|favicon.ico).*)',
  ],
}