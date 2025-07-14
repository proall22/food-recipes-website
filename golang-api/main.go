package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"recipehub/handlers"
	"recipehub/middleware"
	"recipehub/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize services
	dbService := services.NewDatabaseService()
	authService := services.NewAuthService()
	fileService := services.NewFileService()
	chapaService := services.NewChapaService()
	hasuraService := services.NewHasuraService()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, dbService, hasuraService)
	fileHandler := handlers.NewFileHandler(fileService)
	paymentHandler := handlers.NewPaymentHandler(chapaService, dbService, hasuraService)

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Auth routes (Hasura Actions)
	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.RefreshToken)
		auth.POST("/forgot-password", authHandler.ForgotPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)
		auth.POST("/verify-email", authHandler.VerifyEmail)
		auth.POST("/google-login", authHandler.GoogleLogin)
		auth.POST("/facebook-login", authHandler.FacebookLogin)
	}

	// File upload routes
	upload := r.Group("/upload")
	upload.Use(middleware.AuthMiddleware(authService))
	{
		upload.POST("/image", fileHandler.UploadImage)
		upload.POST("/multiple", fileHandler.UploadMultipleImages)
		upload.DELETE("/image/:filename", fileHandler.DeleteImage)
	}

	// Payment routes (Hasura Actions)
	payment := r.Group("/payment")
	payment.Use(middleware.AuthMiddleware(authService))
	{
		payment.POST("/initialize", paymentHandler.InitializePayment)
		payment.POST("/verify", paymentHandler.VerifyPayment)
		payment.POST("/webhook", paymentHandler.WebhookHandler)
	}

	// Recipe actions
	recipe := r.Group("/recipe")
	recipe.Use(middleware.AuthMiddleware(authService))
	{
		recipe.POST("/view", handlers.TrackRecipeView)
		recipe.POST("/recommend", handlers.GetRecommendations)
	}

	// Static file serving
	r.Static("/uploads", "./uploads")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
