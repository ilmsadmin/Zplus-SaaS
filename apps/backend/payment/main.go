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

	log.Fatal(app.Listen(":8083"))
}