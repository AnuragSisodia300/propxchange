package routes

import (
	"net/http"
	"propxchange/controllers"
)

func FavoriteRoutes() {
	http.HandleFunc("/favorites", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.AddFavorite(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/favorites/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/favorites/"):]
		switch r.Method {
		case http.MethodGet:
			controllers.GetFavorites(w, r, id)
		case http.MethodDelete:
			controllers.RemoveFavorite(w, r, id)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
