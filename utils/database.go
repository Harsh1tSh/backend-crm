package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

// ConnectDB initializes the global DB client for the production environment
func ConnectDB() {
	if DB != nil {
		// Prevent reconnecting if already connected
		return
	}

	// Define MongoDB URI
	uri := "mongodb://localhost:27017" // Update this if using MongoDB Atlas

	// Create MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	DB = client
	log.Println("Connected to MongoDB!")
}
