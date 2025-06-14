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