package routes

import (
	"net/http"
	"propxchange/controllers"
)

func KYCProcessRoutes() {
	http.HandleFunc("/kyc/step1", controllers.AddKYCStep1)
	http.HandleFunc("/kyc/step2", controllers.AddKYCStep2)
	http.HandleFunc("/kyc/step3", controllers.AddKYCStep3)
}
