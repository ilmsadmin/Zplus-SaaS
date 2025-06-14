Kiến trúc nền tảng SaaS đa tenant gồm 3 tầng phân quyền (System → Tenant → Customer), dễ mở rộng ra nhiều module trong tương lai như CRM, POS, LMS...


---

🏗️ Kiến trúc tổng quan

+-------------------------+
|        System          |  <-- Quản trị toàn cục (RBAC, gói dịch vụ, tenant, domain)
+-------------------------+
            |
            v
+-------------------------+
|        Tenant          |  <-- Quản trị trong phạm vi tenant (RBAC, user, module, khách hàng)
+-------------------------+
            |
            v
+-------------------------+
|       Customer         |  <-- Người dùng cuối, sử dụng dịch vụ (CRM, LMS, POS...)
+-------------------------+


---

📦 Tầng 1: System Layer

Dữ liệu chung toàn hệ thống (đặt trong schema public hoặc riêng schema system).

Các chức năng:

Quản trị người dùng hệ thống (admin panel).

Quản lý Tenant:

Tên + Mô tả

Subdomain / domain tùy chỉnh

Trạng thái (active, suspended)


RBAC cho admin system (nếu có nhiều người vận hành).

Quản lý Plan/Subscription:

Plan: tên, mô tả, giá, giới hạn (user, dung lượng...)

Subscription: tenant nào đang dùng plan nào, thời hạn


Quản lý các Modules có thể kích hoạt cho từng tenant:

CRM, POS, LMS, HRM, Checkin...

Module có thể bao gồm tên, mô tả, cấu hình bật/tắt


(Tùy chọn) Quản lý thanh toán (Stripe/Billing API...)


Tables gợi ý:

Table	Description

system_users	Admin hệ thống
tenants	Danh sách tenant (id, name, domain)
plans	Các gói cước
subscriptions	Gói đang dùng cho từng tenant
modules	Các module được hỗ trợ
tenant_modules	Các module được bật cho tenant



---

🏢 Tầng 2: Tenant Layer

Dữ liệu riêng trong từng schema PostgreSQL: tenant_acme, tenant_zin100...

Chức năng:

RBAC cho tenant: user, role, permission.

Quản lý người dùng nội bộ của tenant (admin, nhân viên...).

Quản lý khách hàng cuối (customers) tùy theo module đang dùng.

Theo dõi usage, cấu hình tenant.

Kích hoạt module nào sẽ hiển thị/ẩn tính năng tương ứng.

Cấu hình tích hợp (zalo OA, email, sms...) tùy tenant.


Tables cơ bản:

Table	Description

users	Người dùng nội bộ của tenant
roles	Vai trò
permissions	Quyền
user_roles	Gán người dùng vào vai trò
customers	Khách hàng cuối
modules_config	Bật/tắt các tính năng trong tenant


> Bạn có thể định nghĩa 1 BaseModule interface để các module mới thêm dễ dàng (CRM, POS,...).




---

👤 Tầng 3: Customer Layer

Người dùng cuối sử dụng dịch vụ, ví dụ:

Học viên (LMS)

Khách hàng CRM

Khách mua hàng (POS)

Nhân viên (HRM)


> Tùy vào module được bật mà dữ liệu và flow của tầng này sẽ khác.



Ví dụ với LMS: | Table          | Description                      | |----------------|----------------------------------| | students     | Học viên                         | | courses      | Khóa học                         | | enrollments  | Ghi danh                         | | lessons      | Nội dung                         | | quizzes      | Bài kiểm tra                     |


---

🔗 Module System - Gợi ý mở rộng

Hệ thống cần hỗ trợ module-based feature toggle:

Module registry: khai báo module

Tenant-specific config: schema tenant_zin100.modules_config

API Gateway GraphQL điều phối route theo module đang bật

Middleware kiểm tra module permission theo X-Tenant-ID



---

🧠 Kiến trúc kỹ thuật tóm tắt

Layer	Scope	Database Schema	Auth	Tool/Libs

System	Global	system / public	JWT + RBAC (admin)	Go + Fiber + GORM
Tenant	Per-tenant (isolated)	tenant_xyz	JWT + RBAC (tenant)	Fiber + GORM + Redis
Customer	End-user (module-specific)	Depends on module	Session/JWT (lightweight)	Fiber/module-specific logic



---

🚦 Luồng request điển hình

1. student-zin100.myapp.com/api/graphql


2. Traefik xác định subdomain: zin100 → Header: X-Tenant-ID: zin100


3. GraphQL Gateway:

Gọi middleware check tenant + module + RBAC

Route đến microservice Fiber tương ứng



4. Microservice dùng GetTenantDB(c) để truy cập schema tương ứng.
