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
