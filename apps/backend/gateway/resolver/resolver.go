package resolver

import (
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/types"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// Add any dependencies here like database connections, external services, etc.
	// db     *gorm.DB
	// redis  *redis.Client
	// logger *log.Logger
}

// NewResolver creates a new resolver instance
func NewResolver() *Resolver {
	return &Resolver{
		// Initialize dependencies here
	}
}

// Helper methods for multi-tenant operations

// validateTenantAccess ensures the user has access to the current tenant
func (r *Resolver) validateTenantAccess(ctx *types.RequestContext) error {
	return ctx.ValidateTenantAccess()
}

// requireAuth ensures the user is authenticated
func (r *Resolver) requireAuth(ctx *types.RequestContext) error {
	if !ctx.IsAuthenticated() {
		return ErrUnauthenticated
	}
	return nil
}

// requireTenantAuth ensures the user is authenticated and belongs to the tenant
func (r *Resolver) requireTenantAuth(ctx *types.RequestContext) error {
	if err := r.requireAuth(ctx); err != nil {
		return err
	}
	return r.validateTenantAccess(ctx)
}

// requireSystemAdmin ensures the user is a system administrator
func (r *Resolver) requireSystemAdmin(ctx *types.RequestContext) error {
	if err := r.requireAuth(ctx); err != nil {
		return err
	}
	if !ctx.IsSystemAdmin() {
		return ErrForbidden
	}
	return nil
}

// requireTenantAdmin ensures the user is an admin within their tenant
func (r *Resolver) requireTenantAdmin(ctx *types.RequestContext) error {
	if err := r.requireTenantAuth(ctx); err != nil {
		return err
	}
	if !ctx.IsTenantAdmin() {
		return ErrForbidden
	}
	return nil
}

// requirePermission checks if the user has a specific permission
func (r *Resolver) requirePermission(ctx *types.RequestContext, resource, action string) error {
	if err := r.requireTenantAuth(ctx); err != nil {
		return err
	}
	if !ctx.User.CanAccessResource(resource, action) {
		return ErrForbidden
	}
	return nil
}
