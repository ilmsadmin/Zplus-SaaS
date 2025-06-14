# Thiết kế Kiến trúc Dự án - Zplus SaaS

## 1. Tổng quan Kiến trúc

Zplus SaaS được xây dựng trên kiến trúc **3-tier Multi-tenant** với sự phân tách rõ ràng giữa các tầng:

```
+-------------------------+
|        System          |  ← Quản trị toàn cục (RBAC, gói dịch vụ, tenant, domain)
+-------------------------+
            |
            v
+-------------------------+
|        Tenant          |  ← Quản trị trong phạm vi tenant (RBAC, user, module, khách hàng)
+-------------------------+
            |
            v
+-------------------------+
|       Customer         |  ← Người dùng cuối, sử dụng dịch vụ (CRM, LMS, POS...)
+-------------------------+
```

## 2. Kiến trúc Microservices

### 2.1 Sơ đồ Tổng quan

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│    Frontend     │    │   API Gateway   │    │   Load Balancer │
│  (React/Vue)    │◄──►│   (GraphQL)     │◄──►│    (Traefik)    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                ┌───────────────┼───────────────┐
                │               │               │
        ┌───────▼───────┐ ┌─────▼─────┐ ┌──────▼──────┐
        │ System Service│ │Tenant Svc │ │Customer Svc │
        │   (Go/Fiber)  │ │(Go/Fiber) │ │ (Go/Fiber)  │
        └───────────────┘ └───────────┘ └─────────────┘
                │               │               │
        ┌───────▼───────┐ ┌─────▼─────┐ ┌──────▼──────┐
        │  PostgreSQL   │ │PostgreSQL │ │  PostgreSQL │
        │ (system db)   │ │(tenant_*)  │ │(module dbs) │
        └───────────────┘ └───────────┘ └─────────────┘
```

### 2.2 Technology Stack

| Component | Technology | Purpose |
|-----------|------------|---------|
| **Frontend** | React/Vue.js | User Interface |
| **API Gateway** | GraphQL (Go) | API Orchestration |
| **Load Balancer** | Traefik | Traffic Routing |
| **Backend Services** | Go + Fiber | Business Logic |
| **Database** | PostgreSQL | Data Persistence |
| **Cache** | Redis | Session & Cache |
| **Message Queue** | Redis/RabbitMQ | Async Processing |
| **Auth** | JWT + RBAC | Authentication |

## 3. Luồng Request

### 3.1 Request Flow Diagram

```
[Client] → [Traefik] → [GraphQL Gateway] → [Microservice] → [Database]
   │           │              │                  │              │
   │           │              │                  │              │
   │        Domain/         Middleware:       Business        Tenant
   │       Subdomain        - Auth             Logic          Schema
   │       Detection        - RBAC                            Access
   │                       - Rate Limit
   │                       - Tenant ID
