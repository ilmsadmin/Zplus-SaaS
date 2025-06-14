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

	log.Fatal(app.Listen(":8086"))
}