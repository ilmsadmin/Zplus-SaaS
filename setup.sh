#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Zplus SaaS - Project Setup Script${NC}"
echo "================================================="

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to print status
print_status() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

# Check required tools
echo -e "${BLUE}📋 Checking required tools...${NC}"

if ! command_exists go; then
    print_error "Go is not installed. Please install Go 1.21+ from https://golang.org/dl/"
    exit 1
else
    print_status "Go is installed: $(go version)"
fi

if ! command_exists node; then
    print_error "Node.js is not installed. Please install Node.js from https://nodejs.org/"
    exit 1
else
    print_status "Node.js is installed: $(node --version)"
fi

if ! command_exists npm; then
    print_error "npm is not installed. Please install npm"
    exit 1
else
    print_status "npm is installed: $(npm --version)"
fi

if ! command_exists docker; then
    print_warning "Docker is not installed. You can install it from https://docker.com/"
    print_info "Docker is optional but recommended for running databases"
else
    print_status "Docker is installed: $(docker --version)"
fi

if ! command_exists docker-compose; then
    print_warning "Docker Compose is not installed"
    print_info "Docker Compose is optional but recommended for running databases"
else
    print_status "Docker Compose is installed: $(docker-compose --version)"
fi

echo ""

# Setup environment files
echo -e "${BLUE}🔧 Setting up environment files...${NC}"

if [ ! -f ".env" ]; then
    cp .env.example .env
    print_status "Created .env file from .env.example"
else
    print_info ".env file already exists"
fi

# Copy environment file to backend services
backend_services=("gateway" "auth" "file" "payment" "crm" "hrm" "pos" "shared")
for service in "${backend_services[@]}"; do
    if [ -d "apps/backend/$service" ]; then
        if [ ! -f "apps/backend/$service/.env" ]; then
            cp .env "apps/backend/$service/.env"
            print_status "Created .env for $service service"
        fi
    fi
done

# Copy environment file to frontend
if [ ! -f "apps/frontend/web/.env.local" ]; then
    cp .env.example apps/frontend/web/.env.local
    print_status "Created .env.local for frontend"
fi

echo ""

# Setup Go modules
echo -e "${BLUE}📦 Setting up Go modules...${NC}"

# Setup pkg module
cd pkg
if [ ! -f "go.mod" ]; then
    go mod init github.com/ilmsadmin/Zplus-SaaS/pkg
    print_status "Initialized pkg module"
else
    print_info "pkg module already exists"
fi
go mod tidy
print_status "Updated pkg dependencies"
cd ..

# Setup shared module
cd apps/backend/shared
go mod tidy
print_status "Updated shared module dependencies"
cd ../../..

# Setup backend services
for service in "${backend_services[@]}"; do
    if [ -d "apps/backend/$service" ] && [ -f "apps/backend/$service/go.mod" ]; then
        cd "apps/backend/$service"
        go mod tidy
        print_status "Updated $service module dependencies"
        cd ../../..
    fi
done

echo ""

# Setup frontend dependencies
echo -e "${BLUE}🌐 Setting up frontend dependencies...${NC}"

cd apps/frontend/web
npm install
print_status "Installed frontend dependencies"
cd ../../..

# Setup UI components
if [ -d "apps/frontend/ui" ]; then
    cd apps/frontend/ui
    npm install
    print_status "Installed UI components dependencies"
    cd ../../..
fi

echo ""

# Database setup
echo -e "${BLUE}🗄️ Database setup...${NC}"

if command_exists docker && command_exists docker-compose; then
    print_info "You can start the database services with:"
    echo "  cd infra/docker && docker-compose up -d postgres mongodb redis"
    print_info "Or start all services with:"
    echo "  cd infra/docker && docker-compose up -d"
else
    print_warning "Docker not available. Please setup PostgreSQL, MongoDB, and Redis manually"
    print_info "Database requirements:"
    echo "  - PostgreSQL (port 5432)"
    echo "  - MongoDB (port 27017)"
    echo "  - Redis (port 6379)"
fi

echo ""

# Create run scripts
echo -e "${BLUE}📝 Creating run scripts...${NC}"

# Backend run script
cat > run-backend.sh << 'EOF'
#!/bin/bash

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🚀 Starting Zplus SaaS Backend Services${NC}"

# Array of services with their ports
declare -A services=(
    ["gateway"]="8000"
    ["auth"]="8001"
    ["file"]="8002"
    ["payment"]="8003"
    ["crm"]="8004"
    ["hrm"]="8005"
    ["pos"]="8006"
)

# Function to start service
start_service() {
    local service=$1
    local port=$2
    
    if [ -d "apps/backend/$service" ] && [ -f "apps/backend/$service/main.go" ]; then
        echo -e "${GREEN}Starting $service service on port $port${NC}"
        cd "apps/backend/$service"
        go run main.go &
        echo $! > "../../../.$service.pid"
        cd ../../..
        sleep 2
    else
        echo "Service $service not found or missing main.go"
    fi
}

# Start all services
for service in "${!services[@]}"; do
    start_service "$service" "${services[$service]}"
done

