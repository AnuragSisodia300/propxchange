package routes

import (
	"net/http"
	"propxchange/controllers"
)

func CampaignRoutes() {
	http.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.CreateCampaign(w, r)
		case http.MethodGet:
			controllers.GetCampaigns(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/campaigns/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/campaigns/"):]
		switch r.Method {
		case http.MethodGet:
			controllers.GetCampaignByID(w, r, id)
		case http.MethodPut:
			controllers.UpdateCampaign(w, r, id)
		case http.MethodDelete:
			controllers.DeleteCampaign(w, r, id)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
