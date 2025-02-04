package routes

import (
	"net/http"
	"propxchange/controllers"
)

func WalletRoutes() {
	http.HandleFunc("/wallet/", WalletHandler)

}
func WalletHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		action := r.URL.Query().Get("action")
		if action == "add" {
			controllers.AddMoneyToWallet(w, r)
		} else if action == "purchase" {
			controllers.PurchaseUnitsFromWallet(w, r)
		} else if action == "reward" {
			controllers.AddMoneyToWallet(w, r)
		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
