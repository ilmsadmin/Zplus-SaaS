# Zplus SaaS - Multi-Tenant Platform

Nền tảng SaaS đa tenant hiện đại với kiến trúc 3 tầng phân quyền rõ ràng (System → Tenant → Customer), hỗ trợ tích hợp nhiều module như CRM, LMS, POS, HRM...

## 🚀 Tổng quan dự án

Zplus SaaS là một nền tảng Software as a Service được xây dựng với công nghệ hiện đại, cho phép quản lý nhiều khách hàng (tenant) với dữ liệu hoàn toàn tách biệt. Hệ thống được thiết kế với khả năng mở rộng cao và hỗ trợ module-based architecture.

### ✨ Tính năng chính

- **🏗️ Kiến trúc 3 tầng**: System → Tenant → Customer với phân quyền rõ ràng
- **🔧 Module linh hoạt**: CRM, LMS, POS, HRM có thể bật/tắt theo nhu cầu
- **🔒 Bảo mật cao**: Multi-level RBAC và JWT authentication
- **⚡ Hiệu suất tối ưu**: Redis caching và microservices architecture
- **🌐 Multi-domain**: Hỗ trợ custom domain/subdomain cho từng tenant

## 🛠️ Công nghệ sử dụng

- **Frontend**: Next.js
- **Backend**: Go Fiber + GORM
- **Database**: PostgreSQL + MongoDB
- **Cache**: Redis
- **Architecture**: Microservices với GraphQL Gateway

## 📚 Tài liệu dự án

### Tài liệu thiết kế (Vietnamese)

- [📋 Thiết kế tổng quan dự án](./docs/thiet-ke-tong-quan-du-an.md) - Mục tiêu, phạm vi, đối tượng sử dụng
- [🏗️ Thiết kế kiến trúc dự án](./docs/thiet-ke-kien-truc-du-an.md) - Kiến trúc hệ thống, công nghệ, deployment
- [🗄️ Thiết kế kiến trúc database](./docs/thiet-ke-kien-truc-database.md) - Multi-tenant database design
- [🎨 Thiết kế UX/UI](./docs/thiet-ke-ux-ui.md) - Design system, components, accessibility
- [📖 Architecture Documentation](./docs/architecture.md) - Technical architecture overview

## 🎨 HTML Mockups

Các mockup HTML thể hiện kiến trúc 3 tầng của hệ thống:

- [🏠 **Mockup Gallery**](./mock/index.html) - Trang tổng quan các mockup
- [⚙️ **System Admin Dashboard**](./mock/system-admin-dashboard.html) - Quản lý global tenant, plans, modules
- [🏢 **Tenant Admin Dashboard**](./mock/tenant-admin-dashboard.html) - Quản lý tổ chức, team, cấu hình
- [📊 **CRM Dashboard**](./mock/customer-crm-dashboard.html) - Sales pipeline, quản lý khách hàng
- [🎓 **LMS Student Portal**](./mock/customer-lms-portal.html) - Học tập, certificates, progress tracking

## 🏗️ Cấu trúc dự án

```
zplus-saas/
│
├── apps/
│   ├── backend/
│   │   ├── gateway/            # GraphQL/REST gateway + auth/tenant
│   │   ├── auth/               # Authentication + RBAC
│   │   ├── file/               # File management
│   │   ├── payment/            # Transactions & subscriptions
│   │   ├── crm/                # Customer management
│   │   ├── hrm/                # HR management
│   │   ├── pos/                # Point of sale
│   │   └── shared/             # Shared libraries
│   │
│   ├── frontend/
│   │   ├── web/                # Unified Next.js app with multi-tenant routing
│   │   │   ├── app/
│   │   │   │   ├── admin/      # System admin (/admin)
│   │   │   │   ├── [tenant-slug]/ # Alternative tenant routing
│   │   │   │   └── tenant/[slug]/ # Main tenant routing & admin
│   │   │   └── middleware.ts   # Subdomain routing logic
│   │   ├── admin/              # Legacy tenant admin (deprecated)
│   │   └── ui/                 # Shared UI components
│
├── pkg/                        # Reusable SDKs & libraries
│
├── infra/
│   ├── db/                     # Database migrations
│   ├── k8s/                    # Kubernetes manifests
│   ├── docker/                 # Dockerfiles
│   └── ci-cd/                  # CI/CD configurations
│
├── docs/                       # Project documentation
└── mock/                       # HTML mockups
```

## 🚦 Multi-Tenant Routing Structure

The frontend now supports all required routing patterns:

1. **System Administration**: `/admin` - Global system management
2. **Tenant Admin**: 
   - `/tenant/[slug]/admin` - Organization administration  
   - `tenant-slug.domain.com/admin` - Subdomain-based admin (via middleware)
3. **System Homepage**: `/` - Main landing page
4. **Tenant Customer Pages**: 
   - `/tenant/[slug]` - Tenant service portal
   - `tenant-slug.domain.com` - Subdomain-based customer portal (via middleware)
   - `/[tenant-slug]` - Alternative direct routing (redirects to `/tenant/[slug]`)

## 🎯 Kiến trúc 3 tầng

```
┌─────────────────────────────┐
│        System Layer         │ ← Global management (tenants, plans, modules)
└─────────────────────────────┘
              │
              ▼
┌─────────────────────────────┐
│        Tenant Layer         │ ← Organization management (users, RBAC, config)
└─────────────────────────────┘
              │
              ▼
┌─────────────────────────────┐
│       Customer Layer        │ ← End-user interfaces (CRM, LMS, POS, HRM)
└─────────────────────────────┘
```

## 🚦 Request Flow

1. **Subdomain routing**: `tenant.myapp.com` → `X-Tenant-ID: tenant`
2. **Gateway processing**: Authentication + tenant validation
3. **Module routing**: Route to appropriate microservice
4. **Database isolation**: Access tenant-specific schema

## 🔧 Development

```bash
# Backend (Go Fiber)
cd apps/backend
go mod tidy
go run main.go

# Frontend (Next.js)
cd apps/frontend/web
npm install
npm run dev
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Zplus SaaS** - Powering the future of multi-tenant applications