```

### 3.2 Luồng xử lý chi tiết

1. **Client Request**: `student-zin100.myapp.com/api/graphql`

2. **Traefik Processing**:
   - Phân tích subdomain: `zin100`
   - Thêm header: `X-Tenant-ID: zin100` 
   - Route đến GraphQL Gateway

3. **GraphQL Gateway**:
   - Middleware authentication (JWT)
   - Middleware authorization (RBAC)
   - Middleware tenant validation
   - Middleware rate limiting
   - Route đến microservice tương ứng

4. **Microservice Processing**:
   - Lấy tenant context: `GetTenantDB(c)`  
   - Xử lý business logic
   - Truy cập database schema tương ứng
   - Trả về kết quả

## 4. Database Architecture

### 4.1 Multi-tenant Strategy: Schema per Tenant

```
┌─────────────────────────────────────────────────────────────┐
│                      PostgreSQL Server                      │
├─────────────────────────────────────────────────────────────┤
│  Schema: public/system                                      │
│  ├── system_users       (System Admins)                    │
│  ├── tenants           (Tenant Registry)                   │
│  ├── plans             (Service Plans)                     │
│  ├── subscriptions     (Tenant Subscriptions)             │
│  ├── modules           (Available Modules)                 │
│  └── tenant_modules    (Enabled Modules per Tenant)       │
├─────────────────────────────────────────────────────────────┤
│  Schema: tenant_zin100                                      │
│  ├── users             (Tenant Users)                      │
│  ├── roles             (Tenant Roles)                      │
│  ├── permissions       (Tenant Permissions)                │
│  ├── customers         (End Customers)                     │
│  └── modules_config    (Module Configurations)             │
├─────────────────────────────────────────────────────────────┤
│  Schema: tenant_acme                                        │
│  ├── users, roles, permissions...                          │
│  └── (Same structure as above)                             │
└─────────────────────────────────────────────────────────────┘
```

### 4.2 Module-specific Tables

Mỗi module sẽ có bảng riêng trong schema của tenant:

**LMS Module**:
```sql
-- Schema: tenant_xyz
CREATE TABLE students (id, name, email, ...);
CREATE TABLE courses (id, title, description, ...);
CREATE TABLE enrollments (student_id, course_id, ...);
CREATE TABLE lessons (course_id, title, content, ...);
CREATE TABLE quizzes (course_id, questions, ...);
```

## 5. Security Architecture

### 5.1 Multi-layer Security

```
┌─────────────────────────────────────────────┐
│              Security Layers                │
├─────────────────────────────────────────────┤
│ 1. Network Security (SSL/TLS, Firewall)    │
├─────────────────────────────────────────────┤
│ 2. Application Security (JWT, Rate Limit)  │
├─────────────────────────────────────────────┤
│ 3. Authorization (RBAC Multi-tier)         │
├─────────────────────────────────────────────┤
│ 4. Data Security (Encryption, Audit)       │
├─────────────────────────────────────────────┤
│ 5. Tenant Isolation (Schema Separation)    │
└─────────────────────────────────────────────┘
```

### 5.2 RBAC Model

**System Level**:
- Super Admin
- System Admin  
- Support Staff

**Tenant Level**:
- Tenant Admin
- Manager
- User
- Viewer

**Customer Level**:
- Module-specific roles (Student, Teacher, Salesperson...)

## 6. Scalability & Performance

### 6.1 Horizontal Scaling

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  App Instance 1 │    │  App Instance 2 │    │  App Instance N │
│   (Container)   │    │   (Container)   │    │   (Container)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │   Load Balancer │
                    │    (Traefik)    │
                    └─────────────────┘
```

### 6.2 Database Scaling

- **Read Replicas**: Cho read-heavy operations
- **Connection Pooling**: Optimize database connections  
- **Redis Caching**: Cache frequently accessed data
- **Database Sharding**: Phân chia tenant theo regions

### 6.3 Performance Optimization

- **CDN**: Static assets delivery
- **Lazy Loading**: Module loading on demand
- **Database Indexing**: Optimized queries
- **Background Jobs**: Async processing với Redis Queue

## 7. Module System Architecture

### 7.1 Module Registry

```go
type BaseModule interface {
    GetName() string
    GetVersion() string
    GetDependencies() []string
    Install(tenantID string) error
    Uninstall(tenantID string) error
    GetRoutes() []Route
    GetPermissions() []Permission
}
```

### 7.2 Module Lifecycle

1. **Module Registration**: Đăng ký module trong system
2. **Tenant Activation**: Kích hoạt module cho tenant
3. **Schema Migration**: Tạo tables cần thiết
4. **Route Registration**: Đăng ký API routes
5. **Permission Setup**: Thiết lập permissions

## 8. Deployment Architecture

### 8.1 Container Orchestration

```yaml
# docker-compose.yml
version: '3.8'
services:
  traefik:
    image: traefik:v2.10
    ports: ["80:80", "443:443"]
  
  app:
    build: .
    deploy:
      replicas: 3
    environment:
      - DATABASE_URL=${DATABASE_URL}
      - REDIS_URL=${REDIS_URL}
  
  postgres:
    image: postgres:15
    environment:
      - POSTGRES_DB=zplus_saas
  
  redis:
    image: redis:7-alpine
```

### 8.2 Environment Architecture

- **Development**: Single instance, local database
- **Staging**: Multi-instance, shared database  
- **Production**: Clustered, replicated database

## 9. Monitoring & Observability

### 9.1 Monitoring Stack

- **Metrics**: Prometheus + Grafana
- **Logging**: ELK Stack (Elasticsearch, Logstash, Kibana)
- **Tracing**: Jaeger
- **Health Checks**: Built-in endpoints
- **Alerting**: Grafana Alerts

### 9.2 Key Metrics

- **System**: CPU, Memory, Disk usage
- **Application**: Response time, Error rate, Throughput
- **Business**: Tenant count, Active users, Revenue