echo -e "${GREEN}✓ All backend services started${NC}"
echo "To stop services, run: ./stop-backend.sh"

# Keep script running
wait
EOF

chmod +x run-backend.sh
print_status "Created run-backend.sh"

# Backend stop script
cat > stop-backend.sh << 'EOF'
#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${RED}🛑 Stopping Zplus SaaS Backend Services${NC}"

services=("gateway" "auth" "file" "payment" "crm" "hrm" "pos")

for service in "${services[@]}"; do
    if [ -f ".$service.pid" ]; then
        pid=$(cat ".$service.pid")
        if kill -0 "$pid" 2>/dev/null; then
            kill "$pid"
            echo -e "${GREEN}✓ Stopped $service service (PID: $pid)${NC}"
        fi
        rm ".$service.pid"
    fi
done

echo -e "${GREEN}✓ All backend services stopped${NC}"
EOF

chmod +x stop-backend.sh
print_status "Created stop-backend.sh"

# Frontend run script
cat > run-frontend.sh << 'EOF'
#!/bin/bash

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🌐 Starting Zplus SaaS Frontend${NC}"

cd apps/frontend/web
npm run dev &
echo $! > "../../../.frontend.pid"
cd ../../..

echo -e "${GREEN}✓ Frontend started on http://localhost:3000${NC}"
echo "To stop frontend, run: ./stop-frontend.sh"

wait
EOF

chmod +x run-frontend.sh
print_status "Created run-frontend.sh"

# Frontend stop script
cat > stop-frontend.sh << 'EOF'
#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${RED}🛑 Stopping Zplus SaaS Frontend${NC}"

if [ -f ".frontend.pid" ]; then
    pid=$(cat ".frontend.pid")
    if kill -0 "$pid" 2>/dev/null; then
        kill "$pid"
        echo -e "${GREEN}✓ Stopped frontend (PID: $pid)${NC}"
    fi
    rm ".frontend.pid"
fi

echo -e "${GREEN}✓ Frontend stopped${NC}"
EOF

chmod +x stop-frontend.sh
print_status "Created stop-frontend.sh"

# Complete run script
cat > run-all.sh << 'EOF'
#!/bin/bash

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🚀 Starting Complete Zplus SaaS Application${NC}"

# Start databases if Docker is available
if command -v docker >/dev/null 2>&1 && command -v docker-compose >/dev/null 2>&1; then
    echo -e "${BLUE}🗄️ Starting database services...${NC}"
    cd infra/docker
    docker-compose up -d postgres mongodb redis
    cd ../..
    sleep 5
fi

# Start backend services
echo -e "${BLUE}🔧 Starting backend services...${NC}"
./run-backend.sh &
sleep 10

# Start frontend
echo -e "${BLUE}🌐 Starting frontend...${NC}"
./run-frontend.sh &

echo -e "${GREEN}✅ Zplus SaaS is now running!${NC}"
echo ""
echo "📱 Frontend: http://localhost:3000"
echo "🔧 Gateway API: http://localhost:8000"
echo "🔐 Auth Service: http://localhost:8001"
echo ""
echo "To stop all services, run: ./stop-all.sh"

wait
EOF

chmod +x run-all.sh
print_status "Created run-all.sh"

# Complete stop script
cat > stop-all.sh << 'EOF'
#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${RED}🛑 Stopping Complete Zplus SaaS Application${NC}"

# Stop frontend
./stop-frontend.sh

# Stop backend
./stop-backend.sh

# Stop databases if Docker is available
if command -v docker >/dev/null 2>&1 && command -v docker-compose >/dev/null 2>&1; then
    echo -e "${RED}🗄️ Stopping database services...${NC}"
    cd infra/docker
    docker-compose down
    cd ../..
fi

echo -e "${GREEN}✅ All services stopped${NC}"
EOF

chmod +x stop-all.sh
print_status "Created stop-all.sh"

echo ""

# Final instructions
echo -e "${GREEN}🎉 Setup completed successfully!${NC}"
echo "================================================="
echo ""
echo -e "${BLUE}📋 Next steps:${NC}"
echo ""
echo "1. 🗄️ Start databases:"
echo "   cd infra/docker && docker-compose up -d"
echo ""
echo "2. 🚀 Start the complete application:"
echo "   ./run-all.sh"
echo ""
echo "3. 🌐 Or start services individually:"
echo "   ./run-backend.sh    # Start all backend services"
echo "   ./run-frontend.sh   # Start frontend only"
echo ""
echo "4. 🛑 Stop services:"
echo "   ./stop-all.sh       # Stop everything"
echo ""
echo -e "${BLUE}📱 Application URLs:${NC}"
echo "   Frontend: http://localhost:3000"
echo "   Gateway API: http://localhost:8000"
echo "   Auth Service: http://localhost:8001"
echo "   Traefik Dashboard: http://localhost:8080 (if using Docker)"
echo ""
echo -e "${YELLOW}⚠️ Important:${NC}"
echo "   - Make sure to start databases before backend services"
echo "   - Update .env file with your configuration"
echo "   - Check firewall settings if services don't start"
echo ""
echo -e "${GREEN}✅ Happy coding!${NC}"
