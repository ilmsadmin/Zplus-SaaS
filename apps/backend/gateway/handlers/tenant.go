package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/services"
)

// TenantHandler handles tenant CRUD operations
type TenantHandler struct {
	tenantService *services.TenantService
}

// NewTenantHandler creates a new tenant handler
func NewTenantHandler(tenantService *services.TenantService) *TenantHandler {
	return &TenantHandler{
		tenantService: tenantService,
	}
}

// GetTenants retrieves all tenants with pagination and filtering
func (h *TenantHandler) GetTenants(c *fiber.Ctx) error {
	// Parse pagination params
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	if limit > 100 {
		limit = 100 // Max limit
	}
	offset := (page - 1) * limit

	// Parse filters
	filter := services.TenantFilter{
		Status: c.Query("status"),
		Search: c.Query("search"),
	}

	if planID := c.Query("plan_id"); planID != "" {
		if id, err := uuid.Parse(planID); err == nil {
			filter.PlanID = &id
		}
	}

	// Get tenants
	tenants, total, err := h.tenantService.ListTenants(filter, offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve tenants",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": tenants,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetTenant retrieves a single tenant by ID
func (h *TenantHandler) GetTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	tenantID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid tenant ID",
			"message": "Tenant ID must be a valid UUID",
		})
	}

	tenant, err := h.tenantService.GetTenant(tenantID)
	if err != nil {
		if err.Error() == "tenant not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Tenant not found",
				"message": "No tenant found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve tenant",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": tenant,
	})
}

// CreateTenant creates a new tenant
func (h *TenantHandler) CreateTenant(c *fiber.Ctx) error {
	var input services.CreateTenantInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if input.Name == "" || input.Slug == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Missing required fields",
			"message": "Name and slug are required",
		})
	}

	tenant, err := h.tenantService.CreateTenant(input)
	if err != nil {
		if err.Error() == "tenant with slug '"+input.Slug+"' already exists" {
			return c.Status(409).JSON(fiber.Map{
				"error":   "Tenant already exists",
				"message": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create tenant",
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data":    tenant,
		"message": "Tenant created successfully",
	})
}

// UpdateTenant updates an existing tenant
func (h *TenantHandler) UpdateTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	tenantID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid tenant ID",
			"message": "Tenant ID must be a valid UUID",
		})
	}

	var input services.UpdateTenantInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	tenant, err := h.tenantService.UpdateTenant(tenantID, input)
	if err != nil {
		if err.Error() == "tenant not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Tenant not found",
				"message": "No tenant found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update tenant",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":    tenant,
		"message": "Tenant updated successfully",
	})
}

// DeleteTenant deletes a tenant
func (h *TenantHandler) DeleteTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	tenantID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid tenant ID",
			"message": "Tenant ID must be a valid UUID",
		})
	}

	err = h.tenantService.DeleteTenant(tenantID)
	if err != nil {
		if err.Error() == "tenant not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Tenant not found",
				"message": "No tenant found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete tenant",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tenant deleted successfully",
	})
}

// SuspendTenant suspends a tenant
func (h *TenantHandler) SuspendTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	tenantID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid tenant ID",
			"message": "Tenant ID must be a valid UUID",
		})
	}

	err = h.tenantService.SuspendTenant(tenantID)
	if err != nil {
		if err.Error() == "tenant not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Tenant not found",
				"message": "No tenant found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to suspend tenant",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tenant suspended successfully",
	})
}

// ActivateTenant activates a tenant
func (h *TenantHandler) ActivateTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	tenantID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid tenant ID",
			"message": "Tenant ID must be a valid UUID",
		})
	}

	err = h.tenantService.ActivateTenant(tenantID)
	if err != nil {
		if err.Error() == "tenant not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Tenant not found",
				"message": "No tenant found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to activate tenant",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tenant activated successfully",
	})
}