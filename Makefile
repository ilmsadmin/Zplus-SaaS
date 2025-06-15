# Zplus SaaS - Makefile
.PHONY: help setup install start stop clean build test

# Default target
help:
	@echo "🚀 Zplus SaaS - Available Commands"
	@echo "=================================="
	@echo ""
	@echo "Setup & Installation:"
	@echo "  make setup     - Run complete project setup"
	@echo "  make install   - Install all dependencies"
	@echo ""
	@echo "Development:"
	@echo "  make start     - Start all services (databases + backend + frontend)"
	@echo "  make backend   - Start backend services only"
	@echo "  make frontend  - Start frontend only"
	@echo "  make db        - Start database services only"
	@echo ""
	@echo "Management:"
	@echo "  make stop      - Stop all services"
	@echo "  make clean     - Clean build artifacts and dependencies"
	@echo "  make restart   - Restart all services"
	@echo ""
	@echo "Build & Test:"
	@echo "  make build     - Build all services"
	@echo "  make test      - Run all tests"
	@echo ""
	@echo "Database:"
	@echo "  make db-up     - Start database services"
	@echo "  make db-down   - Stop database services"
	@echo "  make db-reset  - Reset database (⚠️  destructive)"
	@echo ""

# Setup and installation
setup:
	@echo "🔧 Running project setup..."
	@./setup.sh

install:
	@echo "📦 Installing dependencies..."
	@cd pkg && go mod tidy
	@cd apps/backend/shared && go mod tidy
	@cd apps/backend/gateway && go mod tidy
	@cd apps/backend/auth && go mod tidy
	@cd apps/backend/file && go mod tidy
	@cd apps/backend/payment && go mod tidy
	@cd apps/backend/crm && go mod tidy
	@cd apps/backend/hrm && go mod tidy
	@cd apps/backend/pos && go mod tidy
	@cd apps/frontend/web && npm install
	@if [ -d "apps/frontend/ui" ]; then cd apps/frontend/ui && npm install; fi
	@echo "✅ Dependencies installed"

# Development commands
start:
	@echo "🚀 Starting all services..."
	@./run-all.sh

backend:
	@echo "🔧 Starting backend services..."
	@./run-backend.sh

frontend:
	@echo "🌐 Starting frontend..."
	@./run-frontend.sh

db:
	@echo "🗄️ Starting database services..."
	@cd infra/docker && docker-compose up -d postgres mongodb redis

# Database management
db-up:
	@echo "🗄️ Starting database services..."
	@cd infra/docker && docker-compose up -d postgres mongodb redis

db-down:
	@echo "🛑 Stopping database services..."
	@cd infra/docker && docker-compose down

db-reset:
	@echo "⚠️  Resetting database (this will delete all data)..."
	@read -p "Are you sure? Type 'yes' to continue: " confirm && [ "$$confirm" = "yes" ] || exit 1
	@cd infra/docker && docker-compose down -v
	@cd infra/docker && docker-compose up -d postgres mongodb redis
	@echo "✅ Database reset completed"

# Management commands
stop:
	@echo "🛑 Stopping all services..."
	@./stop-all.sh

restart: stop start

clean:
	@echo "🧹 Cleaning build artifacts..."
	@find . -name "node_modules" -type d -exec rm -rf {} + 2>/dev/null || true
	@find . -name "*.pid" -type f -delete 2>/dev/null || true
	@find . -name ".next" -type d -exec rm -rf {} + 2>/dev/null || true
	@find . -name "dist" -type d -exec rm -rf {} + 2>/dev/null || true
	@go clean -cache -modcache -testcache
	@echo "✅ Cleanup completed"

# Build commands
build:
	@echo "🔨 Building all services..."
	@cd apps/backend/gateway && go build -o gateway .
	@cd apps/backend/auth && go build -o auth-service .
	@cd apps/backend/file && go build -o file-service .
	@cd apps/backend/payment && go build -o payment-service .
	@cd apps/backend/crm && go build -o crm-service .
	@cd apps/backend/hrm && go build -o hrm-service .
	@cd apps/backend/pos && go build -o pos-service .
	@cd apps/frontend/web && npm run build
	@echo "✅ Build completed"

# Test commands
test:
	@echo "🧪 Running tests..."
	@cd apps/backend/gateway && go test ./...
	@cd apps/backend/auth && go test ./...
	@cd apps/backend/shared && go test ./...
	@cd pkg && go test ./...
	@echo "✅ Tests completed"

# Development helpers
dev-setup: setup install db-up
	@echo "✅ Development environment ready!"
	@echo "Run 'make start' to start all services"

logs:
	@echo "📋 Showing service logs..."
	@cd infra/docker && docker-compose logs -f

status:
	@echo "📊 Service status:"
	@echo ""
	@echo "🗄️ Database Services:"
	@cd infra/docker && docker-compose ps
	@echo ""
	@echo "🔧 Backend Services:"
	@pgrep -f "main.go" | while read pid; do echo "  PID $$pid: $$(ps -p $$pid -o comm=)"; done || echo "  No backend services running"
	@echo ""
	@echo "🌐 Frontend:"
	@pgrep -f "next-server" | while read pid; do echo "  PID $$pid: Next.js"; done || echo "  Frontend not running"

# Quick start for new developers
quickstart:
	@echo "🚀 Quick start for new developers..."
	@make dev-setup
	@sleep 5
	@make start
	@echo ""
	@echo "🎉 Zplus SaaS is now running!"
	@echo "📱 Frontend: http://localhost:3000"
	@echo "🔧 API Gateway: http://localhost:8000"
	@echo ""
	@echo "Run 'make stop' to stop all services"
