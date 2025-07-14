package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"recipehub/models"
	"recipehub/services"
)

type AuthHandler struct {
	authService   *services.AuthService
	dbService     *services.DatabaseService
	hasuraService *services.HasuraService
}

func NewAuthHandler(authService *services.AuthService, dbService *services.DatabaseService, hasuraService *services.HasuraService) *AuthHandler {
	return &AuthHandler{
		authService:   authService,
		dbService:     dbService,
		hasuraService: hasuraService,
	}
}

type SignupRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required,min=3,max=50"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Bio       string `json:"bio"`
	Avatar    string `json:"avatar"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	User         *models.User `json:"user,omitempty"`
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Check if user already exists
	existingUser, _ := h.dbService.GetUserByEmail(req.Email)
	if existingUser != nil {
		c.JSON(http.StatusConflict, AuthResponse{
			Success: false,
			Message: "User with this email already exists",
		})
		return
	}

	// Check if username is taken
	existingUser, _ = h.dbService.GetUserByUsername(req.Username)
	if existingUser != nil {
		c.JSON(http.StatusConflict, AuthResponse{
			Success: false,
			Message: "Username is already taken",
		})
		return
	}

	// Hash password
	hashedPassword, err := h.authService.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to process password",
		})
		return
	}

	// Create user
	user := &models.User{
		Email:     req.Email,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  hashedPassword,
		Bio:       req.Bio,
		Avatar:    req.Avatar,
	}

	if err := h.dbService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to create user account",
		})
		return
	}

	// Generate email verification token
	verificationToken, err := h.authService.GenerateEmailVerificationToken()
	if err != nil {
		// Log error but don't fail signup
		// In production, you'd want to send verification email
	}

	c.JSON(http.StatusCreated, AuthResponse{
		Success: true,
		Message: "Account created successfully. Please check your email for verification.",
		User:    user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Get user by email
	user, err := h.dbService.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	// Check password
	if !h.authService.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Message: "Account is deactivated",
		})
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := h.authService.GenerateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to generate authentication tokens",
		})
		return
	}

	// Remove password from response
	user.Password = ""

	c.JSON(http.StatusOK, AuthResponse{
		Success:      true,
		Message:      "Login successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Message: "Authorization header required",
		})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := h.authService.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Message: "Invalid refresh token",
		})
		return
	}

	// Get user
	user, err := h.dbService.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, AuthResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}

	// Generate new tokens
	accessToken, refreshToken, err := h.authService.GenerateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to generate new tokens",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Success:      true,
		Message:      "Token refreshed successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Invalid email address",
		})
		return
	}

	// Check if user exists
	user, err := h.dbService.GetUserByEmail(req.Email)
	if err != nil {
		// Don't reveal if email exists or not
		c.JSON(http.StatusOK, AuthResponse{
			Success: true,
			Message: "If an account with this email exists, a password reset link has been sent.",
		})
		return
	}

	// Generate reset token
	resetToken, err := h.authService.GenerateResetToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to generate reset token",
		})
		return
	}

	// Save reset token to database (implement this in dbService)
	if err := h.dbService.SavePasswordResetToken(user.ID, resetToken); err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to process password reset request",
		})
		return
	}

	// In production, send email with reset link
	// emailService.SendPasswordResetEmail(user.Email, resetToken)

	c.JSON(http.StatusOK, AuthResponse{
		Success: true,
		Message: "If an account with this email exists, a password reset link has been sent.",
	})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req struct {
		Token    string `json:"token" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Invalid request data",
		})
		return
	}

	// Verify reset token
	userID, err := h.dbService.VerifyPasswordResetToken(req.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Invalid or expired reset token",
		})
		return
	}

	// Hash new password
	hashedPassword, err := h.authService.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to process new password",
		})
		return
	}

	// Update password
	if err := h.dbService.UpdateUserPassword(userID, hashedPassword); err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to update password",
		})
		return
	}

	// Delete used reset token
	h.dbService.DeletePasswordResetToken(req.Token)

	c.JSON(http.StatusOK, AuthResponse{
		Success: true,
		Message: "Password reset successfully",
	})
}

func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Invalid verification token",
		})
		return
	}

	// Verify email token
	userID, err := h.dbService.VerifyEmailToken(req.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Invalid or expired verification token",
		})
		return
	}

	// Mark email as verified
	if err := h.dbService.MarkEmailAsVerified(userID); err != nil {
		c.JSON(http.StatusInternalServerError, AuthResponse{
			Success: false,
			Message: "Failed to verify email",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Success: true,
		Message: "Email verified successfully",
	})
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var req struct {
		GoogleToken string `json:"google_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Google token required",
		})
		return
	}

	// Verify Google token and get user info
	// This would integrate with Google OAuth API
	// For now, we'll return a placeholder response

	c.JSON(http.StatusNotImplemented, AuthResponse{
		Success: false,
		Message: "Google login not implemented yet",
	})
}

func (h *AuthHandler) FacebookLogin(c *gin.Context) {
	var req struct {
		FacebookToken string `json:"facebook_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AuthResponse{
			Success: false,
			Message: "Facebook token required",
		})
		return
	}

	// Verify Facebook token and get user info
	// This would integrate with Facebook Graph API
	// For now, we'll return a placeholder response

	c.JSON(http.StatusNotImplemented, AuthResponse{
		Success: false,
		Message: "Facebook login not implemented yet",
	})
}
