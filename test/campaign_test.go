package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"propxchange/controllers"
	"propxchange/schema"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test helper to reset the in-memory campaigns data
func resetCampaigns() {
	controllers.Campaigns = []schema.Campaign{}
}

// Test helper to create dummy campaign data
func createDummyCampaign() schema.Campaign {
	campaign := schema.Campaign{
		ID:         primitive.NewObjectID(),
		UserID:     primitive.NewObjectID(),
		PropertyID: primitive.NewObjectID(),
		StartDate:  time.Now(),
		EndDate:    time.Now().AddDate(0, 1, 0),
	}
	controllers.Campaigns = append(controllers.Campaigns, campaign)
	return campaign
}

func TestCreateCampaign(t *testing.T) {
	resetCampaigns()

	// Prepare request payload
	payload := schema.Campaign{
		UserID:     primitive.NewObjectID(),
		PropertyID: primitive.NewObjectID(),
		StartDate:  time.Now(),
		EndDate:    time.Now().AddDate(0, 1, 0),
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/campaigns", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Call the handler
	controllers.CreateCampaign(w, req)

	// Validate the response
	assert.Equal(t, http.StatusCreated, w.Code)

	var response schema.Campaign
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, payload.UserID, response.UserID)
	assert.Equal(t, payload.PropertyID, response.PropertyID)
	assert.NotEmpty(t, response.ID)
	assert.NotEmpty(t, response.CreatedAt)
	assert.NotEmpty(t, response.UpdatedAt)
}

func TestGetCampaigns(t *testing.T) {
	resetCampaigns()
	dummyCampaign := createDummyCampaign()

	req := httptest.NewRequest(http.MethodGet, "/campaigns", nil)
	w := httptest.NewRecorder()

	// Call the handler
	controllers.GetCampaigns(w, req)

	// Validate the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response []schema.Campaign
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, dummyCampaign.ID, response[0].ID)
}

func TestGetCampaignByID(t *testing.T) {
	resetCampaigns()
	dummyCampaign := createDummyCampaign()

	req := httptest.NewRequest(http.MethodGet, "/campaigns/"+dummyCampaign.ID.Hex(), nil)
	w := httptest.NewRecorder()

	// Call the handler
	controllers.GetCampaignByID(w, req, dummyCampaign.ID.Hex())

	// Validate the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response schema.Campaign
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, dummyCampaign.ID, response.ID)
}

func TestGetCampaignByID_NotFound(t *testing.T) {
	resetCampaigns()

	req := httptest.NewRequest(http.MethodGet, "/campaigns/000000000000000000000000", nil)
	w := httptest.NewRecorder()

	// Call the handler
	controllers.GetCampaignByID(w, req, "000000000000000000000000")

	// Validate the response
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Campaign not found", response["error"])
}

func TestUpdateCampaign(t *testing.T) {
	resetCampaigns()
	dummyCampaign := createDummyCampaign()

	// Prepare updated data
	updatedCampaign := schema.Campaign{
		UserID:     dummyCampaign.UserID,
		PropertyID: dummyCampaign.PropertyID,
		StartDate:  time.Now(),
		EndDate:    time.Now().AddDate(0, 2, 0),
	}
	body, _ := json.Marshal(updatedCampaign)

	req := httptest.NewRequest(http.MethodPut, "/campaigns/"+dummyCampaign.ID.Hex(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Call the handler
	controllers.UpdateCampaign(w, req, dummyCampaign.ID.Hex())

	// Validate the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response schema.Campaign
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, updatedCampaign.EndDate, response.EndDate)
	assert.Equal(t, dummyCampaign.ID, response.ID)
}

func TestUpdateCampaign_NotFound(t *testing.T) {
	resetCampaigns()

	updatedCampaign := schema.Campaign{
		UserID:     primitive.NewObjectID(),
		PropertyID: primitive.NewObjectID(),
		StartDate:  time.Now(),
		EndDate:    time.Now().AddDate(0, 1, 0),
	}
	body, _ := json.Marshal(updatedCampaign)

	req := httptest.NewRequest(http.MethodPut, "/campaigns/000000000000000000000000", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Call the handler
	controllers.UpdateCampaign(w, req, "000000000000000000000000")

	// Validate the response
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Campaign not found", response["error"])
}

func TestDeleteCampaign(t *testing.T) {
	resetCampaigns()
	dummyCampaign := createDummyCampaign()

	req := httptest.NewRequest(http.MethodDelete, "/campaigns/"+dummyCampaign.ID.Hex(), nil)
	w := httptest.NewRecorder()

	// Call the handler
	controllers.DeleteCampaign(w, req, dummyCampaign.ID.Hex())

	// Validate the response
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Len(t, controllers.Campaigns, 0)
}

func TestDeleteCampaign_NotFound(t *testing.T) {
	resetCampaigns()

	req := httptest.NewRequest(http.MethodDelete, "/campaigns/000000000000000000000000", nil)
	w := httptest.NewRecorder()

	// Call the handler
	controllers.DeleteCampaign(w, req, "000000000000000000000000")

	// Validate the response
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Campaign not found", response["error"])
}
