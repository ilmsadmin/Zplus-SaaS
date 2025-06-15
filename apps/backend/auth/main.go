package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth/handlers"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
				"code":  "SERVER_ERROR",
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Initialize handlers
	authHandler := handlers.NewAuthHandler()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "auth",
			"status":  "running",
			"message": "Authentication & RBAC Service",
			"version": "1.0.0",
		})
	})

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"timestamp": c.Context().Time(),
		})
	})

	// Authentication endpoints
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)
	app.Post("/refresh", authHandler.RefreshToken)

	// Development endpoint to see mock users
	app.Get("/users", authHandler.GetUsers)

	// Development endpoint to see active sessions
	app.Get("/sessions", authHandler.GetSessions)

	// Register endpoint (placeholder for future implementation)
	app.Post("/register", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Register endpoint - to be implemented",
			"code":    "NOT_IMPLEMENTED",
		})
	})

	log.Printf("Auth service starting on port 8081...")
	log.Fatal(app.Listen(":8081"))
}