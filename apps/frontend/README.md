# Frontend Applications

This directory contains all frontend applications for the Zplus SaaS platform.

## Applications

### Web Application (Port: 3000)
- Main website and landing pages
- Built with Next.js 14 and React 18
- **Unified Multi-Tenant Routing**: 
  - `/admin` - System administration
  - `/tenant/[slug]` - Tenant customer pages
  - `/tenant/[slug]/admin` - Tenant administration
  - `/[tenant-slug]` - Alternative tenant routing (redirects)
- **Middleware-based subdomain support**:
  - `tenant.domain.com` → `/tenant/[slug]`
  - `tenant.domain.com/admin` → `/tenant/[slug]/admin`

### Legacy Applications (Deprecated)
#### System Admin Interface (Port: 3001)
- ⚠️ **Deprecated**: Use `/admin` route in main web app instead
- Global system administration
- Tenant management
- System-wide settings and monitoring

#### Tenant Admin Interface (Port: 3002)
- ⚠️ **Deprecated**: Use `/tenant/[slug]/admin` route in main web app instead
- Organization-specific administration
- User management within tenant
- Module configuration per tenant

### Shared UI Components
- Reusable React components
- Design system implementation
- Common utilities and hooks

## Development

### Main Web Application (Recommended)
```bash
cd apps/frontend/web
npm install
npm run dev  # Runs on port 3000
```

### Legacy Applications (Deprecated)
```bash
# System admin (deprecated - use /admin instead)
cd apps/frontend/web/system  
npm install
npm run dev  # Port 3001

# Tenant admin (deprecated - use /tenant/[slug]/admin instead)
cd apps/frontend/admin
npm install
npm run dev  # Port 3002
```

## Architecture

- **Next.js 14** with App Router
- **TypeScript** for type safety
- **Tailwind CSS** for styling
- **Shared UI components** for consistency

## Multi-Tenant Architecture

- **Unified Routing Structure**: Single Next.js app handles all routing patterns
- **System-level interfaces** for global management (`/admin`)
- **Tenant-level interfaces** with isolated data (`/tenant/[slug]` and `/tenant/[slug]/admin`)
- **Dynamic theming** based on tenant configuration
- **Subdomain-based tenant routing** via Next.js middleware
- **Backwards compatibility** with alternative routing patterns (`/[tenant-slug]`)