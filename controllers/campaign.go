package controllers

import (
	"encoding/json"
	"net/http"
	"propxchange/schema"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Campaigns []schema.Campaign

// ErrorResponse defines the structure of error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateCampaign handles the request to create a new campaign
func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var newCampaign schema.Campaign
	if err := json.NewDecoder(r.Body).Decode(&newCampaign); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request payload"})
		return
	}

	newCampaign.ID = primitive.NewObjectID()
	newCampaign.CreatedAt = time.Now()
	newCampaign.UpdatedAt = time.Now()

	// Add to campaigns list
	Campaigns = append(Campaigns, newCampaign)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCampaign)
}

// GetCampaigns handles the request to get all campaigns
func GetCampaigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Campaigns)
}

// GetCampaignByID handles the request to get a campaign by ID
func GetCampaignByID(w http.ResponseWriter, r *http.Request, id string) {
	campaignID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid campaign ID"})
		return
	}

	for _, campaign := range Campaigns {
		if campaign.ID == campaignID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(campaign)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Campaign not found"})
}

// UpdateCampaign handles the request to update an existing campaign
func UpdateCampaign(w http.ResponseWriter, r *http.Request, id string) {
	campaignID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid campaign ID"})
		return
	}

	var updatedCampaign schema.Campaign
	if err := json.NewDecoder(r.Body).Decode(&updatedCampaign); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request payload"})
		return
	}

	for i, campaign := range Campaigns {
		if campaign.ID == campaignID {
			updatedCampaign.ID = campaignID
			updatedCampaign.CreatedAt = campaign.CreatedAt
			updatedCampaign.UpdatedAt = time.Now()
			Campaigns[i] = updatedCampaign

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCampaign)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Campaign not found"})
}

// DeleteCampaign handles the request to delete a campaign
func DeleteCampaign(w http.ResponseWriter, r *http.Request, id string) {
	campaignID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid campaign ID"})
		return
	}

	for i, campaign := range Campaigns {
		if campaign.ID == campaignID {
			Campaigns = append(Campaigns[:i], Campaigns[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Campaign not found"})
}
