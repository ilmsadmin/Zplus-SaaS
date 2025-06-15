package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/middleware"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/services"
)

// UserHandler handles tenant-scoped user CRUD operations
type UserHandler struct {
	getUserService func(tenantID string) *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(getUserService func(tenantID string) *services.UserService) *UserHandler {
	return &UserHandler{
		getUserService: getUserService,
	}
}

// GetUsers retrieves all users in the current tenant with pagination and filtering
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	// Get tenant context
	tenantCtx := middleware.GetTenantContext(c)
	if tenantCtx == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Tenant context not available",
		})
	}

	userService := h.getUserService(string(tenantCtx.ID))
	if userService == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "User service not available",
		})
	}

	// Parse pagination params
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	if limit > 100 {
		limit = 100 // Max limit
	}
	offset := (page - 1) * limit

	// Parse filters
	filter := services.UserFilter{
		Status: c.Query("status"),
		Role:   c.Query("role"),
		Search: c.Query("search"),
	}

	// Get users
	users, total, err := userService.ListUsers(filter, offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetUser retrieves a single user by ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	// Get tenant context
	tenantCtx := middleware.GetTenantContext(c)
	if tenantCtx == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Tenant context not available",
		})
	}

	userService := h.getUserService(string(tenantCtx.ID))
	if userService == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "User service not available",
		})
	}

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid UUID",
		})
	}

	user, err := userService.GetUser(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "User not found",
				"message": "No user found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve user",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// CreateUser creates a new user in the current tenant
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Get tenant context
	tenantCtx := middleware.GetTenantContext(c)
	if tenantCtx == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Tenant context not available",
		})
	}

	userService := h.getUserService(string(tenantCtx.ID))
	if userService == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "User service not available",
		})
	}

	var input services.CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if input.Email == "" || input.Password == "" || input.FirstName == "" || input.LastName == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Missing required fields",
			"message": "Email, password, first name, and last name are required",
		})
	}

	user, err := userService.CreateUser(input)
	if err != nil {
		if err.Error() == "user with email '"+input.Email+"' already exists in this tenant" {
			return c.Status(409).JSON(fiber.Map{
				"error":   "User already exists",
				"message": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data":    user,
		"message": "User created successfully",
	})
}

// UpdateUser updates an existing user
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	// Get tenant context
	tenantCtx := middleware.GetTenantContext(c)
	if tenantCtx == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Tenant context not available",
		})
	}

	userService := h.getUserService(string(tenantCtx.ID))
	if userService == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "User service not available",
		})
	}

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid UUID",
		})
	}

	var input services.UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	user, err := userService.UpdateUser(userID, input)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "User not found",
				"message": "No user found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update user",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":    user,
		"message": "User updated successfully",
	})
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	// Get tenant context
	tenantCtx := middleware.GetTenantContext(c)
	if tenantCtx == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Tenant context not available",
		})
	}

	userService := h.getUserService(string(tenantCtx.ID))
	if userService == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "User service not available",
		})
	}

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid UUID",
		})
	}

	err = userService.DeleteUser(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "User not found",
				"message": "No user found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete user",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

// ChangePassword changes a user's password
func (h *UserHandler) ChangePassword(c *fiber.Ctx) error {
	// Get tenant context
	tenantCtx := middleware.GetTenantContext(c)
	if tenantCtx == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Tenant context not available",
		})
	}

	userService := h.getUserService(string(tenantCtx.ID))
	if userService == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "User service not available",
		})
	}

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid UUID",
		})
	}

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	if input.OldPassword == "" || input.NewPassword == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Missing required fields",
			"message": "Old password and new password are required",
		})
	}

	err = userService.ChangePassword(userID, input.OldPassword, input.NewPassword)
	if err != nil {
		if err.Error() == "current password is incorrect" {
			return c.Status(400).JSON(fiber.Map{
				"error":   "Invalid password",
				"message": err.Error(),
			})
		}
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "User not found",
				"message": "No user found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to change password",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}

// AssignRoles assigns roles to a user
func (h *UserHandler) AssignRoles(c *fiber.Ctx) error {
	// Get tenant context
	tenantCtx := middleware.GetTenantContext(c)
	if tenantCtx == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Tenant context not available",
		})
	}

	userService := h.getUserService(string(tenantCtx.ID))
	if userService == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "User service not available",
		})
	}

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid UUID",
		})
	}

	var input struct {
		RoleIDs []uuid.UUID `json:"role_ids"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	err = userService.AssignRoles(userID, input.RoleIDs)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "User not found",
				"message": "No user found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to assign roles",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Roles assigned successfully",
	})
}