package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth/models"
)

// RoleHandler handles role and permission management endpoints
type RoleHandler struct {
	// Mock storage for roles and permissions
	roles           map[int]*models.Role
	permissions     map[int]*models.Permission
	userRoles       map[string][]int // user_id -> role_ids
	rolePermissions map[int][]int    // role_id -> permission_ids
	nextRoleID      int
	nextPermID      int
}

// NewRoleHandler creates a new role handler with mock data
func NewRoleHandler() *RoleHandler {
	handler := &RoleHandler{
		roles:           make(map[int]*models.Role),
		permissions:     make(map[int]*models.Permission),
		userRoles:       make(map[string][]int),
		rolePermissions: make(map[int][]int),
		nextRoleID:      1,
		nextPermID:      1,
	}

	// Initialize with default roles and permissions
	handler.initializeDefaultData()
	return handler
}

// initializeDefaultData creates default roles and permissions
func (h *RoleHandler) initializeDefaultData() {
	now := time.Now()

	// Create default roles
	roles := []*models.Role{
		{
			ID:           h.nextRoleID,
			Name:         "system_admin",
			DisplayName:  "System Administrator",
			Description:  "Full system access and management",
			IsSystemRole: true,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			ID:           h.nextRoleID + 1,
			Name:         "tenant_admin",
			DisplayName:  "Tenant Administrator",
			Description:  "Full tenant access and management",
			IsSystemRole: true,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			ID:           h.nextRoleID + 2,
			Name:         "manager",
			DisplayName:  "Manager",
			Description:  "Management level access",
			IsSystemRole: true,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			ID:           h.nextRoleID + 3,
			Name:         "user",
			DisplayName:  "User",
			Description:  "Standard user access",
			IsSystemRole: true,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	for _, role := range roles {
		h.roles[role.ID] = role
		h.nextRoleID++
	}

	// Create default permissions
	permissions := []*models.Permission{
		{ID: h.nextPermID, Name: "system:manage", Resource: "system", Action: "manage", Description: "System management", CreatedAt: now},
		{ID: h.nextPermID + 1, Name: "tenants:read", Resource: "tenants", Action: "read", Description: "Read tenants", CreatedAt: now},
		{ID: h.nextPermID + 2, Name: "tenants:write", Resource: "tenants", Action: "write", Description: "Write tenants", CreatedAt: now},
		{ID: h.nextPermID + 3, Name: "users:read", Resource: "users", Action: "read", Description: "Read users", CreatedAt: now},
		{ID: h.nextPermID + 4, Name: "users:write", Resource: "users", Action: "write", Description: "Write users", CreatedAt: now},
		{ID: h.nextPermID + 5, Name: "customers:read", Resource: "customers", Action: "read", Description: "Read customers", CreatedAt: now},
		{ID: h.nextPermID + 6, Name: "customers:write", Resource: "customers", Action: "write", Description: "Write customers", CreatedAt: now},
		{ID: h.nextPermID + 7, Name: "employees:read", Resource: "employees", Action: "read", Description: "Read employees", CreatedAt: now},
		{ID: h.nextPermID + 8, Name: "employees:write", Resource: "employees", Action: "write", Description: "Write employees", CreatedAt: now},
		{ID: h.nextPermID + 9, Name: "products:read", Resource: "products", Action: "read", Description: "Read products", CreatedAt: now},
		{ID: h.nextPermID + 10, Name: "products:write", Resource: "products", Action: "write", Description: "Write products", CreatedAt: now},
	}

	for _, perm := range permissions {
		h.permissions[perm.ID] = perm
		h.nextPermID++
	}

	// Assign permissions to roles
	// System Admin - all permissions
	h.rolePermissions[1] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// Tenant Admin - tenant-level permissions
	h.rolePermissions[2] = []int{4, 5, 6, 7, 8, 9, 10, 11}
	// Manager - read/write business data
	h.rolePermissions[3] = []int{4, 6, 7, 8, 9, 10, 11}
	// User - read-only access
	h.rolePermissions[4] = []int{4, 6, 8, 10}
}

// GetRoles returns all roles
func (h *RoleHandler) GetRoles(c *fiber.Ctx) error {
	roles := make([]*models.Role, 0, len(h.roles))
	for _, role := range h.roles {
		roles = append(roles, role)
	}

	return c.JSON(fiber.Map{
		"data":  roles,
		"count": len(roles),
	})
}

// GetRole returns a specific role by ID
func (h *RoleHandler) GetRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid role ID",
			Code:    "INVALID_ID",
			Message: "Role ID must be a valid integer",
		})
	}

	role, exists := h.roles[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Role not found",
			Code:    "ROLE_NOT_FOUND",
			Message: fmt.Sprintf("Role with ID %d not found", id),
		})
	}

	return c.JSON(fiber.Map{
		"data": role,
	})
}

