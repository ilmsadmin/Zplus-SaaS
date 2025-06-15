#!/bin/bash

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🚀 Starting Zplus SaaS Backend Services${NC}"

# Array of services with their ports (zsh/macOS compatible)
services=(
    "gateway:8000"
    "auth:8001"
    "file:8002"
    "payment:8003"
    "crm:8004"
    "hrm:8005"
    "pos:8006"
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
for service_port in "${services[@]}"; do
    service="${service_port%%:*}"
    port="${service_port##*:}"
    start_service "$service" "$port"
done

echo -e "${GREEN}✓ All backend services started${NC}"
echo "To stop services, run: ./stop-backend.sh"

# Keep script running
wait
