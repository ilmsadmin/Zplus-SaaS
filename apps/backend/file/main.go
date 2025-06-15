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
			"service": "file",
			"status":  "running",
			"message": "File Management Service",
		})
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "File upload endpoint - to be implemented",
		})
	})

	app.Get("/download/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "File download endpoint - to be implemented",
		})
	})

	log.Printf("File service starting on port %s...", getEnv("FILE_PORT", "8002"))
	log.Fatal(app.Listen(":" + getEnv("FILE_PORT", "8002")))
}
