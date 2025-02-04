package controllers

import (
	"encoding/json"
	"net/http"
	"propxchange/schema"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var favorites []schema.Favorite

// AddFavorite handles the request to add a property to the user's favorites
func AddFavorite(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	propertyID := r.URL.Query().Get("propertyId")

	// Parse user ID and property ID
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	pID, err := primitive.ObjectIDFromHex(propertyID)
	if err != nil {
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}

	// Check if property is already in favorites
	for _, fav := range favorites {
		if fav.UserID == uID && fav.PropertyID == pID {
			http.Error(w, "Property already in favorites", http.StatusBadRequest)
			return
		}
	}

	// Add to favorites
	newFavorite := schema.Favorite{
		UserID:     uID,
		PropertyID: pID,
	}
	favorites = append(favorites, newFavorite)

	for i, u := range users {
		if u.ID == uID {
			users[i].Favorites = append(users[i].Favorites, pID)
			break
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newFavorite)
}

// GetFavorites handles the request to get all properties marked as favorites by a user
func GetFavorites(w http.ResponseWriter, r *http.Request, userID string) {
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var userFavorites []schema.Favorite
	for _, fav := range favorites {
		if fav.UserID == uID {
			userFavorites = append(userFavorites, fav)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userFavorites)
}

// RemoveFavorite handles the request to remove a property from a user's favorites
func RemoveFavorite(w http.ResponseWriter, r *http.Request, userID string) {
	propertyID := r.URL.Query().Get("propertyId")

	// Parse user ID and property ID
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	pID, err := primitive.ObjectIDFromHex(propertyID)
	if err != nil {
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}

	// Find and remove the favorite
	for i, fav := range favorites {
		if fav.UserID == uID && fav.PropertyID == pID {
			favorites = append(favorites[:i], favorites[i+1:]...)
			break
		}
	}

	// Update user's favorites array
	for i, u := range users {
		if u.ID == uID {
			for j, id := range u.Favorites {
				if id == pID {
					users[i].Favorites = append(users[i].Favorites[:j], users[i].Favorites[j+1:]...)
					break
				}
			}
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
