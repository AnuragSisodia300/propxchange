package routes

import (
	"net/http"
	"propxchange/controllers"
)

func WatchlistRoutes() {
	http.HandleFunc("/watchlist", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.AddToWatchlist(w, r) // Add an item to the watchlist
		case http.MethodGet:
			controllers.ListWatchlist(w, r) // List all items in the watchlist
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/watchlist/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/watchlist/"):]
		switch r.Method {
		case http.MethodGet:
			controllers.GetWatchlist(w, r, id) // Get a specific watchlist item by ID
		case http.MethodPut:
			controllers.UpdateWatchlist(w, r, id) // Update a specific watchlist item by ID
		case http.MethodDelete:
			controllers.DeleteWatchlist(w, r, id) // Delete a specific watchlist item by ID
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
