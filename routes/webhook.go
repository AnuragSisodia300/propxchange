package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/webhook"
)

func WebhookRoutes() {
	http.HandleFunc("/webhook", StripeWebhookHandler)
}

// StripeWebhookHandler handles Stripe webhook events
func StripeWebhookHandler(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	const endpointSecret = "whsec_your_webhook_secret"

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), endpointSecret)
	if err != nil {
		http.Error(w, "Webhook signature verification failed", http.StatusBadRequest)
		return
	}

	// Handle the event
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			http.Error(w, "Failed to parse webhook data", http.StatusBadRequest)
			return
		}
		// Update database, mark payment as completed
		fmt.Printf("Payment succeeded: %s\n", paymentIntent.ID)

	case "payment_intent.payment_failed":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			http.Error(w, "Failed to parse webhook data", http.StatusBadRequest)
			return
		}
		// Handle payment failure
		fmt.Printf("Payment failed: %s\n", paymentIntent.ID)

	default:
		fmt.Printf("Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}
