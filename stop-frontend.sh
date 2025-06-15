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
