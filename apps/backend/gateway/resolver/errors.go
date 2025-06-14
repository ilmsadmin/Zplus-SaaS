package resolver

import "errors"

// Common GraphQL errors for multi-tenant operations
var (
	ErrUnauthenticated = errors.New("authentication required")
	ErrForbidden       = errors.New("access forbidden")
	ErrNotFound        = errors.New("resource not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrTenantMismatch  = errors.New("tenant mismatch")
	ErrInactiveTenant  = errors.New("tenant is not active")
	ErrFeatureDisabled = errors.New("feature not enabled for this tenant")
)