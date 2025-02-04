package routes_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"propxchange/controllers"
	"propxchange/schema"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddFavorite(t *testing.T) {
	userID := primitive.NewObjectID()
	propertyID := primitive.NewObjectID()

	// Simulate the request
	req := httptest.NewRequest(http.MethodPost, "/favorites?userId="+userID.Hex()+"&propertyId="+propertyID.Hex(), nil)
	w := httptest.NewRecorder()

	controllers.AddFavorite(w, req)

	// Validate the response
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created; got %v", res.StatusCode)
	}

	var favorite schema.Favorite
	if err := json.NewDecoder(res.Body).Decode(&favorite); err != nil {
		t.Errorf("error decoding response: %v", err)
	}

	if favorite.UserID != userID || favorite.PropertyID != propertyID {
		t.Errorf("expected userID %v and propertyID %v; got userID %v and propertyID %v",
			userID, propertyID, favorite.UserID, favorite.PropertyID)
	}
}

func TestGetFavorites(t *testing.T) {
	userID := primitive.NewObjectID()
	propertyID := primitive.NewObjectID()

	// Add a favorite for testing
	controllers.AddFavorite(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/favorites?userId="+userID.Hex()+"&propertyId="+propertyID.Hex(), nil))

	// Simulate the request
	req := httptest.NewRequest(http.MethodGet, "/favorites/"+userID.Hex(), nil)
	w := httptest.NewRecorder()

	controllers.GetFavorites(w, req, userID.Hex())

	// Validate the response
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	var favorites []schema.Favorite
	if err := json.NewDecoder(res.Body).Decode(&favorites); err != nil {
		t.Errorf("error decoding response: %v", err)
	}

	if len(favorites) != 1 || favorites[0].PropertyID != propertyID {
		t.Errorf("expected 1 favorite with propertyID %v; got %v", propertyID, favorites)
	}
}

func TestRemoveFavorite(t *testing.T) {
	userID := primitive.NewObjectID()
	propertyID := primitive.NewObjectID()

	// Add a favorite for testing
	controllers.AddFavorite(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/favorites?userId="+userID.Hex()+"&propertyId="+propertyID.Hex(), nil))

	// Simulate the request
	req := httptest.NewRequest(http.MethodDelete, "/favorites/"+userID.Hex()+"?propertyId="+propertyID.Hex(), nil)
	w := httptest.NewRecorder()

	controllers.RemoveFavorite(w, req, userID.Hex())

	// Validate the response
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("expected status NoContent; got %v", res.StatusCode)
	}

	// Ensure the favorite is removed
	getReq := httptest.NewRequest(http.MethodGet, "/favorites/"+userID.Hex(), nil)
	getRes := httptest.NewRecorder()
	controllers.GetFavorites(getRes, getReq, userID.Hex())

	var favorites []schema.Favorite
	if err := json.NewDecoder(getRes.Body).Decode(&favorites); err != nil {
		t.Errorf("error decoding response: %v", err)
	}

	if len(favorites) != 0 {
		t.Errorf("expected no favorites; got %v", favorites)
	}
}
