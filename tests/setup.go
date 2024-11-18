package tests

import (
	"backend-crm/routes"
	"context"
	"log"
	"testing"
	"time"

	"backend-crm/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TestDB *mongo.Database
var ctx context.Context
var cancel context.CancelFunc

func SetupRouter() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)     // Register user routes
	routes.CustomerRoutes(r)     // Register customer routes
	routes.TicketRoutes(r)       // Register ticket routes
	routes.NotificationRoutes(r) // Register notification routes
	return r
}

// SetupTestDB initializes a test database connection
func SetupTestDB() {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initialize MongoDB test client and database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to test MongoDB: %v", err)
	}

	TestDB = client.Database("test_crm")
	utils.DB = client // Set this client to the global DB if needed
	log.Println("Connected to Test MongoDB!")
}

// TeardownTestDB cleans up the test database
func TeardownTestDB() {
	if TestDB != nil {
		err := TestDB.Drop(ctx)
		if err != nil {
			log.Fatalf("Failed to drop test database: %v", err)
		}
		log.Println("Test database dropped!")
	}
}

// TestMain is the entry point for testing
func TestMain(m *testing.M) {
	SetupTestDB()
	defer TeardownTestDB()

	m.Run()
}
