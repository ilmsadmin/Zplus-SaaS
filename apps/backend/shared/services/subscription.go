package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/models"
)

// SubscriptionService handles CRUD operations for subscriptions
type SubscriptionService struct {
	db *gorm.DB
}

// NewSubscriptionService creates a new subscription service
func NewSubscriptionService(db *gorm.DB) *SubscriptionService {
	return &SubscriptionService{db: db}
}

// CreateSubscriptionInput represents input for creating a subscription
type CreateSubscriptionInput struct {
	TenantID      uuid.UUID              `json:"tenant_id" validate:"required"`
	PlanID        uuid.UUID              `json:"plan_id" validate:"required"`
	StartDate     *time.Time             `json:"start_date"`
	TrialEndDate  *time.Time             `json:"trial_end_date"`
	BillingCycle  string                 `json:"billing_cycle"` // monthly, yearly
	Amount        *float64               `json:"amount"`
	Currency      string                 `json:"currency"`
	Metadata      map[string]interface{} `json:"metadata"`
}

// UpdateSubscriptionInput represents input for updating a subscription
type UpdateSubscriptionInput struct {
	PlanID       *uuid.UUID             `json:"plan_id"`
	Status       *string                `json:"status"`
	EndDate      *time.Time             `json:"end_date"`
	BillingCycle *string                `json:"billing_cycle"`
	Amount       *float64               `json:"amount"`
	Currency     *string                `json:"currency"`
	Metadata     map[string]interface{} `json:"metadata"`
}

// SubscriptionFilter represents filtering options for subscriptions
type SubscriptionFilter struct {
	TenantID uuid.UUID `json:"tenant_id"`
	PlanID   *uuid.UUID `json:"plan_id"`
	Status   string    `json:"status"`
}

// CreateSubscription creates a new subscription
func (s *SubscriptionService) CreateSubscription(input CreateSubscriptionInput) (*models.Subscription, error) {
	// Validate tenant exists
	var tenant models.Tenant
	if err := s.db.First(&tenant, input.TenantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("tenant not found")
		}
		return nil, fmt.Errorf("failed to verify tenant: %v", err)
	}

	// Validate plan exists
	var plan models.Plan
	if err := s.db.First(&plan, input.PlanID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("plan not found")
		}
		return nil, fmt.Errorf("failed to verify plan: %v", err)
	}

	// Check if tenant already has an active subscription
	var existingCount int64
	if err := s.db.Model(&models.Subscription{}).
		Where("tenant_id = ? AND status IN ('active', 'trial')", input.TenantID).
		Count(&existingCount).Error; err != nil {
		return nil, fmt.Errorf("failed to check existing subscriptions: %v", err)
	}

	if existingCount > 0 {
		return nil, fmt.Errorf("tenant already has an active subscription")
	}

	// Set default values
	startDate := time.Now()
	if input.StartDate != nil {
		startDate = *input.StartDate
	}

	billingCycle := "monthly"
	if input.BillingCycle != "" {
		billingCycle = input.BillingCycle
	}

	amount := plan.Price
	if input.Amount != nil {
		amount = *input.Amount
	}

	currency := "USD"
	if input.Currency != "" {
		currency = input.Currency
	}

	status := "active"
	if input.TrialEndDate != nil {
		status = "trial"
	}

	subscription := &models.Subscription{
		TenantID:     input.TenantID,
		PlanID:       input.PlanID,
		Status:       status,
		StartDate:    startDate,
		TrialEndDate: input.TrialEndDate,
		BillingCycle: billingCycle,
		Amount:       amount,
		Currency:     currency,
		Metadata:     input.Metadata,
	}

	if err := s.db.Create(subscription).Error; err != nil {
		return nil, fmt.Errorf("failed to create subscription: %v", err)
	}

	// Update tenant's plan_id
	if err := s.db.Model(&tenant).Update("plan_id", input.PlanID).Error; err != nil {
		return nil, fmt.Errorf("failed to update tenant plan: %v", err)
	}

	// Load relationships
	if err := s.db.Preload("Tenant").Preload("Plan").First(subscription, subscription.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to load subscription details: %v", err)
	}

	return subscription, nil
}

// GetSubscription retrieves a subscription by ID
func (s *SubscriptionService) GetSubscription(id uuid.UUID) (*models.Subscription, error) {
	var subscription models.Subscription
	err := s.db.Preload("Tenant").Preload("Plan").First(&subscription, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("subscription not found")
		}
		return nil, fmt.Errorf("failed to get subscription: %v", err)
	}
	return &subscription, nil
}

// GetSubscriptionByTenant retrieves the active subscription for a tenant
func (s *SubscriptionService) GetSubscriptionByTenant(tenantID uuid.UUID) (*models.Subscription, error) {
	var subscription models.Subscription
	err := s.db.Preload("Tenant").Preload("Plan").
		Where("tenant_id = ? AND status IN ('active', 'trial')", tenantID).
		First(&subscription).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no active subscription found for tenant")
		}
		return nil, fmt.Errorf("failed to get subscription: %v", err)
	}
	return &subscription, nil
}

