package routes

import (
	"net/http"
	"propxchange/controllers"
)

func PropertyRoutes() {
	http.HandleFunc("/properties", PropertiesHandler)
	http.HandleFunc("/properties/", PropertyHandler)
	http.HandleFunc("/properties/purchase", PurchaseUnitsHandler)
}

// PropertiesHandler handles the /properties endpoint (GET and POST)
func PropertiesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.ListProperties(w)
	case http.MethodPost:
		controllers.CreateProperty(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// PropertyHandler handles the /properties/{id} endpoint (GET, PUT, DELETE)
func PropertyHandler(w http.ResponseWriter, r *http.Request) {
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
}
func PurchaseUnitsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	controllers.PurchaseUnits(w, r)
}
