package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"propxchange/schema"

	"cloud.google.com/go/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Step 1: Personal and Residential Info
func AddKYCStep1(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")

	// Parse user ID
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req struct {
		PersonalInfo    schema.PersonalInfo    `json:"personal_info"`
		ResidentialInfo schema.ResidentialInfo `json:"residential_info"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Update the user's KYC info
	for i, user := range users {
		if user.ID == uID {
			users[i].KYC.PersonalInfo = req.PersonalInfo
			users[i].KYC.ResidentialInfo = req.ResidentialInfo
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users[i].KYC)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// Step 2: Financial Info
func AddKYCStep2(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")

	// Parse user ID
	uID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var financialInfo schema.FinancialInfo
	if err := json.NewDecoder(r.Body).Decode(&financialInfo); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Update the user's KYC info
	for i, user := range users {
		if user.ID == uID {
			users[i].KYC.FinancialInfo = financialInfo
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users[i].KYC)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

const (
	bucketName   = "wdl-varun"
	folderName   = "propxchange"
	credFilePath = "./config/woodland-397213-420bc2b011a3.json"
)

// AddKYCStep3 handles the upload of KYC documents (ID proof, Address proof, Income proof).
func AddKYCStep3(w http.ResponseWriter, r *http.Request) {
	// Set GCP credentials explicitly
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credFilePath)
	if err != nil {
		http.Error(w, "Failed to set GCP credentials", http.StatusInternalServerError)
		return
	}

	// Parse the multipart form with a maximum memory of 10 MB
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to process form data", http.StatusBadRequest)
		return
	}

	// Get user ID from the form
	userID := r.FormValue("userId")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Validate user ID
	_, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Connect to GCP bucket
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		http.Error(w, "Failed to connect to GCP storage", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Define a helper function to upload a file
	uploadFile := func(fileKey string) (string, error) {
		file, header, err := r.FormFile(fileKey)
		if err != nil {
			return "", fmt.Errorf("error retrieving %s file: %v", fileKey, err)
		}
		defer file.Close()

		// Generate file name using user ID and file key
		fileName := fmt.Sprintf("%s/kyc/%s/%s_%s", folderName, userID, fileKey, filepath.Base(header.Filename))

		// Upload the file to the GCP bucket
		bucket := client.Bucket(bucketName)
		obj := bucket.Object(fileName)
		writer := obj.NewWriter(ctx)
		defer writer.Close()

		if _, err := io.Copy(writer, file); err != nil {
			return "", fmt.Errorf("failed to upload %s: %v", fileKey, err)
		}

		return fileName, nil
	}

	// Upload each document
	idProofPath, err := uploadFile("id_proof")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addressProofPath, err := uploadFile("address_proof")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	incomeProofPath, err := uploadFile("income_proof")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with success and file paths
	response := map[string]string{
		"id_proof":      idProofPath,
		"address_proof": addressProofPath,
		"income_proof":  incomeProofPath,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
