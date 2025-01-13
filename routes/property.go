package routes

import (
	"net/http"
	"propxchange/controllers"
)

func PropertyRoutes() {
	http.HandleFunc("/properties", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.ListProperties(w)
		case http.MethodPost:
			controllers.CreateProperty(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/properties/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/properties/"):]
		switch r.Method {
		case http.MethodGet:
			controllers.GetProperty(w, id)
		case http.MethodPut:
			controllers.UpdateProperty(w, r, id)
		case http.MethodDelete:
			controllers.DeleteProperty(w, id)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/properties/purchase", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.PurchaseUnits(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