// CreateRole creates a new role
func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var req models.CreateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if req.Name == "" || req.DisplayName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Missing required fields",
			Code:    "VALIDATION_ERROR",
			Message: "Name and display_name are required",
		})
	}

	// Check if role name already exists
	for _, role := range h.roles {
		if strings.EqualFold(role.Name, req.Name) {
			return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
				Error:   "Role already exists",
				Code:    "ROLE_EXISTS",
				Message: fmt.Sprintf("Role with name '%s' already exists", req.Name),
			})
		}
	}

	// Create new role
	now := time.Now()
	role := &models.Role{
		ID:           h.nextRoleID,
		Name:         strings.ToLower(strings.TrimSpace(req.Name)),
		DisplayName:  strings.TrimSpace(req.DisplayName),
		Description:  strings.TrimSpace(req.Description),
		IsSystemRole: req.IsSystemRole,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	h.roles[role.ID] = role
	h.nextRoleID++

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    role,
		"message": "Role created successfully",
	})
}

// UpdateRole updates an existing role
func (h *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid role ID",
			Code:    "INVALID_ID",
			Message: "Role ID must be a valid integer",
		})
	}

	role, exists := h.roles[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Role not found",
			Code:    "ROLE_NOT_FOUND",
			Message: fmt.Sprintf("Role with ID %d not found", id),
		})
	}

	var req models.UpdateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "Please provide valid JSON data",
		})
	}

	// Update role fields
	if req.DisplayName != "" {
		role.DisplayName = strings.TrimSpace(req.DisplayName)
	}
	if req.Description != "" {
		role.Description = strings.TrimSpace(req.Description)
	}
	role.UpdatedAt = time.Now()

	return c.JSON(fiber.Map{
		"data":    role,
		"message": "Role updated successfully",
	})
}

// DeleteRole deletes a role
func (h *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid role ID",
			Code:    "INVALID_ID",
			Message: "Role ID must be a valid integer",
		})
	}

	role, exists := h.roles[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Role not found",
			Code:    "ROLE_NOT_FOUND",
			Message: fmt.Sprintf("Role with ID %d not found", id),
		})
	}

	// Prevent deletion of system roles
	if role.IsSystemRole {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error:   "Cannot delete system role",
			Code:    "SYSTEM_ROLE_PROTECTED",
			Message: "System roles cannot be deleted",
		})
	}

	delete(h.roles, id)
	delete(h.rolePermissions, id)

	return c.JSON(fiber.Map{
		"message": "Role deleted successfully",
	})
}

// GetPermissions returns all permissions
func (h *RoleHandler) GetPermissions(c *fiber.Ctx) error {
	permissions := make([]*models.Permission, 0, len(h.permissions))
	for _, perm := range h.permissions {
		permissions = append(permissions, perm)
	}

	return c.JSON(fiber.Map{
		"data":  permissions,
		"count": len(permissions),
	})
}

// CreatePermission creates a new permission
func (h *RoleHandler) CreatePermission(c *fiber.Ctx) error {
	var req models.CreatePermissionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if req.Name == "" || req.Resource == "" || req.Action == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Missing required fields",
			Code:    "VALIDATION_ERROR",
			Message: "Name, resource, and action are required",
		})
	}

	// Check if permission already exists
	for _, perm := range h.permissions {
		if strings.EqualFold(perm.Name, req.Name) {
			return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
				Error:   "Permission already exists",
				Code:    "PERMISSION_EXISTS",
				Message: fmt.Sprintf("Permission with name '%s' already exists", req.Name),
			})
		}
	}

	// Create new permission
	now := time.Now()
	permission := &models.Permission{
		ID:          h.nextPermID,
		Name:        fmt.Sprintf("%s:%s", strings.ToLower(req.Resource), strings.ToLower(req.Action)),
		Resource:    strings.ToLower(strings.TrimSpace(req.Resource)),
		Action:      strings.ToLower(strings.TrimSpace(req.Action)),
		Description: strings.TrimSpace(req.Description),
		CreatedAt:   now,
	}

	h.permissions[permission.ID] = permission
	h.nextPermID++

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    permission,
		"message": "Permission created successfully",
	})
}

