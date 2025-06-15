package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gorm.io/gorm"

	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/generated"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/handlers"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/middleware"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/resolver"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/services"
	"github.com/ilmsadmin/Zplus-SaaS/pkg/database"
)

// getEnv returns environment variable or default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt returns environment variable as int or default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func main() {
	// Initialize database connection
	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// Global middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${latency}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Tenant-ID",
		AllowCredentials: true,
	}))

	// Multi-tenant middleware
	app.Use(middleware.TenantMiddleware())
	app.Use(middleware.AuthMiddleware())
	app.Use(middleware.GraphQLContextMiddleware())

	// Health check endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "Zplus SaaS API Gateway",
			"status":  "running",
			"version": "1.0.0",
			"message": "GraphQL-first Multi-tenant API Gateway",
			"endpoints": fiber.Map{
				"graphql":    "/graphql",
				"playground": "/playground",
				"health":     "/health",
				"api":        "/api/v1",
			},
		})
	})

	// Health endpoint for monitoring
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"timestamp": "2024-01-01T00:00:00Z",
		})
	})

	// Create GraphQL server with database integration
	gqlResolver := resolver.NewResolver()
	gqlResolver.SetDatabase(db)

	gqlServer := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: gqlResolver,
		}),
	)

	// GraphQL endpoint with context injection
	app.All("/graphql", func(c *fiber.Ctx) error {
		// Get request context from middleware
		requestCtx := middleware.GetRequestContext(c)

		// Create GraphQL context with request context
		ctx := context.WithValue(c.Context(), "request_context", requestCtx)

		// Adapt Fiber to net/http for GraphQL handler
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the context with request information
			r = r.WithContext(ctx)
			gqlServer.ServeHTTP(w, r)
		})(c.Context())

		return nil
	})

	// GraphQL Playground for development
	app.Get("/playground", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(
			playground.Handler("GraphQL Playground", "/graphql"),
		)(c.Context())
		return nil
	})

	// REST API endpoints for backward compatibility
	setupRESTRoutes(app, db, gqlResolver)

	// Get port from environment variable
	port := getEnv("GATEWAY_PORT", "8000")

	log.Printf("🚀 Zplus SaaS Gateway starting on :%s", port)
	log.Printf("📊 GraphQL endpoint: http://localhost:%s/graphql", port)
	log.Printf("🛝 GraphQL Playground: http://localhost:%s/playground", port)
	log.Printf("🔗 REST API: http://localhost:%s/api/v1", port)
	log.Fatal(app.Listen(":" + port))
}

// initializeDatabase sets up the database connection
func initializeDatabase() (*gorm.DB, error) {
	// Get database configuration from environment variables
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvInt("DB_PORT", 5432),
		Username: getEnv("DB_USERNAME", "zplus_user"),
		Password: getEnv("DB_PASSWORD", "zplus_password"),
		Database: getEnv("DB_DATABASE", "zplus_saas"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	db, err := database.Connect(dbConfig)
	if err != nil {
		return nil, err
	}

	log.Printf("🗄️  Database connected successfully")
	return db, nil
}

// setupRESTRoutes configures REST API endpoints for backward compatibility
func setupRESTRoutes(app *fiber.App, db *gorm.DB, gqlResolver *resolver.Resolver) {
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      "healthy",
			"api_version": "1.0",
		})
	})

	// User endpoints
	users := api.Group("/users")
	userHandler := handlers.NewUserHandler(gqlResolver.GetUserService)
	users.Get("/me", func(c *fiber.Ctx) error {
		userCtx := middleware.GetUserContext(c)
		if userCtx == nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authentication required",
			})
		}

		return c.JSON(fiber.Map{
			"id":         userCtx.ID,
			"email":      userCtx.Email,
			"first_name": userCtx.FirstName,
			"last_name":  userCtx.LastName,
			"tenant_id":  userCtx.TenantID,
			"roles":      userCtx.Roles,
		})
	})
	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
	users.Post("/:id/password", userHandler.ChangePassword)
	users.Post("/:id/roles", userHandler.AssignRoles)

	// Tenant endpoints (system admin only)
	tenants := api.Group("/tenants")
	tenantHandler := handlers.NewTenantHandler(services.NewTenantService(db))
	tenants.Get("/current", func(c *fiber.Ctx) error {
		tenantCtx := middleware.GetTenantContext(c)
		if tenantCtx == nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Tenant context not available",
			})
		}

		return c.JSON(fiber.Map{
			"id":       tenantCtx.ID,
			"slug":     tenantCtx.Slug,
			"name":     tenantCtx.Name,
			"status":   tenantCtx.Status,
			"plan_id":  tenantCtx.PlanID,
			"features": tenantCtx.Features,
		})
	})
	tenants.Get("/", tenantHandler.GetTenants)
	tenants.Get("/:id", tenantHandler.GetTenant)
	tenants.Post("/", tenantHandler.CreateTenant)
	tenants.Put("/:id", tenantHandler.UpdateTenant)
	tenants.Delete("/:id", tenantHandler.DeleteTenant)
	tenants.Post("/:id/suspend", tenantHandler.SuspendTenant)
	tenants.Post("/:id/activate", tenantHandler.ActivateTenant)

	// Plan endpoints (system admin only)
	plans := api.Group("/plans")
	planHandler := handlers.NewPlanHandler(services.NewPlanService(db))
	plans.Get("/", planHandler.GetPlans)
	plans.Get("/:id", planHandler.GetPlan)
	plans.Post("/", planHandler.CreatePlan)
	plans.Put("/:id", planHandler.UpdatePlan)
	plans.Delete("/:id", planHandler.DeletePlan)
	plans.Get("/:id/usage", planHandler.GetPlanUsage)

	// Subscription endpoints
	subscriptions := api.Group("/subscriptions")
	subscriptionHandler := handlers.NewSubscriptionHandler(services.NewSubscriptionService(db))
	subscriptions.Get("/", subscriptionHandler.GetSubscriptions)
	subscriptions.Get("/:id", subscriptionHandler.GetSubscription)
	subscriptions.Get("/tenant/:tenant_id", subscriptionHandler.GetTenantSubscription)
	subscriptions.Post("/", subscriptionHandler.CreateSubscription)
	subscriptions.Put("/:id", subscriptionHandler.UpdateSubscription)
	subscriptions.Post("/:id/cancel", subscriptionHandler.CancelSubscription)
	subscriptions.Get("/stats", subscriptionHandler.GetSubscriptionStats)
}

// errorHandler handles Fiber errors
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
		"code":  code,
	})
}
