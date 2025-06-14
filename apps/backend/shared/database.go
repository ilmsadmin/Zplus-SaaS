package shared

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase initializes the database connection and runs migrations
func InitDatabase() error {
	// Get database configuration from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "zplus_saas")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// Create DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, dbname, port, sslmode)

	// Connect to database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto-migrate the schema
	err = DB.AutoMigrate(
		&User{},
		&Tenant{},
		&Role{},
		&Permission{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	// Seed initial data
	if err := seedInitialData(); err != nil {
		log.Printf("Warning: failed to seed initial data: %v", err)
	}

	log.Println("Database initialized successfully")
	return nil
}

// seedInitialData creates initial roles and permissions
func seedInitialData() error {
	// Create system-level roles
	systemRoles := []Role{
		{Name: "Super Admin", Level: "system", Description: "Full system access"},
		{Name: "System Admin", Level: "system", Description: "System administration"},
		{Name: "Support Staff", Level: "system", Description: "Customer support"},
	}

	// Create tenant-level roles
	tenantRoles := []Role{
		{Name: "Tenant Admin", Level: "tenant", Description: "Tenant administration"},
		{Name: "Manager", Level: "tenant", Description: "Team management"},
		{Name: "User", Level: "tenant", Description: "Regular user"},
		{Name: "Viewer", Level: "tenant", Description: "Read-only access"},
	}

	// Create customer-level roles
	customerRoles := []Role{
		{Name: "Student", Level: "customer", Description: "LMS student"},
		{Name: "Teacher", Level: "customer", Description: "LMS teacher"},
		{Name: "Salesperson", Level: "customer", Description: "CRM salesperson"},
		{Name: "Employee", Level: "customer", Description: "HRM employee"},
		{Name: "Cashier", Level: "customer", Description: "POS cashier"},
	}

	// Combine all roles
	allRoles := append(systemRoles, append(tenantRoles, customerRoles...)...)

	// Create roles if they don't exist
	for _, role := range allRoles {
		var existingRole Role
		if err := DB.Where("name = ? AND level = ?", role.Name, role.Level).First(&existingRole).Error; err == gorm.ErrRecordNotFound {
			if err := DB.Create(&role).Error; err != nil {
				return fmt.Errorf("failed to create role %s: %v", role.Name, err)
			}
		}
	}

	// Create basic permissions
	permissions := []Permission{
		// System permissions
		{Name: "system.manage", Resource: "system", Action: "manage", Description: "Manage system settings"},
		{Name: "tenant.manage", Resource: "tenant", Action: "manage", Description: "Manage tenants"},
		{Name: "user.manage", Resource: "user", Action: "manage", Description: "Manage users"},
		
		// Tenant permissions
		{Name: "tenant.view", Resource: "tenant", Action: "view", Description: "View tenant data"},
		{Name: "user.create", Resource: "user", Action: "create", Description: "Create users"},
		{Name: "user.update", Resource: "user", Action: "update", Description: "Update users"},
		{Name: "user.delete", Resource: "user", Action: "delete", Description: "Delete users"},
		
		// Module permissions
		{Name: "crm.access", Resource: "crm", Action: "access", Description: "Access CRM module"},
		{Name: "lms.access", Resource: "lms", Action: "access", Description: "Access LMS module"},
		{Name: "pos.access", Resource: "pos", Action: "access", Description: "Access POS module"},
		{Name: "hrm.access", Resource: "hrm", Action: "access", Description: "Access HRM module"},
	}

	// Create permissions if they don't exist
	for _, permission := range permissions {
		var existingPermission Permission
		if err := DB.Where("name = ?", permission.Name).First(&existingPermission).Error; err == gorm.ErrRecordNotFound {
			if err := DB.Create(&permission).Error; err != nil {
				return fmt.Errorf("failed to create permission %s: %v", permission.Name, err)
			}
		}
	}

	// Create default system admin user if not exists
	var adminUser User
	if err := DB.Where("email = ?", "admin@zplus.com").First(&adminUser).Error; err == gorm.ErrRecordNotFound {
		// Get Super Admin role
		var superAdminRole Role
		if err := DB.Where("name = ? AND level = ?", "Super Admin", "system").First(&superAdminRole).Error; err != nil {
			return fmt.Errorf("failed to find Super Admin role: %v", err)
		}

		// Hash password properly - this is just for demo, in production use bcrypt
		hashedPassword := "$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPmOXz7w8MRoW" // "admin123"

		adminUser = User{
			Email:     "admin@zplus.com",
			Password:  hashedPassword,
			FirstName: "System",
			LastName:  "Administrator",
			IsActive:  true,
			RoleID:    superAdminRole.ID,
		}

		if err := DB.Create(&adminUser).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %v", err)
		}
	}

	// Create demo tenant if not exists
	var demoTenant Tenant
	if err := DB.Where("slug = ?", "demo").First(&demoTenant).Error; err == gorm.ErrRecordNotFound {
		demoTenant = Tenant{
			Name:       "Demo Organization",
			Slug:       "demo",
			Domain:     "demo.zplus.com",
			IsActive:   true,
			SchemaName: "tenant_demo",
		}

		if err := DB.Create(&demoTenant).Error; err != nil {
			return fmt.Errorf("failed to create demo tenant: %v", err)
		}
	}

	// Create demo tenant admin user if not exists
	var tenantAdminUser User
	if err := DB.Where("email = ?", "tenant@demo.com").First(&tenantAdminUser).Error; err == gorm.ErrRecordNotFound {
		// Get Tenant Admin role
		var tenantAdminRole Role
		if err := DB.Where("name = ? AND level = ?", "Tenant Admin", "tenant").First(&tenantAdminRole).Error; err != nil {
			return fmt.Errorf("failed to find Tenant Admin role: %v", err)
		}

		hashedPassword := "$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPmOXz7w8MRoW" // "admin123"

		tenantAdminUser = User{
			Email:     "tenant@demo.com",
			Password:  hashedPassword,
			FirstName: "Tenant",
			LastName:  "Admin",
			IsActive:  true,
			TenantID:  &demoTenant.ID,
			RoleID:    tenantAdminRole.ID,
		}

		if err := DB.Create(&tenantAdminUser).Error; err != nil {
			return fmt.Errorf("failed to create tenant admin user: %v", err)
		}
	}

	// Create demo customer user if not exists
	var customerUser User
	if err := DB.Where("email = ?", "user@demo.com").First(&customerUser).Error; err == gorm.ErrRecordNotFound {
		// Get Customer role (Student)
		var customerRole Role
		if err := DB.Where("name = ? AND level = ?", "Student", "customer").First(&customerRole).Error; err != nil {
			return fmt.Errorf("failed to find Student role: %v", err)
		}

		hashedPassword := "$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPmOXz7w8MRoW" // "admin123"

		customerUser = User{
			Email:     "user@demo.com",
			Password:  hashedPassword,
			FirstName: "Demo",
			LastName:  "User",
			IsActive:  true,
			TenantID:  &demoTenant.ID,
			RoleID:    customerRole.ID,
		}

		if err := DB.Create(&customerUser).Error; err != nil {
			return fmt.Errorf("failed to create customer user: %v", err)
		}
	}

	return nil
}

// getEnv gets environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// SetTenantContext sets the tenant context for database operations
func SetTenantContext(db *gorm.DB, tenantID uint) *gorm.DB {
	return db.Where("tenant_id = ? OR tenant_id IS NULL", tenantID)
}