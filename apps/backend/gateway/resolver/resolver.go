package resolver

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/types"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// Database connection
	db *gorm.DB
	
	// Services
	tenantService      *services.TenantService
	planService        *services.PlanService
	subscriptionService *services.SubscriptionService
}

// NewResolver creates a new resolver instance
func NewResolver() *Resolver {
	return &Resolver{
		// Initialize dependencies here
		// For now, we'll use nil until we integrate database connection
	}
}

// SetDatabase sets the database connection and initializes services
func (r *Resolver) SetDatabase(db *gorm.DB) {
	r.db = db
	r.tenantService = services.NewTenantService(db)
	r.planService = services.NewPlanService(db)
	r.subscriptionService = services.NewSubscriptionService(db)
}

// GetUserService returns a user service for the given tenant
func (r *Resolver) GetUserService(tenantID string) *services.UserService {
	if r.db == nil {
		return nil
	}
	// Parse UUID from string
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		return nil
	}
	return services.NewUserService(r.db, tenantUUID)
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
