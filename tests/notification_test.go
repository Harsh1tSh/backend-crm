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

func TestCreateNotification(t *testing.T) {
	router := SetupRouter()

	// Clean up the "notifications" collection before running the test
	TestDB.Collection("notifications").DeleteMany(context.TODO(), bson.M{})

	payload := map[string]interface{}{
		"message": "Meeting scheduled at 3 PM",
		"user_id": "643a14ab12f6558342a5c2e9", // Replace with a valid ObjectID in string format
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/notifications/", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Meeting scheduled at 3 PM")
}

func TestGetNotifications(t *testing.T) {
	router := SetupRouter()

	// Seed the "notifications" collection
	TestDB.Collection("notifications").InsertOne(context.TODO(), bson.M{
		"message": "Meeting scheduled at 3 PM",
		"user_id": "643a14ab12f6558342a5c2e9",
	})

	req, _ := http.NewRequest("GET", "/notifications/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Meeting scheduled at 3 PM")
}
