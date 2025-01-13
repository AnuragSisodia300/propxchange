package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"propxchange/schema"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var properties []schema.Property
var Properties []schema.Property

// CreateProperty handles property creation
func CreateProperty(w http.ResponseWriter, r *http.Request) {
	var property schema.Property
	err := json.NewDecoder(r.Body).Decode(&property)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign new ID and timestamps
	property.ID = primitive.NewObjectID()
	property.CreatedAt = time.Now()
	property.UpdatedAt = time.Now()

	// Calculate units and fund raised
	if property.PerUnitPrice > 0 {
		property.Units = int(property.Price / property.PerUnitPrice)
		property.FundRaised = 0
	}

	// Add property to the in-memory store
	Properties = append(Properties, property)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(property)
}

// ListProperties lists all properties
func ListProperties(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Properties)
}

// GetProperty retrieves a property by ID
func GetProperty(w http.ResponseWriter, id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	for _, property := range Properties {
		if property.ID == objectID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(property)
			return
		}
	}

	http.Error(w, "Property not found", http.StatusNotFound)
}

// UpdateProperty updates a property
func UpdateProperty(w http.ResponseWriter, r *http.Request, id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedProperty schema.Property
	err = json.NewDecoder(r.Body).Decode(&updatedProperty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, property := range Properties {
		if property.ID == objectID {
			// Update fields and recalculate
			updatedProperty.ID = property.ID
			updatedProperty.CreatedAt = property.CreatedAt
			updatedProperty.UpdatedAt = time.Now()

			if updatedProperty.PerUnitPrice > 0 {
				updatedProperty.Units = int(updatedProperty.Price / updatedProperty.PerUnitPrice)
				updatedProperty.FundRaised = float64(updatedProperty.Units) * updatedProperty.PerUnitPrice
			}

			Properties[i] = updatedProperty

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProperty)
			return
		}
	}

	http.Error(w, "Property not found", http.StatusNotFound)
}

// DeleteProperty deletes a property
func DeleteProperty(w http.ResponseWriter, id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	for i, property := range Properties {
		if property.ID == objectID {
			Properties = append(Properties[:i], Properties[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Property not found", http.StatusNotFound)
}

// PurchaseUnits purchase number of units
func PurchaseUnits(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UserID     string  `json:"user_id"`
		PropertyID string  `json:"property_id"`
		Units      int     `json:"units"`
		PaymentID  string  `json:"payment_id"`
		Amount     float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Convert IDs
	userID, err := primitive.ObjectIDFromHex(request.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	propertyID, err := primitive.ObjectIDFromHex(request.PropertyID)
	if err != nil {
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}

	// Update property fund and units
	for i, property := range Properties {
		if property.ID == propertyID {
			if request.Units > property.Units {
				http.Error(w, "Not enough units available", http.StatusBadRequest)
				return
			}

			Properties[i].Units -= request.Units
			Properties[i].FundRaised += request.Amount

			// Update user's purchased properties
			for j, user := range users {
				if user.ID == userID {
					user.Purchased = append(user.Purchased, schema.UserProperty{
						PropertyID: propertyID,
						Units:      request.Units,
						AmountPaid: request.Amount,
						PaymentID:  request.PaymentID,
					})
					users[j] = user

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(map[string]string{
						"status": "Purchase successful",
					})
					return
				}
			}

			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
	}

	http.Error(w, "Property not found", http.StatusNotFound)
}
