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
- **GraphQL-first** - Primary API protocol with gqlgen
- **REST API** - Backward compatibility and simple operations

### Databases
- **PostgreSQL** - Primary relational database
- **MongoDB** - Document storage for unstructured data
- **Redis** - Caching and session storage

### Infrastructure
- **Docker** - Containerization
- **Kubernetes** - Container orchestration
- **Traefik** - Load balancing, reverse proxy, and multi-tenant routing

### Message Queue Strategy
- **Redis Queue** - Background job processing and async operations
- **Redis Pub/Sub** - Real-time notifications and live updates

*Note: Kafka is not currently part of the architecture. Redis provides sufficient message processing capabilities for our current scale and use cases. Kafka would be considered for future requirements involving high-throughput event streaming (>100K msg/s) or event sourcing patterns.*

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
- **GraphQL-first** approach for modern frontend development
- **REST API** compatibility for legacy systems and simple operations
- **Multi-tenant aware** - All operations scoped to tenant context
- **Real-time subscriptions** for live updates
- Consistent API contracts and schema evolution
- Version management and backward compatibility

### 4. Security
- JWT-based authentication
- Multi-level RBAC (System → Tenant → Customer)
- Tenant context isolation

## Project Structure

```
zplus-saas/
├── apps/
│   ├── backend/
│   │   ├── gateway/            # API Gateway (Go Fiber + GraphQL)
│   │   │   ├── examples/       # GraphQL query examples
│   │   │   ├── generated/      # Auto-generated GraphQL code
│   │   │   ├── middleware/     # Authentication & tenant middleware
│   │   │   ├── resolver/       # GraphQL resolvers
│   │   │   ├── schema/         # GraphQL schema definitions
│   │   │   └── types/          # Custom types and models
│   │   ├── auth/               # Authentication Service
│   │   ├── file/               # File Management Service
│   │   ├── payment/            # Payment & Subscription
│   │   ├── crm/                # Customer Relationship Management
│   │   ├── hrm/                # Human Resource Management
│   │   ├── pos/                # Point of Sale
│   │   └── shared/             # Shared Libraries
│   └── frontend/
│       ├── web/                # Main Next.js Application
│       │   └── app/            # App Router (Next.js 13+)
│       │       ├── [tenant-slug]/     # Dynamic tenant routing
│       │       ├── tenant/            # Tenant management
│       │       │   └── [slug]/        # Tenant-specific pages
│       │       │       └── admin/     # Tenant admin interface
│       │       └── admin/             # System admin interface
│       └── ui/                 # Shared UI Components
│           └── components/     # Reusable React components
├── pkg/                        # Shared Go packages & utilities
│   ├── auth/                   # JWT authentication utilities
│   ├── database/               # Database connection utilities
│   └── utils/                  # Helper functions
├── infra/                      # Infrastructure Code
│   ├── db/                     # Database migrations & schemas
│   ├── k8s/                    # Kubernetes manifests
│   ├── docker/                 # Docker configurations & Traefik
│   └── ci-cd/                  # CI/CD pipelines
├── docs/                       # Documentation
│   ├── architecture.md         # System architecture (English)
│   ├── thiet-ke-kien-truc-du-an.md    # Architecture design (Vietnamese)
│   ├── thiet-ke-kien-truc-database.md # Database design
│   ├── thiet-ke-tong-quan-du-an.md    # Project overview
│   └── thiet-ke-ux-ui.md      # UX/UI design guidelines
└── mock/                       # Static HTML mockups & prototypes
    ├── index.html             # Landing page mockup
    ├── system-admin-dashboard.html    # System admin UI mockup
    ├── tenant-admin-dashboard.html    # Tenant admin UI mockup
    ├── customer-crm-dashboard.html    # CRM module mockup
    └── customer-lms-portal.html       # LMS module mockup
```

## Service Architecture

### Gateway Service
- **GraphQL-first** entry point for all API requests
- **Real-time subscriptions** for live data updates
- **REST API** endpoints for backward compatibility
- Authentication and authorization (JWT + RBAC)
- Multi-tenant context injection and validation
- Request routing to appropriate microservices
- Rate limiting and security middleware

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
- Real-time data and pub/sub messaging
- Background job queue management
- Async processing workflows

## Message Queue & Event Processing

### Redis-based Message Processing

**Queue Management**:
```go
// Background job processing
type Job struct {
    ID       string                 `json:"id"`
    Type     string                 `json:"type"`
    Payload  map[string]interface{} `json:"payload"`
    TenantID string                 `json:"tenant_id"`
    Priority int                    `json:"priority"`
}

// Common job types:
// - Email notifications
// - PDF report generation
// - Data exports/imports
// - Webhook deliveries
// - Audit log processing
```

