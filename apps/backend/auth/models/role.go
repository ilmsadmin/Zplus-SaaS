package models

import (
	"time"
)

// Role represents a role in the system
type Role struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	DisplayName  string    `json:"display_name" db:"display_name"`
	Description  string    `json:"description" db:"description"`
	IsSystemRole bool      `json:"is_system_role" db:"is_system_role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Permission represents a permission in the system
type Permission struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Resource    string    `json:"resource" db:"resource"`
	Action      string    `json:"action" db:"action"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// UserRole represents the junction table between users and roles
type UserRole struct {
	ID         int       `json:"id" db:"id"`
	UserID     string    `json:"user_id" db:"user_id"`
	RoleID     int       `json:"role_id" db:"role_id"`
	AssignedAt time.Time `json:"assigned_at" db:"assigned_at"`
}

// RolePermission represents the junction table between roles and permissions
type RolePermission struct {
	ID           int `json:"id" db:"id"`
	RoleID       int `json:"role_id" db:"role_id"`
	PermissionID int `json:"permission_id" db:"permission_id"`
}

// CreateRoleRequest represents the request to create a new role
type CreateRoleRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=100"`
	DisplayName  string `json:"display_name" validate:"required,min=3,max=255"`
	Description  string `json:"description" validate:"max=1000"`
	IsSystemRole bool   `json:"is_system_role"`
}

// UpdateRoleRequest represents the request to update a role
type UpdateRoleRequest struct {
	DisplayName string `json:"display_name" validate:"min=3,max=255"`
	Description string `json:"description" validate:"max=1000"`
}

// CreatePermissionRequest represents the request to create a new permission
type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Resource    string `json:"resource" validate:"required,min=3,max=100"`
	Action      string `json:"action" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=1000"`
}

// AssignRoleRequest represents the request to assign a role to a user
type AssignRoleRequest struct {
	UserID string `json:"user_id" validate:"required"`
	RoleID int    `json:"role_id" validate:"required,min=1"`
}

// AssignPermissionRequest represents the request to assign a permission to a role
type AssignPermissionRequest struct {
	RoleID       int `json:"role_id" validate:"required,min=1"`
	PermissionID int `json:"permission_id" validate:"required,min=1"`
}

// RoleWithPermissions represents a role with its associated permissions
type RoleWithPermissions struct {
	Role        Role         `json:"role"`
	Permissions []Permission `json:"permissions"`
}

// UserWithRoles represents a user with their assigned roles
type UserWithRoles struct {
	User  User   `json:"user"`
	Roles []Role `json:"roles"`
}