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

func TestCreateTicket(t *testing.T) {
	router := SetupRouter()

	// Clean up the "tickets" collection before running the test
	TestDB.Collection("tickets").DeleteMany(context.TODO(), bson.M{})

	payload := map[string]interface{}{
		"customer_id": "643a14ab12f6558342a5c2e9", // Replace with a valid ObjectID in string format
		"title":       "Issue with product",
		"description": "Product not working",
		"status":      "Open",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/tickets/", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Issue with product")
}

func TestGetTickets(t *testing.T) {
	router := SetupRouter()

	// Seed the "tickets" collection
	TestDB.Collection("tickets").InsertOne(context.TODO(), bson.M{
		"customer_id": "643a14ab12f6558342a5c2e9",
		"title":       "Issue with product",
		"description": "Product not working",
		"status":      "Open",
	})

	req, _ := http.NewRequest("GET", "/tickets/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Issue with product")
}
