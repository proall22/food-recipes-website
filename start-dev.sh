#!/bin/bash

echo "🍳 Starting RecipeHub Development Environment"
echo "============================================"

# Function to check if a port is in use
check_port() {
    if lsof -Pi :$1 -sTCP:LISTEN -t >/dev/null ; then
        echo "⚠️  Port $1 is already in use"
        return 1
    else
        return 0
    fi
}

# Check required ports
echo "🔍 Checking ports..."
check_port 3000 || echo "   Frontend port 3000 is busy"
check_port 8000 || echo "   API port 8000 is busy"
check_port 8080 || echo "   Hasura port 8080 is busy"
check_port 5432 || echo "   PostgreSQL port 5432 is busy"

echo ""

# Start Docker services
echo "🐳 Starting Docker services..."
cd hasura
docker-compose up -d

echo "⏳ Waiting for services to be ready..."
sleep 10

# Check if services are healthy
echo "🏥 Checking service health..."
if curl -f http://localhost:8080/healthz >/dev/null 2>&1; then
    echo "✅ Hasura is healthy"
else
    echo "❌ Hasura is not responding"
fi

if curl -f http://localhost:8000/health >/dev/null 2>&1; then
    echo "✅ API is healthy"
else
    echo "⚠️  API is not running yet (will start manually)"
fi

cd ..

# Start API in background
echo "🚀 Starting Golang API..."
cd golang-api
go run main.go &
API_PID=$!
cd ..

# Start frontend
echo "🎨 Starting Frontend..."
npm run dev &
FRONTEND_PID=$!

echo ""
echo "🎉 Development environment is starting!"
echo ""
echo "📱 Services:"
echo "   Frontend:  http://localhost:3000"
echo "   API:       http://localhost:8000"
echo "   Hasura:    http://localhost:8080"
echo "   Database:  localhost:5432"
echo ""
echo "🔑 Credentials:"
echo "   Hasura Admin Secret: myadminsecretkey"
echo "   Database User: postgres"
echo "   Database Password: postgrespassword"
echo ""
echo "⏹️  To stop all services, press Ctrl+C"

# Function to cleanup on exit
cleanup() {
    echo ""
    echo "🛑 Stopping services..."
    kill $API_PID 2>/dev/null
    kill $FRONTEND_PID 2>/dev/null
    cd hasura
    docker-compose down
    echo "✅ All services stopped"
    exit 0
}

# Set trap to cleanup on script exit
trap cleanup SIGINT SIGTERM

# Wait for user to stop
wait
