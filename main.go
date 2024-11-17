package main

import (
	"backend-crm/routes"
	"backend-crm/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	utils.ConnectDB()

	// Set up the router
	r := gin.Default()
	routes.RegisterRoutes(r)
	routes.CustomerRoutes(r)     // Add customer Routes
	routes.TicketRoutes(r)       // Add Ticket Routes
	routes.NotificationRoutes(r) // Add Notifications Routs

	// Start the server
	r.Run(":8080")
}
