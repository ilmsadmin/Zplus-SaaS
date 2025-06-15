package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/models"
)

// TenantService handles CRUD operations for tenants
type TenantService struct {
	db *gorm.DB
}

// NewTenantService creates a new tenant service
func NewTenantService(db *gorm.DB) *TenantService {
	return &TenantService{db: db}
}

// CreateTenantInput represents input for creating a tenant
type CreateTenantInput struct {
	Name      string                 `json:"name" validate:"required"`
	Slug      string                 `json:"slug" validate:"required"`
	Domain    *string                `json:"domain"`
	Subdomain *string                `json:"subdomain"`
	PlanID    *uuid.UUID             `json:"plan_id"`
	Settings  map[string]interface{} `json:"settings"`
}

// UpdateTenantInput represents input for updating a tenant
type UpdateTenantInput struct {
	Name      *string                `json:"name"`
	Domain    *string                `json:"domain"`
	Subdomain *string                `json:"subdomain"`
	PlanID    *uuid.UUID             `json:"plan_id"`
	Status    *string                `json:"status"`
	Settings  map[string]interface{} `json:"settings"`
}

// TenantFilter represents filtering options for tenants
type TenantFilter struct {
	Status string  `json:"status"`
	PlanID *uuid.UUID `json:"plan_id"`
	Search string  `json:"search"`
}

// CreateTenant creates a new tenant
func (s *TenantService) CreateTenant(input CreateTenantInput) (*models.Tenant, error) {
	// Validate slug format
	if !isValidSlug(input.Slug) {
		return nil, fmt.Errorf("invalid slug format: must contain only lowercase letters, numbers, and hyphens")
	}

	tenant := &models.Tenant{
		Name:      input.Name,
		Slug:      input.Slug,
		Domain:    input.Domain,
		Subdomain: input.Subdomain,
		PlanID:    input.PlanID,
		Settings:  input.Settings,
		Status:    "active",
	}

	if err := s.db.Create(tenant).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, fmt.Errorf("tenant with slug '%s' already exists", input.Slug)
		}
		return nil, fmt.Errorf("failed to create tenant: %v", err)
	}

	// Create tenant-specific database schema
	if err := s.createTenantSchema(tenant.ID.String()); err != nil {
		// Rollback tenant creation if schema creation fails
		s.db.Delete(tenant)
		return nil, fmt.Errorf("failed to create tenant schema: %v", err)
	}

	return tenant, nil
}

// GetTenant retrieves a tenant by ID
func (s *TenantService) GetTenant(id uuid.UUID) (*models.Tenant, error) {
	var tenant models.Tenant
	err := s.db.Preload("Plan").First(&tenant, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("tenant not found")
		}
		return nil, fmt.Errorf("failed to get tenant: %v", err)
	}
	return &tenant, nil
}

// GetTenantBySlug retrieves a tenant by slug
func (s *TenantService) GetTenantBySlug(slug string) (*models.Tenant, error) {
	var tenant models.Tenant
	err := s.db.Preload("Plan").Where("slug = ?", slug).First(&tenant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("tenant not found")
		}
		return nil, fmt.Errorf("failed to get tenant: %v", err)
	}
	return &tenant, nil
}

// ListTenants retrieves tenants with filtering and pagination
func (s *TenantService) ListTenants(filter TenantFilter, offset, limit int) ([]*models.Tenant, int64, error) {
	query := s.db.Model(&models.Tenant{}).Preload("Plan")

	// Apply filters
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.PlanID != nil {
		query = query.Where("plan_id = ?", *filter.PlanID)
	}
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("name ILIKE ? OR slug ILIKE ?", search, search)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count tenants: %v", err)
	}

	// Get paginated results
	var tenants []*models.Tenant
	err := query.Offset(offset).Limit(limit).Find(&tenants).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list tenants: %v", err)
	}

	return tenants, total, nil
}

// UpdateTenant updates a tenant
func (s *TenantService) UpdateTenant(id uuid.UUID, input UpdateTenantInput) (*models.Tenant, error) {
	var tenant models.Tenant
	if err := s.db.First(&tenant, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("tenant not found")
		}
		return nil, fmt.Errorf("failed to find tenant: %v", err)
	}

	// Update fields
	if input.Name != nil {
		tenant.Name = *input.Name
	}
	if input.Domain != nil {
		tenant.Domain = input.Domain
	}
	if input.Subdomain != nil {
		tenant.Subdomain = input.Subdomain
	}
	if input.PlanID != nil {
		tenant.PlanID = input.PlanID
	}
	if input.Status != nil {
		tenant.Status = *input.Status
	}
	if input.Settings != nil {
		tenant.Settings = input.Settings
	}

	if err := s.db.Save(&tenant).Error; err != nil {
		return nil, fmt.Errorf("failed to update tenant: %v", err)
	}

	return &tenant, nil
}

// DeleteTenant soft deletes a tenant
func (s *TenantService) DeleteTenant(id uuid.UUID) error {
	var tenant models.Tenant
	if err := s.db.First(&tenant, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tenant not found")
		}
		return fmt.Errorf("failed to find tenant: %v", err)
	}

	// Soft delete the tenant
	if err := s.db.Delete(&tenant).Error; err != nil {
		return fmt.Errorf("failed to delete tenant: %v", err)
	}

	return nil
}

// SuspendTenant suspends a tenant
func (s *TenantService) SuspendTenant(id uuid.UUID) error {
	return s.updateTenantStatus(id, "suspended")
}

// ActivateTenant activates a tenant
func (s *TenantService) ActivateTenant(id uuid.UUID) error {
	return s.updateTenantStatus(id, "active")
}

// Helper methods

func (s *TenantService) updateTenantStatus(id uuid.UUID, status string) error {
	result := s.db.Model(&models.Tenant{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return fmt.Errorf("failed to update tenant status: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("tenant not found")
	}
	return nil
}

func (s *TenantService) createTenantSchema(tenantID string) error {
	schemaName := fmt.Sprintf("tenant_%s", strings.ReplaceAll(tenantID, "-", "_"))
	
	// Create schema
	if err := s.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName)).Error; err != nil {
		return err
	}

	// TODO: Create tenant-specific tables here
	// This could be done by reading a schema template and executing it
	
	return nil
}

func isValidSlug(slug string) bool {
	if len(slug) == 0 || len(slug) > 50 {
		return false
	}
	
	for _, r := range slug {
		if !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-') {
			return false
		}
	}
	
	// Cannot start or end with hyphen
	return slug[0] != '-' && slug[len(slug)-1] != '-'
}