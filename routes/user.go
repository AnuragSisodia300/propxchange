package routes

import (
	"net/http"
	"propxchange/controllers"
)

func UserRoutes() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.ListUsers(w, r) // List all users
		case http.MethodPost:
			controllers.CreateUser(w, r) // Create a new user
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/users/"):]
		switch r.Method {
		case http.MethodGet:
			controllers.GetUser(w, r, id) // Get a user by ID
		case http.MethodPut:
			controllers.UpdateUser(w, r, id) // Update a user
		case http.MethodDelete:
			controllers.DeleteUser(w, r, id) // Delete a user
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
