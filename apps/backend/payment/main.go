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
			"service": "payment",
			"status":  "running",
			"message": "Payment & Subscription Service",
		})
	})

	app.Post("/subscriptions", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Subscription management - to be implemented",
		})
	})

	log.Printf("Payment service starting on port %s...", getEnv("PAYMENT_PORT", "8003"))
	log.Fatal(app.Listen(":" + getEnv("PAYMENT_PORT", "8003")))
}