// ListSubscriptions retrieves subscriptions with filtering and pagination
func (s *SubscriptionService) ListSubscriptions(filter SubscriptionFilter, offset, limit int) ([]*models.Subscription, int64, error) {
	query := s.db.Model(&models.Subscription{}).Preload("Tenant").Preload("Plan")

	// Apply filters
	if filter.TenantID != uuid.Nil {
		query = query.Where("tenant_id = ?", filter.TenantID)
	}
	if filter.PlanID != nil {
		query = query.Where("plan_id = ?", *filter.PlanID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count subscriptions: %v", err)
	}

	// Get paginated results
	var subscriptions []*models.Subscription
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&subscriptions).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list subscriptions: %v", err)
	}

	return subscriptions, total, nil
}

// UpdateSubscription updates a subscription
func (s *SubscriptionService) UpdateSubscription(id uuid.UUID, input UpdateSubscriptionInput) (*models.Subscription, error) {
	var subscription models.Subscription
	if err := s.db.First(&subscription, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("subscription not found")
		}
		return nil, fmt.Errorf("failed to find subscription: %v", err)
	}

	// Update fields
	if input.PlanID != nil {
		// Validate new plan exists
		var plan models.Plan
		if err := s.db.First(&plan, *input.PlanID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("plan not found")
			}
			return nil, fmt.Errorf("failed to verify plan: %v", err)
		}
		subscription.PlanID = *input.PlanID
		// Update amount if not explicitly provided
		if input.Amount == nil {
			subscription.Amount = plan.Price
		}
	}
	if input.Status != nil {
		subscription.Status = *input.Status
		if *input.Status == "cancelled" && subscription.EndDate == nil {
			now := time.Now()
			subscription.EndDate = &now
		}
	}
	if input.EndDate != nil {
		subscription.EndDate = input.EndDate
	}
	if input.BillingCycle != nil {
		subscription.BillingCycle = *input.BillingCycle
	}
	if input.Amount != nil {
		subscription.Amount = *input.Amount
	}
	if input.Currency != nil {
		subscription.Currency = *input.Currency
	}
	if input.Metadata != nil {
		subscription.Metadata = input.Metadata
	}

	if err := s.db.Save(&subscription).Error; err != nil {
		return nil, fmt.Errorf("failed to update subscription: %v", err)
	}

	// Load relationships
	if err := s.db.Preload("Tenant").Preload("Plan").First(&subscription, subscription.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to load subscription details: %v", err)
	}

	return &subscription, nil
}

// CancelSubscription cancels a subscription
func (s *SubscriptionService) CancelSubscription(id uuid.UUID) (*models.Subscription, error) {
	now := time.Now()
	return s.UpdateSubscription(id, UpdateSubscriptionInput{
		Status:  stringPtr("cancelled"),
		EndDate: &now,
	})
}

// RenewSubscription renews a subscription
func (s *SubscriptionService) RenewSubscription(id uuid.UUID, endDate *time.Time) (*models.Subscription, error) {
	return s.UpdateSubscription(id, UpdateSubscriptionInput{
		Status:  stringPtr("active"),
		EndDate: endDate,
	})
}

// DeleteSubscription hard deletes a subscription (use with caution)
func (s *SubscriptionService) DeleteSubscription(id uuid.UUID) error {
	var subscription models.Subscription
	if err := s.db.First(&subscription, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("subscription not found")
		}
		return fmt.Errorf("failed to find subscription: %v", err)
	}

	// Hard delete the subscription
	if err := s.db.Unscoped().Delete(&subscription).Error; err != nil {
		return fmt.Errorf("failed to delete subscription: %v", err)
	}

	return nil
}

// GetSubscriptionStats returns subscription statistics
func (s *SubscriptionService) GetSubscriptionStats() (map[string]interface{}, error) {
	var totalActive, totalTrial, totalCancelled, totalExpired int64

	// Count active subscriptions
	if err := s.db.Model(&models.Subscription{}).Where("status = 'active'").Count(&totalActive).Error; err != nil {
		return nil, fmt.Errorf("failed to count active subscriptions: %v", err)
	}

	// Count trial subscriptions
	if err := s.db.Model(&models.Subscription{}).Where("status = 'trial'").Count(&totalTrial).Error; err != nil {
		return nil, fmt.Errorf("failed to count trial subscriptions: %v", err)
	}

	// Count cancelled subscriptions
	if err := s.db.Model(&models.Subscription{}).Where("status = 'cancelled'").Count(&totalCancelled).Error; err != nil {
		return nil, fmt.Errorf("failed to count cancelled subscriptions: %v", err)
	}

	// Count expired subscriptions
	if err := s.db.Model(&models.Subscription{}).Where("status = 'expired'").Count(&totalExpired).Error; err != nil {
		return nil, fmt.Errorf("failed to count expired subscriptions: %v", err)
	}

	// Calculate total monthly revenue (active subscriptions with monthly billing)
	var monthlyRevenue float64
	if err := s.db.Model(&models.Subscription{}).
		Where("status = 'active' AND billing_cycle = 'monthly'").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&monthlyRevenue).Error; err != nil {
		return nil, fmt.Errorf("failed to calculate monthly revenue: %v", err)
	}

	return map[string]interface{}{
		"total_active":     totalActive,
		"total_trial":      totalTrial,
		"total_cancelled":  totalCancelled,
		"total_expired":    totalExpired,
		"monthly_revenue":  monthlyRevenue,
	}, nil
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}