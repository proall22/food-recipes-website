# 🍳 RecipeHub - Food Recipe Website

A comprehensive food recipe platform built with Vue 3, Nuxt 3, Hasura, PostgreSQL, and Golang.

## ✨ Features

### 🔐 Authentication & User Management
- JWT-based authentication system
- User registration and login
- Password reset and email verification
- Social login (Google, Facebook) ready
- User profiles with avatars and bio

### 🍽️ Recipe Management
- Create, edit, and delete recipes
- Multiple image uploads with featured image
- Dynamic ingredients and cooking steps
- Nutrition information tracking
- Recipe categories and difficulty levels
- Premium recipe pricing with Chapa payment integration

### 🔍 Search & Discovery
- Advanced search with filters
- Search by ingredients, category, difficulty, prep time
- Recipe recommendations based on user preferences
- Category browsing
- Creator profiles and following system

### 💝 User Interactions
- Like and bookmark recipes
- Rate and review recipes with images
- Comment system
- User following/followers
- Recipe collections
- Achievement system

### 💳 Payment Integration
- Chapa payment gateway integration
- Premium recipe purchases
- Payment verification and webhooks
- Purchase history tracking

### 📊 Analytics & Tracking
- Recipe view tracking
- Search analytics
- User engagement metrics
- Popular recipes and trending content

## 🛠️ Tech Stack

### Frontend
- **Vue 3** - Progressive JavaScript framework
- **Nuxt 3** - Vue.js framework with SSR
- **Tailwind CSS** - Utility-first CSS framework
- **Vee-Validate** - Form validation
- **Pinia** - State management
- **Apollo Client** - GraphQL client
- **Lucide Vue** - Beautiful icons

### Backend
- **Hasura** - GraphQL API with real-time subscriptions
- **PostgreSQL** - Relational database with advanced features
- **Golang** - High-performance API server
- **JWT** - Secure authentication
- **Chapa** - Ethiopian payment gateway

### Infrastructure
- **Docker** - Containerization
- **Docker Compose** - Multi-container orchestration

## 🚀 Quick Start

### Prerequisites

Make sure you have the following installed:
- **Node.js** (v18 or higher)
- **Go** (v1.21 or higher)
- **Docker** and **Docker Compose**
- **Git**

### 1. Clone the Repository

\`\`\`bash
git clone <your-repo-url>
cd recipehub
\`\`\`

### 2. Run Setup Script

\`\`\`bash
chmod +x setup.sh
./setup.sh
\`\`\`

This script will:
- Create necessary directories
- Set up environment variables
- Install dependencies
- Start Docker services
- Initialize the database

### 3. Start the Services

#### Start Backend Services (Docker)
\`\`\`bash
cd hasura
docker-compose up -d
\`\`\`

#### Start Golang API
\`\`\`bash
cd golang-api
go run main.go
\`\`\`

#### Start Frontend Development Server
\`\`\`bash
npm run dev
\`\`\`

### 4. Access the Application

- **Frontend**: http://localhost:3000
- **Hasura Console**: http://localhost:8080 (admin secret: `myadminsecretkey`)
- **API Health Check**: http://localhost:8000/health
- **PostgreSQL**: localhost:5432 (user: `postgres`, password: `postgrespassword`)

## 📁 Project Structure

\`\`\`
recipehub/
├── assets/                 # CSS, images, fonts
├── components/            # Vue components
├── layouts/              # Nuxt layouts
├── pages/                # Vue pages (auto-routed)
├── stores/               # Pinia stores
├── plugins/              # Nuxt plugins
├── middleware/           # Route middleware
├── hasura/               # Hasura configuration
│   ├── metadata/         # Hasura metadata
│   ├── database/         # SQL schema and functions
│   └── docker-compose.yml
├── golang-api/           # Golang API server
│   ├── handlers/         # HTTP handlers
│   ├── services/         # Business logic
│   ├── models/           # Data models
│   ├── middleware/       # HTTP middleware
│   └── main.go
└── uploads/              # File uploads directory
\`\`\`

## 🔧 Configuration

### Environment Variables

Create `.env` file in the root directory:

\`\`\`env
# Database
DATABASE_URL=postgres://postgres:postgrespassword@localhost:5432/recipehub?sslmode=disable

# JWT
JWT_SECRET=your-256-bit-secret-key-here-make-it-long-and-random

# Hasura
HASURA_ADMIN_SECRET=myadminsecretkey
HASURA_ENDPOINT=http://localhost:8080/v1/graphql

# Chapa Payment
CHAPA_SECRET_KEY=your-chapa-secret-key

# File Upload
UPLOAD_DIR=./uploads
\`\`\`

### Chapa Payment Setup

1. Sign up at [Chapa.co](https://chapa.co)
2. Get your secret key from the dashboard
3. Update `CHAPA_SECRET_KEY` in your `.env` files
4. Configure webhook URL: `http://your-domain.com/payment/webhook`

