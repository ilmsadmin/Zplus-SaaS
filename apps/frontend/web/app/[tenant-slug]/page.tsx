import { redirect } from 'next/navigation'

interface TenantSlugPageProps {
  params: {
    'tenant-slug': string
  }
}

export default function TenantSlugPage({ params }: TenantSlugPageProps) {
  const tenantSlug = params['tenant-slug']
  
  // Redirect to proper tenant route structure
  redirect(`/tenant/${tenantSlug}`)
}