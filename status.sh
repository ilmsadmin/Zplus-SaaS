#!/bin/bash

# Zplus SaaS - Quick Status Check
echo "🔍 Zplus SaaS System Status Check"
echo "=================================="

# Check databases
echo ""
echo "🗄️ Database Services:"
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | grep zplus

# Check backend services
echo ""
echo "🔧 Backend Services:"
for port in 8000 8001 8002 8003 8004 8005 8006; do
    service_name="unknown"
    case $port in
        8000) service_name="Gateway" ;;
        8001) service_name="Auth" ;;
        8002) service_name="File" ;;
        8003) service_name="Payment" ;;
        8004) service_name="CRM" ;;
        8005) service_name="HRM" ;;
        8006) service_name="POS" ;;
    esac
    
    if curl -s --connect-timeout 2 http://localhost:$port/ > /dev/null; then
        echo "✅ $service_name ($port): Running"
    else
        echo "❌ $service_name ($port): Not responding"
    fi
done

# Check frontend
echo ""
echo "🌐 Frontend Service:"
if curl -s --connect-timeout 2 http://localhost:3001/ > /dev/null; then
    echo "✅ Next.js (3001): Running"
elif curl -s --connect-timeout 2 http://localhost:3000/ > /dev/null; then
    echo "✅ Next.js (3000): Running"
else
    echo "❌ Next.js: Not responding"
fi

echo ""
echo "🎯 Quick Access URLs:"
echo "   Frontend: http://localhost:3001"
echo "   API Gateway: http://localhost:8000"
echo "   GraphQL Playground: http://localhost:8000/playground"
echo "   Auth Service: http://localhost:8001"

echo ""
echo "📚 Demo Pages:"
echo "   Login: file://$(pwd)/mock/login.html"
echo "   System Admin: file://$(pwd)/mock/system-admin-dashboard.html"
echo "   Tenant Admin: file://$(pwd)/mock/tenant-admin-dashboard.html"
echo "   CRM Dashboard: file://$(pwd)/mock/customer-crm-dashboard.html"

echo ""
echo "🛠️ Management Commands:"
echo "   ./stop-all.sh     # Stop everything"
echo "   ./run-all.sh      # Start everything"
echo "   make status       # Detailed status"
echo "   make help         # All available commands"
