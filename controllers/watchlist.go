package controllers

import (
	"encoding/json"
	"net/http"
	"propxchange/schema"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Watchlist []schema.Watchlist

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
	Watchlist = append(Watchlist, watchlistItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(watchlistItem)
}

// ListWatchlist returns all properties in the watchlist
func ListWatchlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Watchlist)
}

// GetWatchlist retrieves all watchlist items or a specific item by ID
func GetWatchlist(w http.ResponseWriter, r *http.Request, id string) {
	// Prepare the response
	var response []map[string]interface{}

	if id == "" {
		// If no ID is provided, retrieve all watchlist items
		for _, item := range Watchlist {
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
	} else {
		// Retrieve a specific watchlist item by ID
		for _, item := range Watchlist {
			if item.ID.Hex() == id {
				responseItem := map[string]interface{}{
					"id":         item.ID.Hex(),
					"property":   item.PropertyID.Hex(),
					"user_id":    item.UserID.Hex(),
					"created_at": item.CreatedAt,
					"updated_at": item.UpdatedAt,
				}
				response = append(response, responseItem)
				break
			}
		}
		// If no matching item is found, return a 404 response
		if len(response) == 0 {
			http.Error(w, "Watchlist item not found", http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateWatchlist updates a watchlist item by its ID
func UpdateWatchlist(w http.ResponseWriter, r *http.Request, watchlistID string) {
	objectID, err := primitive.ObjectIDFromHex(watchlistID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedItem schema.Watchlist
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Iterate through the in-memory watchlist and update the matching item
	for i, item := range Watchlist {
		if item.ID == objectID {
			updatedItem.ID = item.ID
			updatedItem.UserID = item.UserID
			updatedItem.CreatedAt = item.CreatedAt
			updatedItem.UpdatedAt = time.Now()

			Watchlist[i] = updatedItem

			response := map[string]interface{}{
				"id":         updatedItem.ID.Hex(),
				"property":   updatedItem.PropertyID.Hex(),
				"user_id":    updatedItem.UserID.Hex(),
				"created_at": updatedItem.CreatedAt,
				"updated_at": updatedItem.UpdatedAt,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":   "Watchlist updated successfully",
				"watchlist": response,
			})
			return // Exit the function after sending a response
		}
	}

	http.Error(w, "Watchlist item not found", http.StatusNotFound)
}

// DeleteWatchlist deletes a watchlist item by its ID
func DeleteWatchlist(w http.ResponseWriter, r *http.Request, watchlistID string) {
	objectID, err := primitive.ObjectIDFromHex(watchlistID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	for i, item := range Watchlist {
		if item.ID == objectID {
			Watchlist = append(Watchlist[:i], Watchlist[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Watchlist deleted successfully",
			})
			return // Exit the function after sending a response
		}
	}

	http.Error(w, "Watchlist item not found", http.StatusNotFound)
}
