package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/services"
)

// SubscriptionHandler handles subscription CRUD operations
type SubscriptionHandler struct {
	subscriptionService *services.SubscriptionService
}

// NewSubscriptionHandler creates a new subscription handler
func NewSubscriptionHandler(subscriptionService *services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

// GetSubscriptions retrieves all subscriptions with pagination and filtering
func (h *SubscriptionHandler) GetSubscriptions(c *fiber.Ctx) error {
	// Parse pagination params
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	if limit > 100 {
		limit = 100 // Max limit
	}
	offset := (page - 1) * limit

	// Parse filters
	filter := services.SubscriptionFilter{
		Status: c.Query("status"),
	}

	if tenantID := c.Query("tenant_id"); tenantID != "" {
		if id, err := uuid.Parse(tenantID); err == nil {
			filter.TenantID = id
		}
	}

	if planID := c.Query("plan_id"); planID != "" {
		if id, err := uuid.Parse(planID); err == nil {
			filter.PlanID = &id
		}
	}

	// Get subscriptions
	subscriptions, total, err := h.subscriptionService.ListSubscriptions(filter, offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve subscriptions",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": subscriptions,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetSubscription retrieves a single subscription by ID
func (h *SubscriptionHandler) GetSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	subscriptionID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid subscription ID",
			"message": "Subscription ID must be a valid UUID",
		})
	}

	subscription, err := h.subscriptionService.GetSubscription(subscriptionID)
	if err != nil {
		if err.Error() == "subscription not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Subscription not found",
				"message": "No subscription found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve subscription",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": subscription,
	})
}

// GetTenantSubscription retrieves the active subscription for a tenant
func (h *SubscriptionHandler) GetTenantSubscription(c *fiber.Ctx) error {
	tenantID := c.Params("tenant_id")
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid tenant ID",
			"message": "Tenant ID must be a valid UUID",
		})
	}

	subscription, err := h.subscriptionService.GetSubscriptionByTenant(tenantUUID)
	if err != nil {
		if err.Error() == "no active subscription found for tenant" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "No active subscription",
				"message": "No active subscription found for this tenant",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve subscription",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": subscription,
	})
}

// CreateSubscription creates a new subscription
func (h *SubscriptionHandler) CreateSubscription(c *fiber.Ctx) error {
	var input services.CreateSubscriptionInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if input.TenantID == uuid.Nil || input.PlanID == uuid.Nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Missing required fields",
			"message": "Tenant ID and Plan ID are required",
		})
	}

	subscription, err := h.subscriptionService.CreateSubscription(input)
	if err != nil {
		if err.Error() == "tenant already has an active subscription" {
			return c.Status(409).JSON(fiber.Map{
				"error":   "Subscription conflict",
				"message": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create subscription",
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data":    subscription,
		"message": "Subscription created successfully",
	})
}

// UpdateSubscription updates an existing subscription
func (h *SubscriptionHandler) UpdateSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	subscriptionID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid subscription ID",
			"message": "Subscription ID must be a valid UUID",
		})
	}

	var input services.UpdateSubscriptionInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	subscription, err := h.subscriptionService.UpdateSubscription(subscriptionID, input)
	if err != nil {
		if err.Error() == "subscription not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Subscription not found",
				"message": "No subscription found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update subscription",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":    subscription,
		"message": "Subscription updated successfully",
	})
}

// CancelSubscription cancels a subscription
func (h *SubscriptionHandler) CancelSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	subscriptionID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid subscription ID",
			"message": "Subscription ID must be a valid UUID",
		})
	}

	subscription, err := h.subscriptionService.CancelSubscription(subscriptionID)
	if err != nil {
		if err.Error() == "subscription not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Subscription not found",
				"message": "No subscription found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to cancel subscription",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":    subscription,
		"message": "Subscription cancelled successfully",
	})
}

// GetSubscriptionStats retrieves subscription statistics
func (h *SubscriptionHandler) GetSubscriptionStats(c *fiber.Ctx) error {
	stats, err := h.subscriptionService.GetSubscriptionStats()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve subscription statistics",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": stats,
	})
}