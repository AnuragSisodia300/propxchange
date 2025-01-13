package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"propxchange/schema"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddMoneyToWallet(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	amountStr := r.URL.Query().Get("amount")

	// Convert user ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Convert amount to float
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Find user and update wallet balance
	for i, user := range users {
		if user.ID == objectID {
			users[i].WalletBalance += amount

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":        "Money added successfully",
				"wallet_balance": users[i].WalletBalance,
			})
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func PurchaseUnitsFromWallet(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	propertyID := r.URL.Query().Get("property_id")
	unitCountStr := r.URL.Query().Get("unit_count")

	// Convert IDs and unit count
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	propertyObjID, err := primitive.ObjectIDFromHex(propertyID)
	if err != nil {
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}
	unitCount, err := strconv.Atoi(unitCountStr)
	if err != nil || unitCount <= 0 {
		http.Error(w, "Invalid unit count", http.StatusBadRequest)
		return
	}

	// Find user and property
	var user *schema.User
	var property *schema.Property
	for i, u := range users {
		if u.ID == objectID {
			user = &users[i]
			break
		}
	}
	for i, p := range Properties {
		if p.ID == propertyObjID {
			property = &Properties[i]
			break
		}
	}

	if user == nil || property == nil {
		http.Error(w, "User or property not found", http.StatusNotFound)
		return
	}

	// Calculate total cost and validate wallet balance
	totalCost := float64(unitCount) * property.PerUnitPrice
	if user.WalletBalance < totalCost {
		http.Error(w, "Insufficient wallet balance", http.StatusPaymentRequired)
		return
	}

	// Deduct amount from wallet and update property fund raised
	user.WalletBalance -= totalCost
	property.FundRaised += totalCost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":        "Units purchased successfully",
		"wallet_balance": user.WalletBalance,
		"fund_raised":    property.FundRaised,
	})
}
