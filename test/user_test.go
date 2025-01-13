package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"propxchange/controllers"
	"propxchange/schema"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestListUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	controllers.ListUsers(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	var users []schema.User
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		t.Errorf("error decoding response: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	newUser := schema.User{
		Name:          "Test User",
		Email:         "test@example.com",
		WalletBalance: 100.50,
	}
	body, _ := json.Marshal(newUser)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	w := httptest.NewRecorder()

	controllers.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created; got %v", res.StatusCode)
	}

	var user schema.User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		t.Errorf("error decoding response: %v", err)
	}

	if user.Name != newUser.Name || user.Email != newUser.Email {
		t.Errorf("expected user %v; got %v", newUser, user)
	}
}

func TestGetUser(t *testing.T) {
	testID := primitive.NewObjectID().Hex()
	req := httptest.NewRequest(http.MethodGet, "/users/"+testID, nil)
	w := httptest.NewRecorder()

	controllers.GetUser(w, req, testID)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Not Found; got %v", res.StatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	testID := primitive.NewObjectID().Hex()
	updatedUser := schema.User{
		Name:          "Updated User",
		Email:         "updated@example.com",
		WalletBalance: 200.75,
	}
	body, _ := json.Marshal(updatedUser)

	req := httptest.NewRequest(http.MethodPut, "/users/"+testID, bytes.NewReader(body))
	w := httptest.NewRecorder()

	controllers.UpdateUser(w, req, testID)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Not Found; got %v", res.StatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	testID := primitive.NewObjectID().Hex()
	req := httptest.NewRequest(http.MethodDelete, "/users/"+testID, nil)
	w := httptest.NewRecorder()

	controllers.DeleteUser(w, req, testID)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Not Found; got %v", res.StatusCode)
	}
}
