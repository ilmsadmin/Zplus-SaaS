import { redirect } from 'next/navigation'

interface TenantSlugPageProps {
  params: Promise<{
    'tenant-slug': string
  }>
}

export default async function TenantSlugPage({ params }: TenantSlugPageProps) {
  const resolvedParams = await params
  const tenantSlug = resolvedParams['tenant-slug']
  
  // Redirect to proper tenant route structure
  redirect(`/tenant/${tenantSlug}`)
}