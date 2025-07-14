package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"recipehub/models"
	"recipehub/services"
)

type PaymentHandler struct {
	chapaService  *services.ChapaService
	dbService     *services.DatabaseService
	hasuraService *services.HasuraService
}

func NewPaymentHandler(chapaService *services.ChapaService, dbService *services.DatabaseService, hasuraService *services.HasuraService) *PaymentHandler {
	return &PaymentHandler{
		chapaService:  chapaService,
		dbService:     dbService,
		hasuraService: hasuraService,
	}
}

type InitializePaymentRequest struct {
	RecipeID string  `json:"recipe_id" binding:"required"`
	Amount   float64 `json:"amount" binding:"required,gt=0"`
}

type PaymentResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	CheckoutURL string `json:"checkout_url,omitempty"`
	TxRef       string `json:"tx_ref,omitempty"`
}

func (h *PaymentHandler) InitializePayment(c *gin.Context) {
	var req InitializePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Get user from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, PaymentResponse{
			Success: false,
			Message: "User not authenticated",
		})
		return
	}

	// Get user details
	user, err := h.dbService.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, PaymentResponse{
			Success: false,
			Message: "Failed to get user details",
		})
		return
	}

	// Generate unique transaction reference
	txRef := fmt.Sprintf("recipe_%s_%d", uuid.New().String()[:8], time.Now().Unix())

	// Create payment request
	paymentReq := &services.PaymentRequest{
		Amount:      req.Amount,
		Currency:    "ETB", // Ethiopian Birr
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		TxRef:       txRef,
		CallbackURL: "http://localhost:8000/payment/webhook",
		ReturnURL:   "http://localhost:3000/payment/success",
		Description: fmt.Sprintf("Purchase recipe - %s", req.RecipeID),
	}

	// Initialize payment with Chapa
	paymentResp, err := h.chapaService.InitializePayment(paymentReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, PaymentResponse{
			Success: false,
			Message: "Failed to initialize payment: " + err.Error(),
		})
		return
	}

	// Save purchase record
	purchase := &models.RecipePurchase{
		RecipeID:         req.RecipeID,
		UserID:           userID.(string),
		Amount:           req.Amount,
		PaymentMethod:    "chapa",
		PaymentReference: txRef,
		Status:           "pending",
	}

	if err := h.dbService.CreateRecipePurchase(purchase); err != nil {
		c.JSON(http.StatusInternalServerError, PaymentResponse{
			Success: false,
			Message: "Failed to create purchase record",
		})
		return
	}

	c.JSON(http.StatusOK, PaymentResponse{
		Success:     true,
		Message:     "Payment initialized successfully",
		CheckoutURL: paymentResp.Data.CheckoutURL,
		TxRef:       txRef,
	})
}

func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	var req struct {
		TxRef string `json:"tx_ref" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Success: false,
			Message: "Transaction reference required",
		})
		return
	}

	// Verify payment with Chapa
	verifyResp, err := h.chapaService.VerifyPayment(req.TxRef)
	if err != nil {
		c.JSON(http.StatusInternalServerError, PaymentResponse{
			Success: false,
			Message: "Failed to verify payment: " + err.Error(),
		})
		return
	}

	// Get purchase record
	purchase, err := h.dbService.GetRecipePurchaseByReference(req.TxRef)
	if err != nil {
		c.JSON(http.StatusNotFound, PaymentResponse{
			Success: false,
			Message: "Purchase record not found",
		})
		return
	}

	// Update purchase status based on verification result
	var status string
	if verifyResp.Data.Status == "success" {
		status = "completed"
	} else {
		status = "failed"
	}

	if err := h.dbService.UpdateRecipePurchaseStatus(purchase.ID, status); err != nil {
		c.JSON(http.StatusInternalServerError, PaymentResponse{
			Success: false,
			Message: "Failed to update purchase status",
		})
		return
	}

	c.JSON(http.StatusOK, PaymentResponse{
		Success: verifyResp.Data.Status == "success",
		Message: fmt.Sprintf("Payment %s", verifyResp.Data.Status),
	})
}

func (h *PaymentHandler) WebhookHandler(c *gin.Context) {
	// Handle Chapa webhook notifications
	var webhookData map[string]interface{}
	if err := c.ShouldBindJSON(&webhookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid webhook data"})
		return
	}

	// Extract transaction reference
	txRef, ok := webhookData["tx_ref"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing transaction reference"})
		return
	}

	// Verify the webhook with Chapa
	verifyResp, err := h.chapaService.VerifyPayment(txRef)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify payment"})
		return
	}

	// Get purchase record
	purchase, err := h.dbService.GetRecipePurchaseByReference(txRef)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase record not found"})
		return
	}

	// Update purchase status
	var status string
	if verifyResp.Data.Status == "success" {
		status = "completed"
	} else {
		status = "failed"
	}

	if err := h.dbService.UpdateRecipePurchaseStatus(purchase.ID, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update purchase status"})
		return
	}

	// If payment successful, trigger Hasura event
	if status == "completed" {
		// Trigger recipe purchase event in Hasura
		h.hasuraService.TriggerEvent("recipe_purchased", map[string]interface{}{
			"user_id":   purchase.UserID,
			"recipe_id": purchase.RecipeID,
			"amount":    purchase.Amount,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "processed"})
}
