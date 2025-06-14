package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared"
	"github.com/ilmsadmin/Zplus-SaaS/pkg/auth"
)

var tokenManager *auth.TokenManager

func initTokenManager(secret, issuer string) {
	tokenManager = auth.NewTokenManager(secret, issuer)
}
func setupRoutes(app *fiber.App) {
	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "auth",
			"status":  "running",
			"message": "Authentication & RBAC Service",
		})
	})

	api := app.Group("/api/v1")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", login)
	auth.Post("/refresh", refreshToken)
	auth.Post("/logout", authMiddleware, logout)
	auth.Get("/me", authMiddleware, getProfile)

	// Protected routes for user management (system/tenant admin only)
	users := api.Group("/users", authMiddleware)
	users.Get("/", getUsers)
	users.Post("/", createUser)
	users.Get("/:id", getUser)
	users.Put("/:id", updateUser)
	users.Delete("/:id", deleteUser)

	// Role and permission routes
	rbac := api.Group("/rbac", authMiddleware)
	rbac.Get("/roles", getRoles)
	rbac.Get("/permissions", getPermissions)
}

// Authentication handlers
func login(c *fiber.Ctx) error {
	var req shared.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Find user by email
	var user shared.User
	query := shared.DB.Preload("Role").Preload("Tenant")
	
	// If tenant slug is provided, filter by tenant
	if req.TenantSlug != "" {
		query = query.Joins("JOIN tenants ON users.tenant_id = tenants.id").
			Where("users.email = ? AND tenants.slug = ?", req.Email, req.TenantSlug)
	} else {
		query = query.Where("email = ?", req.Email)
	}

	if err := query.First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Check if user is active
	if !user.IsActive {
		return c.Status(401).JSON(fiber.Map{"error": "Account is disabled"})
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate tokens
	tenantID := ""
	if user.TenantID != nil {
		tenantID = fmt.Sprintf("%d", *user.TenantID)
	}

	token, err := tokenManager.GenerateToken(
		fmt.Sprintf("%d", user.ID),
		tenantID,
		user.Role.Name,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Generate refresh token (valid for 7 days)
	refreshToken, err := tokenManager.GenerateRefreshToken(fmt.Sprintf("%d", user.ID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate refresh token"})
	}

	response := shared.LoginResponse{
		User:         user,
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    24 * 60 * 60, // 24 hours in seconds
	}

	return c.JSON(response)
}

func refreshToken(c *fiber.Ctx) error {
	var req shared.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate refresh token
	claims, err := tokenManager.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	// Get user
	userID, _ := strconv.ParseUint(claims.UserID, 10, 32)
	var user shared.User
	if err := shared.DB.Preload("Role").Preload("Tenant").First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Generate new access token
	tenantID := ""
	if user.TenantID != nil {
		tenantID = fmt.Sprintf("%d", *user.TenantID)
	}

	newToken, err := tokenManager.GenerateToken(
		fmt.Sprintf("%d", user.ID),
		tenantID,
		user.Role.Name,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"token":      newToken,
		"expires_in": 24 * 60 * 60,
	})
}

func logout(c *fiber.Ctx) error {
	// In a production system, you might want to blacklist the token
	return c.JSON(fiber.Map{"message": "Logged out successfully"})
}

func getProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(shared.User)
	return c.JSON(user)
}

// User management handlers
func getUsers(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(shared.User)
	
	var users []shared.User
	query := shared.DB.Preload("Role").Preload("Tenant")

	// Apply tenant-based filtering
	if currentUser.Role.Level == "tenant" && currentUser.TenantID != nil {
		query = query.Where("tenant_id = ?", *currentUser.TenantID)
	} else if currentUser.Role.Level != "system" {
		return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
	}

	if err := query.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}

	return c.JSON(users)
}

func createUser(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(shared.User)
	
	// Check permissions
	if currentUser.Role.Level != "system" && currentUser.Role.Level != "tenant" {
		return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
	}

	var user shared.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	// Set tenant ID for tenant admins
	if currentUser.Role.Level == "tenant" && currentUser.TenantID != nil {
		user.TenantID = currentUser.TenantID
	}

	if err := shared.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Load related data
	shared.DB.Preload("Role").Preload("Tenant").First(&user, user.ID)

	return c.Status(201).JSON(user)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	currentUser := c.Locals("user").(shared.User)

	var user shared.User
	query := shared.DB.Preload("Role").Preload("Tenant")

	// Apply tenant-based filtering
	if currentUser.Role.Level == "tenant" && currentUser.TenantID != nil {
		query = query.Where("tenant_id = ?", *currentUser.TenantID)
	} else if currentUser.Role.Level != "system" {
		return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
	}

	if err := query.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	currentUser := c.Locals("user").(shared.User)

	// Check permissions
	if currentUser.Role.Level != "system" && currentUser.Role.Level != "tenant" {
		return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
	}

	var user shared.User
	query := shared.DB

	// Apply tenant-based filtering
	if currentUser.Role.Level == "tenant" && currentUser.TenantID != nil {
		query = query.Where("tenant_id = ?", *currentUser.TenantID)
	}

	if err := query.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	var updateData shared.User
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Hash password if provided
	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		updateData.Password = string(hashedPassword)
	}

	if err := shared.DB.Model(&user).Updates(updateData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}

	// Load updated user with related data
	shared.DB.Preload("Role").Preload("Tenant").First(&user, user.ID)

	return c.JSON(user)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	currentUser := c.Locals("user").(shared.User)

	// Check permissions
	if currentUser.Role.Level != "system" && currentUser.Role.Level != "tenant" {
		return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
	}

	var user shared.User
	query := shared.DB

	// Apply tenant-based filtering
	if currentUser.Role.Level == "tenant" && currentUser.TenantID != nil {
		query = query.Where("tenant_id = ?", *currentUser.TenantID)
	}

	if err := query.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if err := shared.DB.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

// RBAC handlers
func getRoles(c *fiber.Ctx) error {
	var roles []shared.Role
	if err := shared.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve roles"})
	}
	return c.JSON(roles)
}

func getPermissions(c *fiber.Ctx) error {
	var permissions []shared.Permission
	if err := shared.DB.Find(&permissions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve permissions"})
	}
	return c.JSON(permissions)
}

// Middleware
func authMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Authorization header required"})
	}

	// Extract token from "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization header format"})
	}

	tokenString := parts[1]

	// Validate token
	claims, err := tokenManager.ValidateToken(tokenString)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Get user from database
	userID, _ := strconv.ParseUint(claims.UserID, 10, 32)
	var user shared.User
	if err := shared.DB.Preload("Role").Preload("Tenant").First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Check if user is still active
	if !user.IsActive {
		return c.Status(401).JSON(fiber.Map{"error": "Account is disabled"})
	}

	// Store user in context
	c.Locals("user", user)
	c.Locals("claims", claims)

	return c.Next()
}

// Utility functions
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}