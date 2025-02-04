package main

import (
	"log"
	"net/http"
	"propxchange/routes"
)

func main() {
	routes.UserRoutes()
	routes.PaymentRoutes()
	routes.FavoriteRoutes()
	routes.CampaignRoutes()
	routes.KYCProcessRoutes()
	routes.WatchlistRoutes()
	routes.PropertyRoutes()
	routes.WalletRoutes()
	routes.WebhookRoutes()
	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
