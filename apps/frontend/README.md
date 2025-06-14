# Frontend Applications

This directory contains all frontend applications for the Zplus SaaS platform.

## Applications

### Web Application (Port: 3000)
- Main website and landing pages
- Built with Next.js 14 and React 18
- Supports both system and tenant interfaces

#### System Admin Interface (Port: 3001)
- Global system administration
- Tenant management
- System-wide settings and monitoring

### Tenant Admin Interface (Port: 3002)
- Organization-specific administration
- User management within tenant
- Module configuration per tenant

### Shared UI Components
- Reusable React components
- Design system implementation
- Common utilities and hooks

## Development

Each application is a standalone Next.js project:

```bash
# Web application
cd apps/frontend/web
npm install
npm run dev

# System admin
cd apps/frontend/web/system  
npm install
npm run dev

# Tenant admin
cd apps/frontend/admin
npm install
npm run dev
```

## Architecture

- **Next.js 14** with App Router
- **TypeScript** for type safety
- **Tailwind CSS** for styling
- **Shared UI components** for consistency

## Multi-Tenant Architecture

- System-level interfaces for global management
- Tenant-level interfaces with isolated data
- Dynamic theming based on tenant configuration
- Subdomain-based tenant routing