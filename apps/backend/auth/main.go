package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared"
)

func main() {
	// Initialize database
	if err := shared.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize JWT token manager
	jwtSecret := getEnv("JWT_SECRET", "your-super-secret-jwt-key")
	initTokenManager(jwtSecret, "zplus-saas")

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Routes
	setupRoutes(app)

	// Start server
	port := getEnv("PORT", "8081")
	log.Printf("Auth service starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}