// AssignRoleToUser assigns a role to a user
func (h *RoleHandler) AssignRoleToUser(c *fiber.Ctx) error {
	var req models.AssignRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if req.UserID == "" || req.RoleID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Missing required fields",
			Code:    "VALIDATION_ERROR",
			Message: "User ID and role ID are required",
		})
	}

	// Check if role exists
	if _, exists := h.roles[req.RoleID]; !exists {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Role not found",
			Code:    "ROLE_NOT_FOUND",
			Message: fmt.Sprintf("Role with ID %d not found", req.RoleID),
		})
	}

	// Check if user already has this role
	userRoles := h.userRoles[req.UserID]
	for _, roleID := range userRoles {
		if roleID == req.RoleID {
			return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
				Error:   "Role already assigned",
				Code:    "ROLE_ALREADY_ASSIGNED",
				Message: "User already has this role assigned",
			})
		}
	}

	// Assign role to user
	h.userRoles[req.UserID] = append(h.userRoles[req.UserID], req.RoleID)

	return c.JSON(fiber.Map{
		"message": "Role assigned successfully",
	})
}

// AssignPermissionToRole assigns a permission to a role
func (h *RoleHandler) AssignPermissionToRole(c *fiber.Ctx) error {
	var req models.AssignPermissionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if req.RoleID == 0 || req.PermissionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Missing required fields",
			Code:    "VALIDATION_ERROR",
			Message: "Role ID and permission ID are required",
		})
	}

	// Check if role exists
	if _, exists := h.roles[req.RoleID]; !exists {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Role not found",
			Code:    "ROLE_NOT_FOUND",
			Message: fmt.Sprintf("Role with ID %d not found", req.RoleID),
		})
	}

	// Check if permission exists
	if _, exists := h.permissions[req.PermissionID]; !exists {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Permission not found",
			Code:    "PERMISSION_NOT_FOUND",
			Message: fmt.Sprintf("Permission with ID %d not found", req.PermissionID),
		})
	}

	// Check if role already has this permission
	rolePermissions := h.rolePermissions[req.RoleID]
	for _, permID := range rolePermissions {
		if permID == req.PermissionID {
			return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
				Error:   "Permission already assigned",
				Code:    "PERMISSION_ALREADY_ASSIGNED",
				Message: "Role already has this permission assigned",
			})
		}
	}

	// Assign permission to role
	h.rolePermissions[req.RoleID] = append(h.rolePermissions[req.RoleID], req.PermissionID)

	return c.JSON(fiber.Map{
		"message": "Permission assigned successfully",
	})
}

// GetRolePermissions returns all permissions for a specific role
func (h *RoleHandler) GetRolePermissions(c *fiber.Ctx) error {
	roleID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid role ID",
			Code:    "INVALID_ID",
			Message: "Role ID must be a valid integer",
		})
	}

	role, exists := h.roles[roleID]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Role not found",
			Code:    "ROLE_NOT_FOUND",
			Message: fmt.Sprintf("Role with ID %d not found", roleID),
		})
	}

	permissionIDs := h.rolePermissions[roleID]
	permissions := make([]*models.Permission, 0, len(permissionIDs))
	for _, permID := range permissionIDs {
		if perm, exists := h.permissions[permID]; exists {
			permissions = append(permissions, perm)
		}
	}

	return c.JSON(fiber.Map{
		"role":        role,
		"permissions": permissions,
		"count":       len(permissions),
	})
}

// GetUserRoles returns all roles for a specific user
func (h *RoleHandler) GetUserRoles(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid user ID",
			Code:    "INVALID_ID",
			Message: "User ID is required",
		})
	}

	roleIDs := h.userRoles[userID]
	roles := make([]*models.Role, 0, len(roleIDs))
	for _, roleID := range roleIDs {
		if role, exists := h.roles[roleID]; exists {
			roles = append(roles, role)
		}
	}

	return c.JSON(fiber.Map{
		"user_id": userID,
		"roles":   roles,
		"count":   len(roles),
	})
}