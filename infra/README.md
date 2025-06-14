# Infrastructure

This directory contains all infrastructure-related code and configurations.

## Structure

### Database (`db/`)
- SQL migration files
- Schema definitions for system and tenant databases
- Seed data and initial configurations

### Kubernetes (`k8s/`)
- Kubernetes deployment manifests
- Service definitions
- ConfigMaps and Secrets
- Persistent Volume Claims

### Docker (`docker/`)
- Dockerfiles for different services
- Docker Compose configurations
- Multi-stage build configurations

### CI/CD (`ci-cd/`)
- GitHub Actions workflows
- Build and deployment scripts
- Environment-specific configurations

## Usage

### Local Development
```bash
# Start all services with Docker Compose
cd infra/docker
docker-compose up -d

# Apply database migrations
psql -h localhost -U zplus_user -d zplus_saas -f ../db/001_system_schema.sql
```

### Kubernetes Deployment
```bash
# Apply all manifests
kubectl apply -f infra/k8s/

# Check deployment status
kubectl get pods -n zplus-saas
```

### Database Migrations
- `001_system_schema.sql` - System-level tables (tenants, plans, modules)
- `002_tenant_schema_template.sql` - Template for tenant-specific schemas

## Environment Variables

Required environment variables for services:
- `DB_HOST` - Database host
- `DB_USER` - Database username  
- `DB_PASSWORD` - Database password
- `REDIS_HOST` - Redis host
- `JWT_SECRET` - JWT signing secret