**Pub/Sub for Real-time Updates**:
```go
// Real-time notification channels
channels := []string{
    "tenant:{slug}:notifications",
    "tenant:{slug}:chat",
    "system:maintenance",
    "module:{module_name}:updates",
}
```

### When to Consider Kafka

Current architecture uses Redis for message processing. Kafka would be considered when:

- **High Throughput**: > 100K messages/second required
- **Event Sourcing**: Need for event store and replay capabilities  
- **Cross-System Events**: Complex event streaming between services
- **Analytics Pipeline**: Real-time data processing and analytics
- **Audit Compliance**: Immutable event log requirements

**Comparison Matrix**:

| Feature | Redis Queue | Kafka | Current Choice |
|---------|-------------|-------|----------------|
| Setup Complexity | Low | High | ✅ Redis |
| Latency | < 1ms | 5-10ms | ✅ Redis |
| Throughput | 100K msg/s | 1M+ msg/s | Redis sufficient |
| Persistence | Optional | Durable | Redis adequate |
| Use Case Fit | Simple queues | Event streaming | ✅ Multi-tenant SaaS |

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

### Traefik Configuration

**Multi-tenant Routing with Traefik**:
```yaml
# Traefik middleware for tenant extraction
http:
  middlewares:
    tenant-headers:
      headers:
        customRequestHeaders:
          X-Tenant-ID: "{{ .Request.Host | regexReplaceAll `^([^.]+)\..*` `$1` }}"
    
    tenant-auth:
      forwardAuth:
        address: "http://auth-service:8081/auth/validate"
        authRequestHeaders:
          - "X-Tenant-ID"
          
  routers:
    api-router:
      rule: "Host(`{subdomain:[a-z0-9-]+}.zplus.com`)"
      service: gateway-service
      middlewares:
        - tenant-headers
        - tenant-auth
      tls:
        certResolver: letsencrypt
```

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

## API Design Strategy

### GraphQL-First Approach

Zplus SaaS prioritizes GraphQL as the primary API protocol while maintaining REST compatibility for specific use cases.

#### Why GraphQL-First?

1. **Frontend Efficiency**: Single request for multiple data sources
2. **Type Safety**: Strong typing with auto-generated TypeScript types
3. **Real-time Capabilities**: Built-in subscription support
4. **Multi-tenant Aware**: Schema-level tenant isolation
5. **Developer Experience**: Interactive playground and introspection

#### GraphQL Schema Design

```graphql
# Multi-tenant base interface
interface TenantEntity {
  id: ID!
  tenantId: TenantID!
  createdAt: DateTime!
  updatedAt: DateTime!
}

# Example: Customer entity
type Customer implements TenantEntity {
  id: ID!
  tenantId: TenantID!
  name: String!
  email: String
  status: CustomerStatus!
  # ... other fields
}
```

#### Multi-Tenant GraphQL Context

```go
// Request context with tenant isolation
type RequestContext struct {
    Tenant *TenantContext `json:"tenant"`
    User   *UserContext   `json:"user"`
}

// Resolver authorization
func (r *customerResolver) Customers(ctx context.Context) ([]*Customer, error) {
    reqCtx := getRequestContext(ctx)
    
    // Validate tenant access
    if err := r.requireTenantAuth(reqCtx); err != nil {
        return nil, err
    }
    
    // Query scoped to tenant
    return r.customerService.GetByTenant(reqCtx.Tenant.ID)
}
```

### REST API Compatibility

REST endpoints are provided for:

1. **Simple CRUD operations** where GraphQL might be overkill
2. **Third-party integrations** that prefer REST
3. **Health checks and monitoring**
4. **File uploads** and binary data operations

#### REST API Design Pattern

```http
# Tenant-scoped REST endpoints
GET    /api/v1/customers
POST   /api/v1/customers
GET    /api/v1/customers/{id}
PUT    /api/v1/customers/{id}
DELETE /api/v1/customers/{id}

# Headers required
Authorization: Bearer {jwt_token}
X-Tenant-ID: {tenant_slug}
```

### Frontend-Backend Integration

#### GraphQL Code Generation

```typescript
// Auto-generated TypeScript types
export interface Customer {
  id: string;
  tenantId: string;
  name: string;
  email?: string;
  status: CustomerStatus;
  createdAt: string;
  updatedAt: string;
}

// Auto-generated React hooks
export const useGetCustomersQuery = (variables?: GetCustomersQueryVariables) => {
  return useQuery<GetCustomersQuery, GetCustomersQueryVariables>(
    GetCustomersDocument,
    variables
  );
};
```

#### Real-time Updates with Subscriptions