## 🗄️ Database Schema

The application uses PostgreSQL with the following main tables:

- `users` - User accounts and profiles
- `categories` - Recipe categories
- `recipes` - Recipe information
- `recipe_ingredients` - Recipe ingredients (many-to-many)
- `recipe_steps` - Cooking instructions
- `recipe_images` - Recipe photos
- `recipe_nutrition` - Nutritional information
- `recipe_likes` - User likes
- `recipe_bookmarks` - User bookmarks
- `recipe_reviews` - User reviews and ratings
- `recipe_purchases` - Premium recipe purchases
- `user_follows` - User following relationships

## 🔐 Authentication Flow

1. User registers with email/password
2. JWT tokens are generated (access + refresh)
3. Tokens are stored in localStorage
4. Apollo client automatically includes tokens in requests
5. Hasura validates JWT and extracts user claims
6. Row-level security enforces data access rules

## 📱 API Endpoints

### Authentication
- `POST /auth/signup` - User registration
- `POST /auth/login` - User login
- `POST /auth/refresh` - Refresh tokens
- `POST /auth/forgot-password` - Password reset
- `POST /auth/verify-email` - Email verification

### File Upload
- `POST /upload/image` - Upload single image
- `POST /upload/multiple` - Upload multiple images
- `DELETE /upload/image/:filename` - Delete image

### Payments
- `POST /payment/initialize` - Initialize payment
- `POST /payment/verify` - Verify payment
- `POST /payment/webhook` - Payment webhook

## 🎨 UI/UX Features

### Design System
- Consistent color palette (Orange/Red gradient theme)
- Responsive design for all screen sizes
- Smooth animations and transitions
- Loading states and error handling
- Accessible components with ARIA labels

### User Experience
- Intuitive navigation with breadcrumbs
- Search with real-time suggestions
- Infinite scroll for recipe lists
- Image lazy loading
- Offline support (PWA ready)
- Print-friendly recipe pages

## 🧪 Testing

### Run Tests
\`\`\`bash
# Frontend tests
npm run test

# Backend tests
cd golang-api
go test ./...
\`\`\`

## 🚀 Deployment

### Frontend (Vercel/Netlify)
\`\`\`bash
npm run build
npm run generate  # For static deployment
\`\`\`

### Backend (Docker)
\`\`\`bash
docker build -t recipehub-api ./golang-api
docker run -p 8000:8000 recipehub-api
\`\`\`

### Database (Production)
- Use managed PostgreSQL (AWS RDS, Google Cloud SQL, etc.)
- Update connection strings in environment variables
- Run migrations and seed data

## 🔒 Security Features

- JWT token authentication
- Password hashing with bcrypt
- SQL injection prevention
- XSS protection
- CORS configuration
- Rate limiting
- File upload validation
- Environment variable protection

## 📈 Performance Optimizations

- Database indexing for fast queries
- Image optimization and compression
- GraphQL query optimization
- Caching strategies
- Lazy loading components
- Code splitting
- CDN integration ready

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Troubleshooting

### Common Issues

1. **Docker services not starting**
   \`\`\`bash
   docker-compose down
   docker-compose up -d --force-recreate
   \`\`\`

2. **Database connection errors**
   - Check if PostgreSQL is running
   - Verify connection string in `.env`
   - Ensure database exists

3. **Hasura metadata errors**
   \`\`\`bash
   hasura metadata reload
   \`\`\`

4. **File upload issues**
   - Check upload directory permissions
   - Verify file size limits
   - Ensure proper CORS configuration

### Getting Help

- Check the [Issues](link-to-issues) page
- Join our [Discord](link-to-discord) community
- Email support: support@recipehub.com

## 🎯 Roadmap

- [ ] Mobile app (React Native)
- [ ] Advanced recipe analytics
- [ ] AI-powered recipe suggestions
- [ ] Video recipe support
- [ ] Meal planning features
- [ ] Grocery list integration
- [ ] Recipe scaling calculator
- [ ] Nutritional analysis improvements

---

Made with ❤️ by the RecipeHub Team
\`\`\`

Finally, let's create a startup script for easy development:
