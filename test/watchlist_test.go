package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"propxchange/controllers"
	"propxchange/schema"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddToWatchlist(t *testing.T) {
	// Prepare a watchlist item
	watchlistItem := schema.Watchlist{
		UserID:     primitive.NewObjectID(),
		PropertyID: primitive.NewObjectID(),
	}
	body, _ := json.Marshal(watchlistItem)

	req := httptest.NewRequest(http.MethodPost, "/watchlist", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	controllers.AddToWatchlist(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, res.Code)
	}

	var createdItem schema.Watchlist
	_ = json.NewDecoder(res.Body).Decode(&createdItem)

	if createdItem.UserID != watchlistItem.UserID {
		t.Errorf("Expected UserID %v, got %v", watchlistItem.UserID, createdItem.UserID)
	}
}

func TestListWatchlist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/watchlist", nil)
	res := httptest.NewRecorder()

	controllers.ListWatchlist(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, res.Code)
	}

	var items []schema.Watchlist
	_ = json.NewDecoder(res.Body).Decode(&items)

	if len(items) == 0 {
		t.Error("Expected non-empty watchlist, got empty")
	}
}

func TestGetWatchlist(t *testing.T) {
	// Create a mock watchlist item
	watchlistItem := schema.Watchlist{
		ID:         primitive.NewObjectID(),
		UserID:     primitive.NewObjectID(),
		PropertyID: primitive.NewObjectID(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Add the item to the mock watchlist (simulate in-memory storage)
	controllers.Watchlist = append(controllers.Watchlist, watchlistItem)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/watchlist/"+watchlistItem.ID.Hex(), nil)
	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	controllers.GetWatchlist(rr, req, watchlistItem.ID.Hex())

	// Assert the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body
	var response []map[string]interface{}
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	// Assert the returned data
	if len(response) == 0 {
		t.Fatalf("expected at least one item in the response, got %d", len(response))
	}

	if response[0]["id"] != watchlistItem.ID.Hex() {
		t.Errorf("handler returned unexpected item ID: got %v want %v", response[0]["id"], watchlistItem.ID.Hex())
	}
}

func TestUpdateWatchlist(t *testing.T) {
	// Create a mock watchlist item
	watchlistItem := schema.Watchlist{
		ID:         primitive.NewObjectID(),
		UserID:     primitive.NewObjectID(),
		PropertyID: primitive.NewObjectID(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Simulate adding the watchlist item to the in-memory storage
	controllers.Watchlist = append(controllers.Watchlist, watchlistItem)

	// Create a payload for updating the watchlist item
	updatePayload := map[string]interface{}{
		"userId":     watchlistItem.UserID.Hex(),
		"propertyId": primitive.NewObjectID().Hex(),
	}
	payloadBytes, _ := json.Marshal(updatePayload)

	// Create a new HTTP PUT request
	req := httptest.NewRequest(http.MethodPut, "/watchlist/"+watchlistItem.ID.Hex(), bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	controllers.UpdateWatchlist(rr, req, watchlistItem.ID.Hex())

	// Assert the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode and validate the response body
	var response map[string]interface{}
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	if response["message"] != "Watchlist updated successfully" {
		t.Errorf("unexpected response message: got %v want %v", response["message"], "Watchlist updated successfully")
	}
}

func TestDeleteWatchlist(t *testing.T) {
	// Create a mock watchlist item
	watchlistItem := schema.Watchlist{
		ID:         primitive.NewObjectID(),
		UserID:     primitive.NewObjectID(),
		PropertyID: primitive.NewObjectID(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Simulate adding the watchlist item to the in-memory storage
	controllers.Watchlist = append(controllers.Watchlist, watchlistItem)

	// Create a DELETE request
	req := httptest.NewRequest(http.MethodDelete, "/watchlist/"+watchlistItem.ID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	controllers.DeleteWatchlist(rr, req, watchlistItem.ID.Hex())

	// Assert the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode and validate the response body
	var response map[string]interface{}
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	if response["message"] != "Watchlist deleted successfully" {
		t.Errorf("unexpected response message: got %v want %v", response["message"], "Watchlist deleted successfully")
	}
}
