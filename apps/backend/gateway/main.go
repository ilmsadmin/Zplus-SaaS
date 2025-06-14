package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/generated"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/middleware"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/resolver"
)

func main() {
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
			},
		})
	})

	// Health endpoint for monitoring
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"timestamp": "2024-01-01T00:00:00Z",
		})
	})

	// Create GraphQL server
	gqlResolver := resolver.NewResolver()
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
	setupRESTRoutes(app)

	log.Printf("🚀 Zplus SaaS Gateway starting on :8080")
	log.Printf("📊 GraphQL endpoint: http://localhost:8080/graphql")
	log.Printf("🛝 GraphQL Playground: http://localhost:8080/playground")
	log.Fatal(app.Listen(":8080"))
}

// setupRESTRoutes configures REST API endpoints for backward compatibility
func setupRESTRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	
	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"api_version": "1.0",
		})
	})
	
	// User endpoints
	users := api.Group("/users")
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
	
	// Tenant endpoints
	tenants := api.Group("/tenants")
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