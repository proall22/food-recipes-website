package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ChapaService struct {
	secretKey string
	baseURL   string
}

func NewChapaService() *ChapaService {
	return &ChapaService{
		secretKey: os.Getenv("CHAPA_SECRET_KEY"),
		baseURL:   "https://api.chapa.co/v1",
	}
}

type PaymentRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Email       string  `json:"email"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	TxRef       string  `json:"tx_ref"`
	CallbackURL string  `json:"callback_url"`
	ReturnURL   string  `json:"return_url"`
	Description string  `json:"description"`
}

type PaymentResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
		TxRef       string `json:"tx_ref"`
	} `json:"data"`
}

type VerificationResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		Amount      float64 `json:"amount"`
		Currency    string  `json:"currency"`
		Email       string  `json:"email"`
		FirstName   string  `json:"first_name"`
		LastName    string  `json:"last_name"`
		TxRef       string  `json:"tx_ref"`
		Status      string  `json:"status"`
		Reference   string  `json:"reference"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
	} `json:"data"`
}

func (s *ChapaService) InitializePayment(req *PaymentRequest) (*PaymentResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", s.baseURL+"/transaction/initialize", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+s.secretKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var paymentResp PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("payment initialization failed: %s", paymentResp.Message)
	}

	return &paymentResp, nil
}

func (s *ChapaService) VerifyPayment(txRef string) (*VerificationResponse, error) {
	url := fmt.Sprintf("%s/transaction/verify/%s", s.baseURL, txRef)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+s.secretKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var verifyResp VerificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&verifyResp); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("payment verification failed: %s", verifyResp.Message)
	}

	return &verifyResp, nil
}
