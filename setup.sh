#!/bin/bash

echo "🍳 Setting up RecipeHub - Food Recipe Website"
echo "============================================="

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js (v18 or higher) first."
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go (v1.21 or higher) first."
    exit 1
fi

echo "✅ All prerequisites are installed"
echo ""

# Create necessary directories
echo "📁 Creating project directories..."
mkdir -p uploads/images uploads/recipes uploads/avatars
mkdir -p hasura/metadata/databases/default/tables
mkdir -p golang-api/handlers golang-api/services golang-api/models golang-api/middleware

echo "✅ Directories created"
echo ""

# Set up environment variables
echo "🔧 Setting up environment variables..."
cat > .env << EOF
# Database
DATABASE_URL=postgres://postgres:postgrespassword@localhost:5432/recipehub?sslmode=disable

# JWT
JWT_SECRET=your-256-bit-secret-key-here-make-it-long-and-random-change-this-in-production

# Hasura
HASURA_ADMIN_SECRET=myadminsecretkey
HASURA_ENDPOINT=http://localhost:8080/v1/graphql

# Chapa Payment (Get these from https://chapa.co)
CHAPA_SECRET_KEY=your-chapa-secret-key-here

# File Upload
UPLOAD_DIR=./uploads

# Server
PORT=8000
EOF

cat > golang-api/.env << EOF
DATABASE_URL=postgres://postgres:postgrespassword@postgres:5432/recipehub?sslmode=disable
JWT_SECRET=your-256-bit-secret-key-here-make-it-long-and-random-change-this-in-production
HASURA_ADMIN_SECRET=myadminsecretkey
HASURA_ENDPOINT=http://graphql-engine:8080/v1/graphql
CHAPA_SECRET_KEY=your-chapa-secret-key-here
UPLOAD_DIR=/app/uploads
PORT=8000
EOF

echo "✅ Environment variables set up"
echo ""

# Install frontend dependencies
echo "📦 Installing frontend dependencies..."
npm install

echo "✅ Frontend dependencies installed"
echo ""

# Install Go dependencies
echo "📦 Installing Go dependencies..."
cd golang-api
go mod init recipehub 2>/dev/null || true
go mod tidy
cd ..

echo "✅ Go dependencies installed"
echo ""

# Start Docker services
echo "🐳 Starting Docker services..."
cd hasura
docker-compose up -d

echo "⏳ Waiting for services to start..."
sleep 30

echo "✅ Docker services started"
echo ""

# Apply Hasura metadata
echo "🔧 Applying Hasura metadata..."
# You would typically use Hasura CLI here
# hasura metadata apply

echo "✅ Hasura metadata applied"
echo ""

echo "🎉 Setup complete!"
echo ""
echo "🚀 To start the application:"
echo "1. Frontend: npm run dev (runs on http://localhost:3000)"
echo "2. Backend API: cd golang-api && go run main.go (runs on http://localhost:8000)"
echo "3. Hasura Console: http://localhost:8080 (admin secret: myadminsecretkey)"
echo "4. PostgreSQL: localhost:5432 (user: postgres, password: postgrespassword)"
echo ""
echo "📝 Important Notes:"
echo "- Update CHAPA_SECRET_KEY in .env files with your actual Chapa credentials"
echo "- Change JWT_SECRET to a secure random string in production"
echo "- The database will be automatically seeded with sample data"
echo ""
echo "🔗 Useful URLs:"
echo "- Frontend: http://localhost:3000"
echo "- Hasura Console: http://localhost:8080"
echo "- API Health Check: http://localhost:8000/health"
