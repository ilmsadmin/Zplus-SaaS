package types

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// TenantID represents a tenant identifier with validation
type TenantID string

// Validate checks if the tenant ID is valid
func (t TenantID) Validate() error {
	if t == "" {
		return fmt.Errorf("tenant ID cannot be empty")
	}
	
	// Tenant ID should be alphanumeric with hyphens, 3-50 characters
	str := string(t)
	if len(str) < 3 || len(str) > 50 {
		return fmt.Errorf("tenant ID must be between 3 and 50 characters")
	}
	
	for _, r := range str {
		if !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-') {
			return fmt.Errorf("tenant ID can only contain lowercase letters, numbers, and hyphens")
		}
	}
	
	if strings.HasPrefix(str, "-") || strings.HasSuffix(str, "-") {
		return fmt.Errorf("tenant ID cannot start or end with a hyphen")
	}
	
	return nil
}

// String returns the string representation
func (t TenantID) String() string {
	return string(t)
}

// IsEmpty checks if tenant ID is empty
func (t TenantID) IsEmpty() bool {
	return string(t) == ""
}

// MarshalGQL implements the graphql.Marshaler interface
func (t TenantID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(t)))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (t *TenantID) UnmarshalGQL(v interface{}) error {
	switch s := v.(type) {
	case string:
		*t = TenantID(s)
		return t.Validate()
	case nil:
		*t = ""
		return nil
	default:
		return fmt.Errorf("tenant ID must be a string")
	}
}

// TenantContext represents the current tenant context for a request
type TenantContext struct {
	ID       TenantID `json:"id"`
	Slug     string   `json:"slug"`
	Name     string   `json:"name"`
	Schema   string   `json:"schema"`
	Status   string   `json:"status"`
	PlanID   string   `json:"plan_id"`
	Features []string `json:"features"`
}

// IsActive checks if the tenant is in active status
func (tc *TenantContext) IsActive() bool {
	return tc.Status == "ACTIVE" || tc.Status == "TRIAL"
}

// HasFeature checks if the tenant has access to a specific feature
func (tc *TenantContext) HasFeature(feature string) bool {
	for _, f := range tc.Features {
		if f == feature {
			return true
		}
	}
	return false
}

// UserContext represents the current user context for a request
type UserContext struct {
	ID          string    `json:"id"`
	TenantID    TenantID  `json:"tenant_id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
	IsAdmin     bool      `json:"is_admin"`
}

// HasRole checks if the user has a specific role
func (uc *UserContext) HasRole(role string) bool {
	for _, r := range uc.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// HasPermission checks if the user has a specific permission
func (uc *UserContext) HasPermission(permission string) bool {
	if uc.IsAdmin {
		return true
	}
	
	for _, p := range uc.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

// CanAccessResource checks if user can access a resource with action
func (uc *UserContext) CanAccessResource(resource, action string) bool {
	permission := fmt.Sprintf("%s:%s", resource, action)
	return uc.HasPermission(permission)
}

// RequestContext combines tenant and user context for GraphQL resolvers
type RequestContext struct {
	Tenant *TenantContext `json:"tenant,omitempty"`
	User   *UserContext   `json:"user,omitempty"`
}

// IsAuthenticated checks if there's a valid user in the context
func (rc *RequestContext) IsAuthenticated() bool {
	return rc.User != nil
}

// IsTenantAdmin checks if the user is an admin within their tenant
func (rc *RequestContext) IsTenantAdmin() bool {
	return rc.User != nil && rc.User.HasRole("admin")
}

// IsSystemAdmin checks if the user is a system-level admin
func (rc *RequestContext) IsSystemAdmin() bool {
	return rc.User != nil && rc.User.IsAdmin
}

// ValidateTenantAccess ensures user belongs to the current tenant
func (rc *RequestContext) ValidateTenantAccess() error {
	if rc.Tenant == nil {
		return fmt.Errorf("no tenant context available")
	}
	
	if rc.User == nil {
		return fmt.Errorf("authentication required")
	}
	
	if rc.User.TenantID != rc.Tenant.ID {
		return fmt.Errorf("user does not belong to the current tenant")
	}
	
	if !rc.Tenant.IsActive() {
		return fmt.Errorf("tenant is not active")
	}
	
	return nil
}