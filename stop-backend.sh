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
