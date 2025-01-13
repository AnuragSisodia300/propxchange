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

func TestCreateProperty(t *testing.T) {
	newProperty := schema.Property{
		Name:         "Test Property",
		Description:  "This is a test property.",
		Price:        500000.00,
		Location:     "New York",
		Bedrooms:     3,
		Bathrooms:    2,
		SquareFeet:   1500.0,
		OwnerID:      primitive.NewObjectID(),
		Types:        []string{"Residential", "Commercial"},
		MinIRR:       8.0,
		MaxIRR:       12.0,
		PerUnitPrice: 5000.0,
	}

	body, _ := json.Marshal(newProperty)

	req := httptest.NewRequest(http.MethodPost, "/properties", bytes.NewReader(body))
	w := httptest.NewRecorder()

	controllers.CreateProperty(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created; got %v", res.StatusCode)
	}

	var property schema.Property
	if err := json.NewDecoder(res.Body).Decode(&property); err != nil {
		t.Errorf("error decoding response: %v", err)
	}

	// Check if the created property matches the sent data
	if property.Name != newProperty.Name || property.Description != newProperty.Description {
		t.Errorf("expected property %v; got %v", newProperty, property)
	}
}

func TestListProperties(t *testing.T) {
	httptest.NewRequest(http.MethodGet, "/properties", nil)
	w := httptest.NewRecorder()

	controllers.ListProperties(w)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}
	var properties []schema.Property

	if err := json.NewDecoder(res.Body).Decode(&properties); err != nil {
		t.Errorf("error decoding response: %v", err)
	}
}

func TestGetProperty(t *testing.T) {
	// Create a property to retrieve
	propertyID := primitive.NewObjectID()
	properties := []schema.Property{
		{
			ID:           propertyID,
			Name:         "Test Property",
			Description:  "This is a test property.",
			Price:        500000.00,
			Location:     "New York",
			Bedrooms:     3,
			Bathrooms:    2,
			SquareFeet:   1500.0,
			OwnerID:      primitive.NewObjectID(),
			Types:        []string{"Residential", "Commercial"},
			MinIRR:       8.0,
			MaxIRR:       12.0,
			PerUnitPrice: 5000.0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	// Simulating adding the property into the in-memory store
	controllers.Properties = append(controllers.Properties, properties[0])

	httptest.NewRequest(http.MethodGet, "/properties/"+propertyID.Hex(), nil)
	w := httptest.NewRecorder()

	controllers.GetProperty(w, propertyID.Hex())

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	var property schema.Property
	if err := json.NewDecoder(res.Body).Decode(&property); err != nil {
		t.Errorf("error decoding response: %v", err)
	}

	if property.ID != propertyID {
		t.Errorf("expected property ID %v; got %v", propertyID, property.ID)
	}
}

func TestUpdateProperty(t *testing.T) {
	// Create a property to update
	propertyID := primitive.NewObjectID()
	properties := []schema.Property{
		{
			ID:           propertyID,
			Name:         "Test Property",
			Description:  "This is a test property.",
			Price:        500000.00,
			Location:     "New York",
			Bedrooms:     3,
			Bathrooms:    2,
			SquareFeet:   1500.0,
			OwnerID:      primitive.NewObjectID(),
			Types:        []string{"Residential", "Commercial"},
			MinIRR:       8.0,
			MaxIRR:       12.0,
			PerUnitPrice: 5000.0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	// Simulating adding the property into the in-memory store
	controllers.Properties = append(controllers.Properties, properties[0])

	updatedProperty := schema.Property{
		Name:        "Updated Property",
		Description: "Updated description.",
		Price:       600000.00,
		Location:    "Los Angeles",
		Bedrooms:    4,
		Bathrooms:   3,
		SquareFeet:  2000.0,
	}

	body, _ := json.Marshal(updatedProperty)

	req := httptest.NewRequest(http.MethodPut, "/properties/"+propertyID.Hex(), bytes.NewReader(body))
	w := httptest.NewRecorder()

	controllers.UpdateProperty(w, req, propertyID.Hex())

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	var property schema.Property
	if err := json.NewDecoder(res.Body).Decode(&property); err != nil {
		t.Errorf("error decoding response: %v", err)
	}

	if property.Name != updatedProperty.Name || property.Description != updatedProperty.Description {
		t.Errorf("expected updated property %v; got %v", updatedProperty, property)
	}
}

func TestDeleteProperty(t *testing.T) {
	// Create a property to delete
	propertyID := primitive.NewObjectID()
	properties := []schema.Property{
		{
			ID:           propertyID,
			Name:         "Test Property",
			Description:  "This is a test property.",
			Price:        500000.00,
			Location:     "New York",
			Bedrooms:     3,
			Bathrooms:    2,
			SquareFeet:   1500.0,
			OwnerID:      primitive.NewObjectID(),
			Types:        []string{"Residential", "Commercial"},
			MinIRR:       8.0,
			MaxIRR:       12.0,
			PerUnitPrice: 5000.0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	// Simulating adding the property into the in-memory store
	controllers.Properties = append(controllers.Properties, properties[0])

	httptest.NewRequest(http.MethodDelete, "/properties/"+propertyID.Hex(), nil)
	w := httptest.NewRecorder()

	controllers.DeleteProperty(w, propertyID.Hex())

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("expected status No Content; got %v", res.StatusCode)
	}

	// Check if the property is deleted
	for _, property := range controllers.Properties {
		if property.ID == propertyID {
			t.Errorf("expected property to be deleted, but found %v", property)
		}
	}
}
