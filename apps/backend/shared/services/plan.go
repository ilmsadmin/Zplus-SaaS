package services

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/models"
)

// PlanService handles CRUD operations for subscription plans
type PlanService struct {
	db *gorm.DB
}

// NewPlanService creates a new plan service
func NewPlanService(db *gorm.DB) *PlanService {
	return &PlanService{db: db}
}

// CreatePlanInput represents input for creating a plan
type CreatePlanInput struct {
	Name        string                 `json:"name" validate:"required"`
	Description *string                `json:"description"`
	Price       float64                `json:"price" validate:"min=0"`
	Features    map[string]interface{} `json:"features"`
	MaxUsers    *int                   `json:"max_users"`
	MaxStorage  *int64                 `json:"max_storage"`
}

// UpdatePlanInput represents input for updating a plan
type UpdatePlanInput struct {
	Name        *string                `json:"name"`
	Description *string                `json:"description"`
	Price       *float64               `json:"price"`
	Features    map[string]interface{} `json:"features"`
	MaxUsers    *int                   `json:"max_users"`
	MaxStorage  *int64                 `json:"max_storage"`
}

// PlanFilter represents filtering options for plans
type PlanFilter struct {
	MinPrice *float64 `json:"min_price"`
	MaxPrice *float64 `json:"max_price"`
	Search   string   `json:"search"`
}

// CreatePlan creates a new subscription plan
func (s *PlanService) CreatePlan(input CreatePlanInput) (*models.Plan, error) {
	plan := &models.Plan{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Features:    input.Features,
		MaxUsers:    input.MaxUsers,
		MaxStorage:  input.MaxStorage,
	}

	if err := s.db.Create(plan).Error; err != nil {
		return nil, fmt.Errorf("failed to create plan: %v", err)
	}

	return plan, nil
}

// GetPlan retrieves a plan by ID
func (s *PlanService) GetPlan(id uuid.UUID) (*models.Plan, error) {
	var plan models.Plan
	err := s.db.First(&plan, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("plan not found")
		}
		return nil, fmt.Errorf("failed to get plan: %v", err)
	}
	return &plan, nil
}

// ListPlans retrieves plans with filtering and pagination
func (s *PlanService) ListPlans(filter PlanFilter, offset, limit int) ([]*models.Plan, int64, error) {
	query := s.db.Model(&models.Plan{})

	// Apply filters
	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("name ILIKE ? OR description ILIKE ?", search, search)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count plans: %v", err)
	}

	// Get paginated results
	var plans []*models.Plan
	err := query.Offset(offset).Limit(limit).Find(&plans).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list plans: %v", err)
	}

	return plans, total, nil
}

// UpdatePlan updates a plan
func (s *PlanService) UpdatePlan(id uuid.UUID, input UpdatePlanInput) (*models.Plan, error) {
	var plan models.Plan
	if err := s.db.First(&plan, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("plan not found")
		}
		return nil, fmt.Errorf("failed to find plan: %v", err)
	}

	// Update fields
	if input.Name != nil {
		plan.Name = *input.Name
	}
	if input.Description != nil {
		plan.Description = input.Description
	}
	if input.Price != nil {
		plan.Price = *input.Price
	}
	if input.Features != nil {
		plan.Features = input.Features
	}
	if input.MaxUsers != nil {
		plan.MaxUsers = input.MaxUsers
	}
	if input.MaxStorage != nil {
		plan.MaxStorage = input.MaxStorage
	}

	if err := s.db.Save(&plan).Error; err != nil {
		return nil, fmt.Errorf("failed to update plan: %v", err)
	}

	return &plan, nil
}

// DeletePlan soft deletes a plan
func (s *PlanService) DeletePlan(id uuid.UUID) error {
	var plan models.Plan
	if err := s.db.First(&plan, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("plan not found")
		}
		return fmt.Errorf("failed to find plan: %v", err)
	}

	// Check if plan is being used by any tenants
	var tenantCount int64
	if err := s.db.Model(&models.Tenant{}).Where("plan_id = ?", id).Count(&tenantCount).Error; err != nil {
		return fmt.Errorf("failed to check plan usage: %v", err)
	}

	if tenantCount > 0 {
		return fmt.Errorf("cannot delete plan: it is currently being used by %d tenant(s)", tenantCount)
	}

	// Soft delete the plan
	if err := s.db.Delete(&plan).Error; err != nil {
		return fmt.Errorf("failed to delete plan: %v", err)
	}

	return nil
}

// GetPlanUsage returns usage statistics for a plan
func (s *PlanService) GetPlanUsage(id uuid.UUID) (map[string]interface{}, error) {
	var plan models.Plan
	if err := s.db.First(&plan, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("plan not found")
		}
		return nil, fmt.Errorf("failed to find plan: %v", err)
	}

	// Count active tenants using this plan
	var activeTenantCount int64
	if err := s.db.Model(&models.Tenant{}).Where("plan_id = ? AND status = 'active'", id).Count(&activeTenantCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count active tenants: %v", err)
	}

	// Count total tenants using this plan
	var totalTenantCount int64
	if err := s.db.Model(&models.Tenant{}).Where("plan_id = ?", id).Count(&totalTenantCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count total tenants: %v", err)
	}

	// Count active subscriptions
	var activeSubscriptionCount int64
	if err := s.db.Model(&models.Subscription{}).Where("plan_id = ? AND status = 'active'", id).Count(&activeSubscriptionCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count active subscriptions: %v", err)
	}

	return map[string]interface{}{
		"plan_id":                   id,
		"plan_name":                 plan.Name,
		"active_tenants":            activeTenantCount,
		"total_tenants":             totalTenantCount,
		"active_subscriptions":      activeSubscriptionCount,
	}, nil
}