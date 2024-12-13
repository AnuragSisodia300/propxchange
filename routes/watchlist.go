package routes

import (
	"net/http"
	"propxchange/controllers"
)

func WatchlistRoutes() {
	http.HandleFunc("/watchlist", WatchlistHandler)
	http.HandleFunc("/watchlist/", WatchlistItemHandler)
}
func WatchlistHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controllers.AddToWatchlist(w, r)
	case http.MethodGet:
		controllers.ListWatchlist(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func WatchlistItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetWatchlist(w, r)
	case http.MethodPut:
		controllers.UpdateWatchlist(w, r)
	case http.MethodDelete:
		controllers.DeleteWatchlist(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
