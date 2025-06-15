package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TenantUser represents a user within a specific tenant
type TenantUser struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TenantID    uuid.UUID      `json:"tenant_id" gorm:"type:uuid;not null"`
	Email       string         `json:"email" gorm:"not null"`
	PasswordHash string        `json:"-" gorm:"not null"`
	FirstName   string         `json:"first_name" gorm:"not null"`
	LastName    string         `json:"last_name" gorm:"not null"`
	Avatar      *string        `json:"avatar"`
	Status      string         `json:"status" gorm:"default:'active'"` // active, inactive, suspended
	LastLoginAt *time.Time     `json:"last_login_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Roles       []Role         `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

// Role represents a role within a tenant
type Role struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TenantID    uuid.UUID      `json:"tenant_id" gorm:"type:uuid;not null"`
	Name        string         `json:"name" gorm:"not null"`
	Description *string        `json:"description"`
	IsSystemRole bool          `json:"is_system_role" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// Relationships
	Permissions []Permission   `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
	Users       []TenantUser   `json:"users,omitempty" gorm:"many2many:user_roles;"`
}

// Permission represents a permission that can be assigned to roles
type Permission struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Resource    string         `json:"resource" gorm:"not null"`
	Action      string         `json:"action" gorm:"not null"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	
	// Relationships
	Roles       []Role         `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
}

// UserRole represents the many-to-many relationship between users and roles
type UserRole struct {
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey"`
	RoleID    uuid.UUID `json:"role_id" gorm:"type:uuid;primaryKey"`
	AssignedAt time.Time `json:"assigned_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// RolePermission represents the many-to-many relationship between roles and permissions
type RolePermission struct {
	RoleID       uuid.UUID `json:"role_id" gorm:"type:uuid;primaryKey"`
	PermissionID uuid.UUID `json:"permission_id" gorm:"type:uuid;primaryKey"`
	AssignedAt   time.Time `json:"assigned_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// TableName methods for tenant-scoped tables
// Note: These will be dynamically set based on tenant context

func (TenantUser) TableName() string {
	return "users" // Will be prefixed with tenant schema
}

func (Role) TableName() string {
	return "roles"
}

func (Permission) TableName() string {
	return "permissions" // System-wide, not tenant-scoped
}

func (UserRole) TableName() string {
	return "user_roles"
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

// Helper methods for TenantUser
func (u *TenantUser) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}

func (u *TenantUser) GetPermissions() []string {
	var permissions []string
	permissionSet := make(map[string]bool)
	
	for _, role := range u.Roles {
		for _, permission := range role.Permissions {
			if !permissionSet[permission.Name] {
				permissions = append(permissions, permission.Name)
				permissionSet[permission.Name] = true
			}
		}
	}
	
	return permissions
}

func (u *TenantUser) HasPermission(permissionName string) bool {
	for _, role := range u.Roles {
		for _, permission := range role.Permissions {
			if permission.Name == permissionName {
				return true
			}
		}
	}
	return false
}

func (u *TenantUser) CanAccessResource(resource, action string) bool {
	permissionName := resource + ":" + action
	return u.HasPermission(permissionName)
}