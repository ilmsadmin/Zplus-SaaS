package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/services"
)

// PlanHandler handles subscription plan CRUD operations
type PlanHandler struct {
	planService *services.PlanService
}

// NewPlanHandler creates a new plan handler
func NewPlanHandler(planService *services.PlanService) *PlanHandler {
	return &PlanHandler{
		planService: planService,
	}
}

// GetPlans retrieves all plans with pagination and filtering
func (h *PlanHandler) GetPlans(c *fiber.Ctx) error {
	// Parse pagination params
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	if limit > 100 {
		limit = 100 // Max limit
	}
	offset := (page - 1) * limit

	// Parse filters
	filter := services.PlanFilter{
		Search: c.Query("search"),
	}

	if minPrice := c.Query("min_price"); minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
			filter.MinPrice = &price
		}
	}

	if maxPrice := c.Query("max_price"); maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			filter.MaxPrice = &price
		}
	}

	// Get plans
	plans, total, err := h.planService.ListPlans(filter, offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve plans",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": plans,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetPlan retrieves a single plan by ID
func (h *PlanHandler) GetPlan(c *fiber.Ctx) error {
	id := c.Params("id")
	planID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid plan ID",
			"message": "Plan ID must be a valid UUID",
		})
	}

	plan, err := h.planService.GetPlan(planID)
	if err != nil {
		if err.Error() == "plan not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Plan not found",
				"message": "No plan found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve plan",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": plan,
	})
}

// CreatePlan creates a new plan
func (h *PlanHandler) CreatePlan(c *fiber.Ctx) error {
	var input services.CreatePlanInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	// Validate required fields
	if input.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Missing required fields",
			"message": "Name is required",
		})
	}

	plan, err := h.planService.CreatePlan(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create plan",
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data":    plan,
		"message": "Plan created successfully",
	})
}

// UpdatePlan updates an existing plan
func (h *PlanHandler) UpdatePlan(c *fiber.Ctx) error {
	id := c.Params("id")
	planID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid plan ID",
			"message": "Plan ID must be a valid UUID",
		})
	}

	var input services.UpdatePlanInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": "Please provide valid JSON data",
		})
	}

	plan, err := h.planService.UpdatePlan(planID, input)
	if err != nil {
		if err.Error() == "plan not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Plan not found",
				"message": "No plan found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update plan",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":    plan,
		"message": "Plan updated successfully",
	})
}

// DeletePlan deletes a plan
func (h *PlanHandler) DeletePlan(c *fiber.Ctx) error {
	id := c.Params("id")
	planID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid plan ID",
			"message": "Plan ID must be a valid UUID",
		})
	}

	err = h.planService.DeletePlan(planID)
	if err != nil {
		if err.Error() == "plan not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Plan not found",
				"message": "No plan found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete plan",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Plan deleted successfully",
	})
}

// GetPlanUsage retrieves plan usage statistics
func (h *PlanHandler) GetPlanUsage(c *fiber.Ctx) error {
	id := c.Params("id")
	planID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid plan ID",
			"message": "Plan ID must be a valid UUID",
		})
	}

	usage, err := h.planService.GetPlanUsage(planID)
	if err != nil {
		if err.Error() == "plan not found" {
			return c.Status(404).JSON(fiber.Map{
				"error":   "Plan not found",
				"message": "No plan found with the specified ID",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to retrieve plan usage",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": usage,
	})
}