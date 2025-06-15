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
