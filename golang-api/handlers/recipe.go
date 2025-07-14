package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recipehub/services"
)

func TrackRecipeView(c *gin.Context) {
	var req struct {
		RecipeID string `json:"recipe_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Recipe ID required",
		})
		return
	}

	// Get user ID from context (optional for anonymous users)
	userID, _ := c.Get("user_id")
	var userIDStr string
	if userID != nil {
		userIDStr = userID.(string)
	}

	// Get IP address and user agent
	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	// Track the view
	dbService := services.NewDatabaseService()
	if err := dbService.TrackRecipeView(req.RecipeID, userIDStr, ipAddress, userAgent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to track recipe view",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Recipe view tracked",
	})
}

func GetRecommendations(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})
		return
	}

	var req struct {
		Limit int `json:"limit"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		req.Limit = 10 // Default limit
	}

	if req.Limit <= 0 || req.Limit > 50 {
		req.Limit = 10
	}

	// This would call the PostgreSQL function for recommendations
	// For now, return a placeholder response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Recommendations retrieved",
		"data":    []interface{}{}, // Placeholder
	})
}
