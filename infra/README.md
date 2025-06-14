# Infrastructure

This directory contains all infrastructure-related code and configurations for Zplus SaaS multi-tenant platform.

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
- Docker Compose configurations with Traefik setup
- Multi-stage build configurations
- Traefik configuration files for multi-tenant routing

### CI/CD (`ci-cd/`)
- GitHub Actions workflows
- Build and deployment scripts
- Environment-specific configurations

## Traefik Multi-tenant Setup

### Key Features
- **Subdomain-based tenant routing**: `tenant-slug.zplus.com` → `X-Tenant-ID: tenant-slug`
- **SSL/TLS termination** with Let's Encrypt certificates
- **Load balancing** across multiple backend instances
- **Rate limiting** per tenant
- **Security headers** injection
- **Access logging** with tenant context

### Configuration Files
- `docker/traefik.yml` - Main Traefik configuration
- `docker/dynamic.yml` - Dynamic routing rules and middlewares
- `docker/docker-compose.yml` - Complete stack with Traefik integration

## Usage

### Local Development
```bash
# Start all services including databases and Traefik
cd infra/docker
docker-compose up -d

# View Traefik dashboard
open http://localhost:8080

# Test tenant routing (requires /etc/hosts entries)
echo "127.0.0.1 tenant1.localhost" >> /etc/hosts
curl -H "Host: tenant1.localhost" http://localhost/api/health

# Apply database migrations
psql -h localhost -U zplus_user -d zplus_saas -f ../db/001_system_schema.sql
```

### Production Deployment with Kubernetes
```bash
# Apply all manifests including Traefik ingress
kubectl apply -f infra/k8s/

# Check deployment status
kubectl get pods -n zplus-saas

# View Traefik logs
kubectl logs -f deployment/traefik -n zplus-saas
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

### Traefik-specific Variables
- `TRAEFIK_ACME_EMAIL` - Email for Let's Encrypt certificates
- `TRAEFIK_DOMAIN` - Base domain for the platform (e.g., zplus.com)

## Message Queue Strategy

### Why Redis over Kafka?

| Aspect | Redis Queue | Kafka | Decision |
|--------|-------------|-------|----------|
| **Setup Complexity** | Simple | Complex | ✅ Redis for MVP |
| **Latency** | < 1ms | 5-10ms | ✅ Redis for real-time |
| **Throughput** | 100K msg/s | 1M+ msg/s | Redis sufficient for current scale |
| **Use Case Fit** | Multi-tenant SaaS | Event streaming | ✅ Perfect for our needs |

### Current Redis Use Cases
- **Background Jobs**: Email notifications, PDF generation, data exports
- **Real-time Updates**: Live chat, dashboard updates, notifications  
- **Session Storage**: User sessions and authentication tokens
- **Caching**: Frequently accessed tenant data

### When to Consider Kafka Migration
- Scale exceeds 100K messages/second
- Need for event sourcing and replay capabilities
- Cross-system event streaming requirements
- Advanced analytics and data pipeline needs

## Security Considerations

### Traefik Security Features
- **SSL/TLS everywhere** with automatic certificate renewal
- **Rate limiting** per tenant to prevent abuse
- **Security headers** injection (HSTS, CSP, XSS protection)
- **Access logging** with tenant context for audit trails

### Network Security
- All services communicate over private Docker networks
- Database ports not exposed to public internet
- Traefik handles all external traffic and SSL termination

## Monitoring Integration

### Traefik Metrics
- Prometheus metrics endpoint exposed
- Request rate, response time, and error rate tracking
- Per-tenant traffic analysis
- SSL certificate expiration monitoring

### Log Aggregation
- Structured JSON logs with tenant context
- Centralized logging with ELK stack integration
- Access logs for security auditing
- Error tracking and alerting