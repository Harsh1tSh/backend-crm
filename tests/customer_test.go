package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateCustomer(t *testing.T) {
	router := SetupRouter()

	// Clean up the "customers" collection before running the test
	TestDB.Collection("customers").DeleteMany(context.TODO(), bson.M{})

	payload := map[string]interface{}{
		"name":    "Customer A",
		"email":   "customer@example.com",
		"company": "Example Corp",
		"status":  "Active",
		"notes":   "Priority customer",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/customers/", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Customer A")
}

func TestGetCustomers(t *testing.T) {
	router := SetupRouter()

	// Seed the "customers" collection
	TestDB.Collection("customers").InsertOne(context.TODO(), bson.M{
		"name":    "Customer A",
		"email":   "customer@example.com",
		"company": "Example Corp",
		"status":  "Active",
		"notes":   "Priority customer",
	})

	req, _ := http.NewRequest("GET", "/customers/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Customer A")
}
