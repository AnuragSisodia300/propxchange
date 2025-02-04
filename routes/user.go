package routes

import (
	"net/http"
	"propxchange/controllers"
)

func UserRoutes() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.ListUsers(w, r)
		case http.MethodPost:
			controllers.CreateUser(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/users/"):]
		switch r.Method {
		case http.MethodGet:
			controllers.GetUser(w, r, id)
		case http.MethodPut:
			controllers.UpdateUser(w, r, id)
		case http.MethodDelete:
			controllers.DeleteUser(w, r, id)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
