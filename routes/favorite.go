package routes

import (
	"net/http"

	"propxchange/controllers"
)

func FavoriteRoutes() {
	http.HandleFunc("/favorites", FavoriteHandler)
	http.HandleFunc("/favorites/", FavoriteItemHandler)
}

func FavoriteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controllers.AddFavorite(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func FavoriteItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetFavorites(w, r)
	case http.MethodDelete:
		controllers.RemoveFavorite(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
