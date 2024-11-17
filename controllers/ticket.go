package controllers

import (
	"backend-crm/models"
	"backend-crm/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.ID = primitive.NewObjectID()
	ticket.CreatedAt = time.Now().Unix()
	ticket.UpdatedAt = ticket.CreatedAt

	// Initialize the collection dynamically
	ticketCollection := utils.DB.Database("crm").Collection("tickets")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := ticketCollection.InsertOne(ctx, ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func GetTickets(c *gin.Context) {
	// Initialize the collection dynamically
	ticketCollection := utils.DB.Database("crm").Collection("tickets")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := ticketCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
		return
	}
	defer cursor.Close(ctx)

	var tickets []models.Ticket
	for cursor.Next(ctx) {
		var ticket models.Ticket
		cursor.Decode(&ticket)
		tickets = append(tickets, ticket)
	}

	c.JSON(http.StatusOK, tickets)
}

func GetTicketByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	// Initialize the collection dynamically
	ticketCollection := utils.DB.Database("crm").Collection("tickets")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ticket models.Ticket
	err = ticketCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&ticket)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func UpdateTicket(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.UpdatedAt = time.Now().Unix()

	// Initialize the collection dynamically
	ticketCollection := utils.DB.Database("crm").Collection("tickets")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"$set": ticket}
	_, err = ticketCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
}

func DeleteTicket(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	// Initialize the collection dynamically
	ticketCollection := utils.DB.Database("crm").Collection("tickets")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = ticketCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
