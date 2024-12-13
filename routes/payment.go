package routes

import (
	"net/http"
	"propxchange/controllers"
)

func PaymentRoutes() {
	http.HandleFunc("/payments", controllers.HandlePayment)
	http.HandleFunc("/purchase", controllers.PurchaseUnits)
}
