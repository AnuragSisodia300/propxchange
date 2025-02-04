package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

// Initialize Stripe
func init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

// PaymentRequest defines the structure of the payment request
type PaymentRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	Source      string  `json:"source"`
}

// HandlePayment processes the payment request
func HandlePayment(w http.ResponseWriter, r *http.Request) {
	var paymentReq PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create a charge
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(paymentReq.Amount * 100)),
		Currency:    stripe.String(paymentReq.Currency),
		Description: stripe.String(paymentReq.Description),
	}
	params.SetSource(paymentReq.Source)

	charge, err := charge.New(params)
	if err != nil {
		http.Error(w, "Payment failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"payment_id": charge.ID,
		"status":     string(charge.Status),
	})
}
