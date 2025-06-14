package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

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

	log.Fatal(app.Listen(":8085"))
}