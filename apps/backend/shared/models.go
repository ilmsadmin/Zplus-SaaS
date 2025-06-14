package shared

import (
	"time"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Email       string         `json:"email" gorm:"uniqueIndex;not null"`
	Password    string         `json:"-" gorm:"not null"` // Don't serialize password
	FirstName   string         `json:"first_name" gorm:"not null"`
	LastName    string         `json:"last_name" gorm:"not null"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	TenantID    *uint          `json:"tenant_id" gorm:"index"` // Nullable for system admins
	RoleID      uint           `json:"role_id" gorm:"not null;index"`
	Role        Role           `json:"role" gorm:"foreignKey:RoleID"`
	Tenant      *Tenant        `json:"tenant,omitempty" gorm:"foreignKey:TenantID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// Tenant represents a tenant/organization in the system
type Tenant struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Domain      string         `json:"domain" gorm:"uniqueIndex"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	SchemaName  string         `json:"schema_name" gorm:"not null"`
	PlanID      *uint          `json:"plan_id" gorm:"index"`
	Users       []User         `json:"-" gorm:"foreignKey:TenantID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// Role represents roles in the RBAC system
type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Level       string         `json:"level" gorm:"not null"` // "system", "tenant", "customer"
	Description string         `json:"description"`
	Permissions []Permission   `json:"permissions" gorm:"many2many:role_permissions;"`
	Users       []User         `json:"-" gorm:"foreignKey:RoleID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// Permission represents permissions in the RBAC system
type Permission struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"uniqueIndex;not null"`
	Resource    string         `json:"resource" gorm:"not null"`
	Action      string         `json:"action" gorm:"not null"`
	Description string         `json:"description"`
	Roles       []Role         `json:"-" gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	TenantSlug string `json:"tenant_slug,omitempty"` // Optional, for tenant-specific login
}

// LoginResponse represents the login response
type LoginResponse struct {
	User         User   `json:"user"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// RefreshTokenRequest represents the refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}