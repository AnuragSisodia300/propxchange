package controllers

import (
	"encoding/json"
	"net/http"
	"propxchange/schema"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var watchlist []schema.Watchlist

// AddToWatchlist adds a property to the watchlist
func AddToWatchlist(w http.ResponseWriter, r *http.Request) {
	var watchlistItem schema.Watchlist
	err := json.NewDecoder(r.Body).Decode(&watchlistItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a new ObjectID
	watchlistItem.ID = primitive.NewObjectID()

	// Assign timestamp values
	watchlistItem.CreatedAt = time.Now()
	watchlistItem.UpdatedAt = time.Now()

	// Add the watchlist item to the in-memory store
	watchlist = append(watchlist, watchlistItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(watchlistItem)
}

// ListWatchlist returns all properties in the watchlist
func ListWatchlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(watchlist)
}

// GetWatchlist retrieves all watchlist items
func GetWatchlist(w http.ResponseWriter, r *http.Request) {
	// Prepare the response with string representations of ObjectIDs
	var response []map[string]interface{}

	for _, item := range watchlist {
		// Convert ObjectID fields to strings using `.Hex()`
		responseItem := map[string]interface{}{
			"id":         item.ID.Hex(),
			"property":   item.PropertyID.Hex(),
			"user_id":    item.UserID.Hex(),
			"created_at": item.CreatedAt,
			"updated_at": item.UpdatedAt,
		}
		response = append(response, responseItem)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateWatchlist updates a watchlist item by its ID
func UpdateWatchlist(w http.ResponseWriter, r *http.Request) {
	// Extract the watchlist ID from the URL
	watchlistID := r.URL.Query().Get("id")

	// Convert the watchlist ID from string to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(watchlistID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Decode the request body to get the updated data
	var updatedItem schema.Watchlist
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Iterate through the in-memory watchlist and update the matching item
	for i, item := range watchlist {
		if item.ID == objectID {
			// Update fields and timestamps
			updatedItem.ID = item.ID
			updatedItem.UserID = item.UserID
			updatedItem.CreatedAt = item.CreatedAt
			updatedItem.UpdatedAt = time.Now()

			watchlist[i] = updatedItem

			// Return the updated item
			response := map[string]interface{}{
				"id":         updatedItem.ID.Hex(),
				"property":   updatedItem.PropertyID.Hex(),
				"user_id":    updatedItem.UserID.Hex(),
				"created_at": updatedItem.CreatedAt,
				"updated_at": updatedItem.UpdatedAt,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// If no item is found, return an error
	http.Error(w, "Watchlist item not found", http.StatusNotFound)
}

// DeleteWatchlist deletes a watchlist item by its ID
func DeleteWatchlist(w http.ResponseWriter, r *http.Request) {
	// Extract the watchlist ID from the URL
	watchlistID := r.URL.Query().Get("id")

	// Convert the watchlist ID from string to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(watchlistID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Iterate through the in-memory watchlist and delete the matching item
	for i, item := range watchlist {
		if item.ID == objectID {
			// Remove the item from the slice
			watchlist = append(watchlist[:i], watchlist[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent) // HTTP 204 No Content
			return
		}
	}

	// If no item is found, return an error
	http.Error(w, "Watchlist item not found", http.StatusNotFound)
}
