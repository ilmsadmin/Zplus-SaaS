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
