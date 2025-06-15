package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth/models"
	"github.com/ilmsadmin/Zplus-SaaS/pkg/auth"
)

// AuthHandler handles authentication endpoints
type AuthHandler struct {
	tokenManager *auth.TokenManager
	users        map[string]*models.User // Mock user store
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler() *AuthHandler {
	// Initialize with mock users for different levels
	tokenManager := auth.NewTokenManager("your-secret-key", "zplus-saas")
	
	// Create mock users with hashed passwords
	users := make(map[string]*models.User)
	
	// System Admin
	systemAdmin := &models.User{
		ID:        "sys-admin-1",
		TenantID:  "system",
		Email:     "admin@zplus.com",
		FirstName: "System",
		LastName:  "Admin",
		Roles:     []string{"system_admin"},
		IsAdmin:   true,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Permissions: []string{
			"system:manage",
			"tenants:read",
			"tenants:write",
			"users:read",
			"users:write",
		},
	}
	systemAdmin.HashPassword("admin123")
	users["admin@zplus.com|system"] = systemAdmin
	
	// Tenant Admin (Demo Corp)
	tenantAdmin := &models.User{
		ID:        "tenant-admin-1",
		TenantID:  "demo-corp",
		Email:     "admin@demo-corp.zplus.com",
		FirstName: "Demo",
		LastName:  "Admin",
		Roles:     []string{"tenant_admin"},
		IsAdmin:   false,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Permissions: []string{
			"users:read",
			"users:write",
			"customers:read",
			"customers:write",
			"employees:read",
			"employees:write",
			"products:read",
			"products:write",
		},
	}
	tenantAdmin.HashPassword("demo123")
	users["admin@demo-corp.zplus.com|demo-corp"] = tenantAdmin
	
	// Customer User (CRM)
	customerUser := &models.User{
		ID:        "customer-1",
		TenantID:  "demo-corp",
		Email:     "john@demo-corp.zplus.com",
		FirstName: "John",
		LastName:  "Doe",
		Roles:     []string{"user"},
		IsAdmin:   false,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Permissions: []string{
			"customers:read",
			"products:read",
		},
	}
	customerUser.HashPassword("user123")
	users["john@demo-corp.zplus.com|demo-corp"] = customerUser
	
	return &AuthHandler{
		tokenManager: tokenManager,
		users:        users,
	}
}

// Login handles user login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" || req.TenantSlug == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Missing required fields",
			Code:    "VALIDATION_ERROR",
			Message: "Email, password, and tenant_slug are required",
		})
	}

	// Normalize email
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.TenantSlug = strings.ToLower(strings.TrimSpace(req.TenantSlug))

	// Find user by email and tenant
	userKey := fmt.Sprintf("%s|%s", req.Email, req.TenantSlug)
	user, exists := h.users[userKey]
	if !exists {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error:   "Invalid credentials",
			Code:    "INVALID_CREDENTIALS",
			Message: "Email or password is incorrect",
		})
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error:   "Invalid credentials",
			Code:    "INVALID_CREDENTIALS",
			Message: "Email or password is incorrect",
		})
	}

	// Check if user is active
	if user.Status != "active" {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error:   "Account disabled",
			Code:    "ACCOUNT_DISABLED",
			Message: "Your account has been disabled. Please contact support",
		})
	}

	// Determine role for JWT
	role := "user"
	if user.IsAdmin && user.TenantID == "system" {
		role = "system_admin"
	} else if len(user.Roles) > 0 {
		role = user.Roles[0]
	}

	// Generate JWT token
	token, err := h.tokenManager.GenerateToken(user.ID, user.TenantID, role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Token generation failed",
			Code:    "TOKEN_ERROR",
			Message: "Unable to generate authentication token",
		})
	}

	// Generate refresh token (same for now, in production use different logic)
	refreshToken, err := h.tokenManager.GenerateToken(user.ID, user.TenantID, role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Token generation failed",
			Code:    "TOKEN_ERROR",
			Message: "Unable to generate refresh token",
		})
	}

	// Return successful login response
	return c.JSON(models.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User:         user,
		ExpiresIn:    86400, // 24 hours in seconds
	})
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// Extract token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Authorization header required",
			Code:    "AUTH_REQUIRED",
			Message: "Please provide a valid authorization token",
		})
	}

	// Validate Bearer token format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid authorization header format",
			Code:    "INVALID_AUTH_FORMAT",
			Message: "Authorization header must be in format: Bearer <token>",
		})
	}

	token := parts[1]

	// Invalidate the token
	if err := h.tokenManager.InvalidateToken(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error:   "Invalid token",
			Code:    "INVALID_TOKEN",
			Message: "Token is invalid or already expired",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully logged out",
	})
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "Please provide valid JSON data",
		})
	}

	if req.RefreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Missing refresh token",
			Code:    "VALIDATION_ERROR",
			Message: "Refresh token is required",
		})
	}

	// Validate refresh token
	claims, err := h.tokenManager.ValidateToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error:   "Invalid refresh token",
			Code:    "INVALID_TOKEN",
			Message: "Refresh token is invalid or expired",
		})
	}

	// Generate new access token
	newToken, err := h.tokenManager.GenerateToken(claims.UserID, claims.TenantID, claims.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Token generation failed",
			Code:    "TOKEN_ERROR",
			Message: "Unable to generate new access token",
		})
	}

	// Find user for response
	var user *models.User
	for _, u := range h.users {
		if u.ID == claims.UserID && u.TenantID == claims.TenantID {
			user = u
			break
		}
	}

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error:   "User not found",
			Code:    "USER_NOT_FOUND",
			Message: "User associated with token not found",
		})
	}

	return c.JSON(models.LoginResponse{
		Token:        newToken,
		RefreshToken: req.RefreshToken, // Keep the same refresh token
		User:         user,
		ExpiresIn:    86400, // 24 hours in seconds
	})
}

// GetUsers returns all mock users (for testing)
func (h *AuthHandler) GetUsers(c *fiber.Ctx) error {
	users := make([]*models.User, 0, len(h.users))
	for _, user := range h.users {
		users = append(users, user)
	}
	return c.JSON(fiber.Map{
		"users": users,
		"count": len(users),
	})
}