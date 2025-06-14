1. Tổng quan kiến trúc 🚀
Kiến trúc tổng thể vẫn tuân theo mô hình Microservices trên Kubernetes, nhưng các thành phần được làm rõ hơn:
Cổng vào (Ingress): Traefik sẽ đóng vai trò là API Gateway, tự động phát hiện các dịch vụ và định tuyến request. Nó sẽ xác định tenant dựa trên subdomain và gắn TenantID vào header.
Tầng API: Một GraphQL Gateway (viết bằng Go) nhận request từ Traefik, xác thực và điều phối đến các microservice nghiệp vụ.
Tầng nghiệp vụ: Các microservice được xây dựng bằng Go với framework Fiber cho hiệu năng cao. Logic truy vấn database sẽ được xử lý qua GORM.
Tầng dữ liệu:
PostgreSQL: 1 database duy nhất, chứa 1000+ schema (ví dụ: tenant_acme, tenant_globex).
MongoDB: Mỗi tenant có một database riêng (acme_docs, globex_docs).
Redis: Dùng chung, phân tách dữ liệu bằng tiền tố (prefix) trên key.
CI/CD: GitHub Actions sẽ build, test và deploy các container lên Kubernetes.
2. Chi tiết triển khai với Fiber và GORM
Đây là phần quan trọng nhất, mô tả cách các công nghệ cụ thể này phối hợp với nhau.
a. Traefik: Cổng vào thông minh
Traefik sẽ được cấu hình làm Ingress Controller trong Kubernetes.
Luồng hoạt động:
Request đến tenant-acme.myapp.com.
Traefik (thông qua IngressRoute CRD) bắt được request này.
Nó sử dụng Middleware để trích xuất acme từ host và tạo một header mới: X-Tenant-ID: acme.
Request được chuyển tiếp đến dịch vụ graphql-gateway.
Lợi ích: Logic xác định tenant được xử lý hoàn toàn ở tầng biên, các service bên trong không cần quan tâm đến subdomain.
b. Microservices với Go Fiber & GORM
Fiber là một lựa chọn tuyệt vời vì nó cực kỳ nhanh và có API tương tự Express.js, dễ sử dụng. GORM là một ORM mạnh mẽ cho Go.
Thách thức chính: Làm thế nào để mọi câu lệnh GORM của một request đều thực thi trên đúng schema của tenant đó?
Giải pháp: Sử dụng một Fiber Middleware kết hợp với một hàm khởi tạo GORM session.
Fiber Middleware để lấy TenantID: Tạo một middleware được thực thi đầu tiên cho mọi request cần truy cập database.
// file: middlewares/tenant.go
func TenantResolver() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Lấy TenantID từ header do Traefik gắn vào
        tenantID := c.Get("X-Tenant-ID")
        if tenantID == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tenant ID is missing"})
        }

        // Lưu tenantID vào context của request để các hàm sau có thể dùng
        c.Locals("tenant_id", tenantID)
        return c.Next()
    }
}


Hàm khởi tạo GORM Session cho Tenant: Tạo một hàm helper để lấy DB session đã được cấu hình đúng search_path.
// file: database/connection.go
var DB *gorm.DB // Biến global chứa kết nối gốc

func InitDatabase() {
    // Khởi tạo kết nối gốc đến PostgreSQL
    dsn := "host=... user=... password=... dbname=saas_app port=5432 sslmode=disable"
    DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Hàm quan trọng nhất!
func GetTenantDB(c *fiber.Ctx) *gorm.DB {
    // Lấy tenantID từ context mà middleware đã lưu
    tenantID, ok := c.Locals("tenant_id").(string)
    if !ok || tenantID == "" {
        return nil // Hoặc panic, tùy vào logic của bạn
    }

    // Tạo một session mới và set search_path cho tenant này
    // Mọi thao tác trên dbSession này sẽ chỉ áp dụng cho schema của tenant
    dbSession := DB.Session(&gorm.Session{})
    dbSession.Exec("SET search_path TO ?", tenantID)
    return dbSession
}


Sử dụng trong một Handler của Fiber: Bây giờ, trong các hàm xử lý request, bạn chỉ cần gọi GetTenantDB để có được kết nối an toàn.
// file: handlers/product_handler.go
func GetProducts(c *fiber.Ctx) error {
    // Lấy DB session đã được scope cho đúng tenant
    db := database.GetTenantDB(c)
    if db == nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not get database session"})
    }

    var products []models.Product
    // GORM sẽ tự động chạy câu lệnh "SELECT * FROM products"
    // trên schema của tenant hiện tại vì đã có `SET search_path`
    db.Find(&products)

    return c.Status(fiber.StatusOK).JSON(products)
}

// Đăng ký route và middleware
// app.Use(middlewares.TenantResolver())
// app.Get("/products", handlers.GetProducts)


c. Quản lý Schema Migration với GORM
Với 1000 tenant, bạn cần một công cụ để tự động cập nhật schema cho tất cả.
Dịch vụ quản lý Tenant: Tạo một microservice riêng (hoặc một lệnh CLI) cho việc này.
Logic migration:
Service này kết nối đến database saas_app.
Nó truy vấn để lấy danh sách tất cả tenantId đang có.
Vòng lặp qua từng tenantId:
Thực thi SET search_path TO ?, tenantId
Chạy db.AutoMigrate(&models.Product{}, &models.Order{}, ...)
GORM sẽ tự động so sánh các struct model của bạn với các bảng trong schema của tenant đó và thêm/thay đổi cột nếu cần.
3. Các thành phần khác và luồng hoạt động
Next.js (Frontend): Hoạt động như đã mô tả, gửi request đến GraphQL Gateway.
GraphQL Gateway: Nhận request, gọi middleware TenantResolver để xác định tenant, sau đó điều phối đến các Fiber microservice tương ứng.
Kafka (Xử lý bất đồng bộ): Khi một Fiber service cần gửi một tác vụ nền, nó sẽ đẩy một message vào Kafka. Message này bắt buộc phải chứa tenantId trong payload để các consumer biết phải thao tác trên schema nào.
GitHub Actions (CI/CD): Quy trình không đổi: push code -> build & test -> build docker image -> push to registry -> kubectl apply.
4. Thách thức và giải pháp cho 1000 Tenant
Các thách thức vẫn tương tự, nhưng giải pháp có thể được tinh chỉnh:
Quản lý kết nối Database: Rất quan trọng khi dùng GORM. Đảm bảo bạn không mở kết nối mới cho mỗi request. Sử dụng connection pool có sẵn của GORM và đặt PgBouncer ở giữa để quản lý hàng nghìn kết nối từ các pod Kubernetes.
Migration Schema hàng loạt: Kịch bản migration bằng GORM AutoMigrate như mô tả ở trên là giải pháp trực tiếp. Cần có cơ chế ghi log và xử lý lỗi cẩn thận.
Giám sát (Monitoring): Sử dụng một thư viện Prometheus client cho Go. Trong middleware TenantResolver, sau khi lấy được tenantID, hãy tăng một bộ đếm (counter) với label tenant_id. Ví dụ: http_requests_total{tenant="acme", method="GET", path="/products"}.
Backup và Restore: Kịch bản pg_dump --schema=<tenant_id> vẫn là giải pháp hiệu quả nhất để sao lưu cho từng tenant.
Bằng cách áp dụng các mẫu thiết kế này, bạn có thể tận dụng sức mạnh của Fiber và GORM để xây dựng một hệ thống SaaS multi-tenant hiệu năng cao, an toàn và dễ bảo trì trên nền tảng Kubernetes.
