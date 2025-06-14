# Backend Services

This directory contains all backend microservices for the Zplus SaaS platform.

## Services

### Gateway Service (Port: 8080)
- Entry point for all API requests
- Handles authentication and tenant routing
- GraphQL and REST API endpoints

### Auth Service (Port: 8081)
- User authentication and authorization
- JWT token management
- Role-based access control (RBAC)

### File Service (Port: 8082)
- File upload and download management
- Multi-tenant file storage
- File metadata and indexing

### Payment Service (Port: 8083)
- Subscription management
- Billing and invoicing
- Payment gateway integration

### CRM Service (Port: 8084)
- Customer relationship management
- Contact and lead management
- Sales pipeline tracking

### HRM Service (Port: 8085)
- Human resource management
- Employee management
- Attendance and payroll

### POS Service (Port: 8086)
- Point of sale operations
- Product and inventory management
- Transaction processing

### Shared Library
- Common utilities and configurations
- Shared models and interfaces
- Database connection helpers

## Development

Each service is a standalone Go application with its own `go.mod` file.

```bash
# Run a specific service
cd apps/backend/gateway
go run main.go

# Run all services (requires multiple terminals)
./scripts/run-all-services.sh
```

## Architecture

All services follow the same structure:
- `main.go` - Entry point
- `handlers/` - HTTP handlers
- `models/` - Data models
- `services/` - Business logic
- `migrations/` - Database migrations