```typescript
// Frontend subscription for live updates
const { data, loading } = useSubscription(NOTIFICATIONS_SUBSCRIPTION, {
  variables: { tenantId: currentTenant.id }
});

// GraphQL subscription
subscription NotificationUpdates($tenantId: TenantID!) {
  notifications(tenantId: $tenantId) {
    id
    type
    title
    message
    createdAt
  }
}
```

### Authentication & Authorization

#### JWT Token Structure

```json
{
  "sub": "user_id",
  "tenant_id": "tenant_slug", 
  "email": "user@example.com",
  "roles": ["admin", "user"],
  "permissions": [
    "customers:read",
    "customers:write",
    "products:read"
  ],
  "exp": 1640995200,
  "iat": 1640908800
}
```

#### GraphQL Authorization Patterns

```go
// Resolver-level authorization
func (r *mutationResolver) CreateCustomer(ctx context.Context, input CreateCustomerInput) (*Customer, error) {
    reqCtx := getRequestContext(ctx)
    
    // Check specific permission
    if err := r.requirePermission(reqCtx, "customers", "write"); err != nil {
        return nil, err
    }
    
    // Business logic...
}

// Field-level authorization
func (r *customerResolver) SensitiveData(ctx context.Context, obj *Customer) (*string, error) {
    reqCtx := getRequestContext(ctx)
    
    if !reqCtx.User.HasRole("admin") {
        return nil, nil // Hide field for non-admins
    }
    
    return &obj.SensitiveData, nil
}
```

### Data Loading Patterns

#### DataLoader for N+1 Prevention

```go
// Batch loading for efficient database queries
type UserLoader struct {
    userService UserService
}

func (l *UserLoader) Load(ctx context.Context, userID string) (*User, error) {
    // Batched loading logic to prevent N+1 queries
    return l.userService.LoadUser(ctx, userID)
}
```

#### Pagination Strategy

```graphql
# Cursor-based pagination for large datasets
type CustomerConnection {
  edges: [CustomerEdge!]!
  pageInfo: PageInfo!
  totalCount: Int!
}

type CustomerEdge {
  node: Customer!
  cursor: String!
}
```

### Error Handling Strategy

#### GraphQL Error Extensions

```json
{
  "data": null,
  "errors": [
    {
      "message": "Access denied",
      "path": ["customers"],
      "extensions": {
        "code": "FORBIDDEN",
        "tenantId": "acme",
        "permission": "customers:read"
      }
    }
  ]
}
```

#### Custom Error Types

```go
// Structured error handling
type GraphQLError struct {
    Message    string            `json:"message"`
    Code       string            `json:"code"`
    TenantID   string            `json:"tenantId,omitempty"`
    Extensions map[string]interface{} `json:"extensions,omitempty"`
}
```

### Performance Optimization

#### Query Complexity Analysis

```go
// Prevent expensive queries
func ComplexityAnalyzer() graphql.HandlerExtension {
    return handler.NewRequestResponse(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
        // Analyze query complexity
        // Reject if too complex
        return next(ctx)
    })
}
```

#### Caching Strategy

```go
// Redis-based caching for frequent queries
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
    cacheKey := fmt.Sprintf("users:%s", getTenantID(ctx))
    
    // Try cache first
    if cached := r.cache.Get(cacheKey); cached != nil {
        return cached.([]*User), nil
    }
    
    // Query database and cache result
    users, err := r.userService.GetAll(ctx)
    if err == nil {
        r.cache.Set(cacheKey, users, 5*time.Minute)
    }
    
    return users, err
}
```

### API Documentation

#### Interactive GraphQL Playground

- **Development**: Available at `/playground`
- **Schema Introspection**: Built-in documentation
- **Query Testing**: Interactive query building

#### OpenAPI for REST Endpoints

```yaml
# OpenAPI 3.0 specification for REST endpoints
openapi: 3.0.0
info:
  title: Zplus SaaS REST API
  version: 1.0.0
paths:
  /api/v1/customers:
    get:
      summary: List customers
      parameters:
        - name: X-Tenant-ID
          in: header
          required: true
          schema:
            type: string
```

### Testing Strategy

#### GraphQL Testing

```go
func TestCustomerQuery(t *testing.T) {
    resolver := NewResolver()
    
    query := `
        query {
            customers {
                id
                name
                email
            }
        }
    `
    
    result := graphql.Do(graphql.Params{
        Schema:        schema,
        RequestString: query,
        Context:       createTestContext(),
    })
    
    assert.NoError(t, result.Errors)
    assert.NotNil(t, result.Data)
}
```

This API design strategy ensures efficient frontend-backend collaboration while maintaining the flexibility and power of GraphQL with the simplicity of REST where appropriate.

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