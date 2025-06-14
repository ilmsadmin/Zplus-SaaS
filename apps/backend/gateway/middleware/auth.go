package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/types"
)

// ContextKey represents the type for context keys
type ContextKey string

const (
	// TenantContextKey is the key for tenant context in request context
	TenantContextKey ContextKey = "tenant"
	// UserContextKey is the key for user context in request context
	UserContextKey ContextKey = "user"
	// RequestContextKey is the key for combined request context
	RequestContextKey ContextKey = "request"
)

// TenantMiddleware extracts tenant information from request and validates it
func TenantMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract tenant slug from subdomain or X-Tenant-ID header
		tenantSlug := extractTenantSlug(c)
		
		if tenantSlug == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Tenant identification required",
				"code":  "TENANT_REQUIRED",
			})
		}
		
		// Validate and get tenant context
		// In a real application, this would query the database
		tenantCtx, err := validateAndGetTenant(tenantSlug)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
				"code":  "INVALID_TENANT",
			})
		}
		
		// Store tenant context in Fiber locals for HTTP handlers
		c.Locals("tenant", tenantCtx)
		
		return c.Next()
	}
}

// AuthMiddleware validates JWT tokens and extracts user information
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip auth for certain endpoints
		if shouldSkipAuth(c.Path()) {
			return c.Next()
		}
		
		// Extract token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header required",
				"code":  "AUTH_REQUIRED",
			})
		}
		
		// Validate Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
				"code":  "INVALID_AUTH_FORMAT",
			})
		}
		
		token := parts[1]
		
		// Validate JWT token and extract user context
		// In a real application, this would validate JWT and query user data
		userCtx, err := validateJWTAndGetUser(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
				"code":  "INVALID_TOKEN",
			})
		}
		
		// Ensure user belongs to the current tenant
		tenantCtx, ok := c.Locals("tenant").(*types.TenantContext)
		if ok && tenantCtx != nil {
			if userCtx.TenantID != tenantCtx.ID {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "User does not belong to the current tenant",
					"code":  "TENANT_MISMATCH",
				})
			}
		}
		
		// Store user context in Fiber locals
		c.Locals("user", userCtx)
		
		return c.Next()
	}
}

// GraphQLContextMiddleware creates GraphQL context with tenant and user information
func GraphQLContextMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get tenant and user from Fiber locals
		var tenantCtx *types.TenantContext
		var userCtx *types.UserContext
		
		if tenant := c.Locals("tenant"); tenant != nil {
			tenantCtx = tenant.(*types.TenantContext)
		}
		
		if user := c.Locals("user"); user != nil {
			userCtx = user.(*types.UserContext)
		}
		
		// Create request context for GraphQL resolvers
		requestCtx := &types.RequestContext{
			Tenant: tenantCtx,
			User:   userCtx,
		}
		
		// Store in Fiber locals for GraphQL handler
		c.Locals("graphql_context", requestCtx)
		
		return c.Next()
	}
}

// extractTenantSlug extracts tenant identifier from request
func extractTenantSlug(c *fiber.Ctx) string {
	// First try X-Tenant-ID header (set by Traefik middleware)
	if tenantID := c.Get("X-Tenant-ID"); tenantID != "" {
		return tenantID
	}
	
	// Fallback to extracting from Host header
	host := c.Get("Host")
	if host == "" {
		return ""
	}
	
	// Extract subdomain from host (e.g., "tenant.zplus.com" -> "tenant")
	parts := strings.Split(host, ".")
	if len(parts) >= 3 {
		return parts[0]
	}
	
	return ""
}

// shouldSkipAuth determines if authentication should be skipped for certain endpoints
func shouldSkipAuth(path string) bool {
	skipPaths := []string{
		"/",
		"/health",
		"/metrics",
		"/graphql", // We'll handle auth in GraphQL resolvers
	}
	
	for _, skipPath := range skipPaths {
		if path == skipPath {
			return true
		}
	}
	
	return false
}

// validateAndGetTenant validates tenant and returns tenant context
// In a real application, this would query the database
func validateAndGetTenant(slug string) (*types.TenantContext, error) {
	tenantID := types.TenantID(slug)
	if err := tenantID.Validate(); err != nil {
		return nil, fmt.Errorf("invalid tenant slug: %v", err)
	}
	
	// Mock tenant validation - replace with actual database query
	mockTenantData := map[string]*types.TenantContext{
		"demo": {
			ID:     "demo",
			Slug:   "demo",
			Name:   "Demo Tenant",
			Schema: "tenant_demo",
			Status: "ACTIVE",
			PlanID: "startup",
			Features: []string{
				"CRM",
				"HRM",
				"POS",
				"BASIC_ANALYTICS",
			},
		},
		"acme": {
			ID:     "acme",
			Slug:   "acme",
			Name:   "ACME Corporation",
			Schema: "tenant_acme",
			Status: "ACTIVE",
			PlanID: "enterprise",
			Features: []string{
				"CRM",
				"HRM",
				"POS",
				"LMS",
				"ADVANCED_ANALYTICS",
				"API_ACCESS",
			},
		},
	}
	
	tenant, exists := mockTenantData[slug]
	if !exists {
		return nil, fmt.Errorf("tenant not found: %s", slug)
	}
	
	if !tenant.IsActive() {
		return nil, fmt.Errorf("tenant is not active: %s", slug)
	}
	
	return tenant, nil
}

// validateJWTAndGetUser validates JWT token and returns user context
// In a real application, this would validate JWT signature and query user data
func validateJWTAndGetUser(token string) (*types.UserContext, error) {
	// Mock JWT validation - replace with actual JWT validation logic
	if token == "" {
		return nil, fmt.Errorf("empty token")
	}
	
	// Mock user data based on token
	mockUserData := map[string]*types.UserContext{
		"demo_admin_token": {
			ID:        "1",
			TenantID:  "demo",
			Email:     "admin@demo.zplus.com",
			FirstName: "Demo",
			LastName:  "Admin",
			Roles:     []string{"admin", "user"},
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
			IsAdmin: false,
		},
		"acme_user_token": {
			ID:        "2",
			TenantID:  "acme",
			Email:     "user@acme.com",
			FirstName: "John",
			LastName:  "Doe",
			Roles:     []string{"user"},
			Permissions: []string{
				"customers:read",
				"products:read",
			},
			IsAdmin: false,
		},
		"system_admin_token": {
			ID:        "999",
			TenantID:  "",
			Email:     "system@zplus.com",
			FirstName: "System",
			LastName:  "Admin",
			Roles:     []string{"system_admin"},
			Permissions: []string{},
			IsAdmin: true,
		},
	}
	
	user, exists := mockUserData[token]
	if !exists {
		return nil, fmt.Errorf("invalid token")
	}
	
	return user, nil
}

// GetRequestContext extracts request context from fiber context
func GetRequestContext(c *fiber.Ctx) *types.RequestContext {
	if ctx := c.Locals("graphql_context"); ctx != nil {
		return ctx.(*types.RequestContext)
	}
	return &types.RequestContext{}
}

// GetTenantContext extracts tenant context from fiber context
func GetTenantContext(c *fiber.Ctx) *types.TenantContext {
	if ctx := c.Locals("tenant"); ctx != nil {
		return ctx.(*types.TenantContext)
	}
	return nil
}

// GetUserContext extracts user context from fiber context
func GetUserContext(c *fiber.Ctx) *types.UserContext {
	if ctx := c.Locals("user"); ctx != nil {
		return ctx.(*types.UserContext)
	}
	return nil
}