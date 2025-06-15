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
			"service": "pos",
			"status":  "running",
			"message": "Point of Sale Service",
		})
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Product list endpoint - to be implemented",
		})
	})

	log.Printf("POS service starting on port %s...", getEnv("POS_PORT", "8006"))
	log.Fatal(app.Listen(":" + getEnv("POS_PORT", "8006")))
}
