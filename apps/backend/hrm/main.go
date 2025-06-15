package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// getEnv returns environment variable or default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "hrm",
			"status":  "running",
			"message": "Human Resource Management Service",
		})
	})

	app.Get("/employees", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Employee list endpoint - to be implemented",
		})
	})

	log.Printf("HRM service starting on port %s...", getEnv("HRM_PORT", "8005"))
	log.Fatal(app.Listen(":" + getEnv("HRM_PORT", "8005")))
}
