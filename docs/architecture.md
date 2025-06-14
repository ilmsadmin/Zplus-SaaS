# Zplus SaaS - System Architecture Documentation

## Overview

Zplus SaaS is a multi-tenant platform built with modern microservices architecture, designed to support various business modules (CRM, HRM, POS, LMS) with complete tenant isolation and scalability.

## Technology Stack

### Frontend
- **Next.js** - Main web application framework
- **React** - UI component library
- **TypeScript** - Type-safe development

### Backend
- **Go Fiber** - High-performance web framework
- **GORM** - ORM for database operations
- **GraphQL/REST** - API gateway protocols

### Databases
- **PostgreSQL** - Primary relational database
- **MongoDB** - Document storage for unstructured data
- **Redis** - Caching and session storage

### Infrastructure
- **Docker** - Containerization
- **Kubernetes** - Container orchestration
- **Traefik** - Load balancing and reverse proxy

## Architecture Principles

### 1. Multi-tenancy
- **Schema-per-tenant** strategy for PostgreSQL
- **Database-per-tenant** strategy for MongoDB
- Complete data isolation between tenants

### 2. Microservices
- Service-per-module architecture
- Independent deployment and scaling
- Shared libraries and utilities

### 3. API-First Design
- GraphQL/REST gateway
- Consistent API contracts
- Version management

### 4. Security
- JWT-based authentication
- Multi-level RBAC (System → Tenant → Customer)
- Tenant context isolation

## Project Structure

```
zplus-saas/
├── apps/
│   ├── backend/
│   │   ├── gateway/            # API Gateway (Go Fiber)
│   │   ├── auth/               # Authentication Service
│   │   ├── file/               # File Management Service
│   │   ├── payment/            # Payment & Subscription
│   │   ├── crm/                # Customer Relationship Management
│   │   ├── hrm/                # Human Resource Management
│   │   ├── pos/                # Point of Sale
│   │   └── shared/             # Shared Libraries
│   └── frontend/
│       ├── web/                # Main Website (Next.js)
│       │   └── system/         # System Admin Interface
│       ├── admin/              # Tenant Admin Interface
│       └── ui/                 # Shared UI Components
├── pkg/                        # Go & JS SDKs
├── infra/                      # Infrastructure Code
│   ├── db/                     # Database migrations
│   ├── k8s/                    # Kubernetes manifests
│   ├── docker/                 # Docker configurations
│   └── ci-cd/                  # CI/CD pipelines
└── docs/                       # Documentation
```

## Service Architecture

### Gateway Service
- Entry point for all API requests
- Authentication and authorization
- Tenant context injection
- Request routing to appropriate services

### Authentication Service
- User management and authentication
- Role-based access control (RBAC)
- JWT token management
- Session handling

### Module Services (CRM/HRM/POS/LMS)
- Independent business logic
- Module-specific data models
- REST/GraphQL endpoints
- Database migrations

### File Service
- File upload/download management
- Multi-tenant file storage
- File metadata and indexing
- Integration with cloud storage

### Payment Service
- Subscription management
- Billing and invoicing
- Payment gateway integration
- Usage tracking

## Database Strategy

### PostgreSQL (Relational Data)
- System-level data (tenants, plans, subscriptions)
- Tenant-specific data (users, customers, transactions)
- Module-specific structured data
- ACID compliance for critical operations

### MongoDB (Document Data)
- File metadata and content
- Audit logs and activity tracking
- Analytics and reporting data
- Flexible schema requirements

### Redis (Cache & Sessions)
- User sessions and authentication tokens
- Frequently accessed data caching
- Real-time data and pub/sub
- Queue management for background jobs

## Multi-Tenant Architecture

### Three-Tier System

1. **System Layer**
   - Global administration
   - Tenant management
   - Plan and subscription management
   - System monitoring and analytics

2. **Tenant Layer**
   - Organization-specific administration
   - User and role management
   - Module configuration
   - Tenant-specific customization

3. **Customer Layer**
   - End-user interfaces
   - Module-specific functionality
   - Customer data and interactions
   - Role-based feature access

### Tenant Isolation

```go
// Middleware for tenant context
func TenantMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Extract tenant from subdomain/header
        tenantSlug := extractTenantSlug(c)
        
        // Validate tenant exists and is active
        tenant, err := validateTenant(tenantSlug)
        if err != nil {
            return fiber.ErrUnauthorized
        }
        
        // Set tenant context
        c.Locals("tenant", tenant)
        c.Locals("tenantDB", getTenantDB(tenant.SchemaName))
        
        return c.Next()
    }
}
```

## Deployment Strategy

### Development Environment
- Docker Compose for local development
- Single database instance
- Hot reloading for rapid development

### Staging Environment
- Kubernetes cluster
- Shared databases with staging data
- CI/CD integration testing

### Production Environment
- Multi-zone Kubernetes deployment
- Database replication and backup
- Monitoring and alerting
- Auto-scaling based on load

## Monitoring and Observability

### Metrics
- **Prometheus** - Metrics collection
- **Grafana** - Visualization and dashboards
- **AlertManager** - Alert management

### Logging
- **ELK Stack** - Centralized logging
- **Structured logging** with tenant context
- **Log aggregation** across services

### Tracing
- **Jaeger** - Distributed tracing
- **Request correlation** across services
- **Performance monitoring**

## Security Considerations

### Authentication & Authorization
- JWT tokens with short expiration
- Refresh token rotation
- Multi-factor authentication support
- Role-based access control

### Data Protection
- Encryption at rest and in transit
- PII data protection and GDPR compliance
- Tenant data isolation
- Regular security audits

### Network Security
- HTTPS/TLS everywhere
- VPC and network segmentation
- WAF (Web Application Firewall)
- DDoS protection

## Scalability Patterns

### Horizontal Scaling
- Stateless service design
- Load balancing across instances
- Database read replicas
- CDN for static assets

### Vertical Scaling
- Resource allocation based on tenant usage
- Auto-scaling policies
- Performance monitoring and optimization

### Database Scaling
- Connection pooling
- Query optimization
- Caching strategies
- Sharding for large tenants

## Development Guidelines

### Code Organization
- Domain-driven design principles
- Clean architecture patterns
- Dependency injection
- Unit and integration testing

### API Design
- RESTful conventions
- GraphQL schema design
- API versioning strategy
- Comprehensive documentation

### Database Design
- Normalized relational schemas
- Denormalized document schemas
- Indexing strategies
- Migration procedures

This architecture provides a solid foundation for building a scalable, secure, and maintainable multi-tenant SaaS platform.