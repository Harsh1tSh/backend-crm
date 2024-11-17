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

// CreateCustomer handles creating a new customer
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	customerCollection := utils.DB.Database("crm").Collection("customers")

	_, err := customerCollection.InsertOne(ctx, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// GetCustomers handles fetching all customers
func GetCustomers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	customerCollection := utils.DB.Database("crm").Collection("customers")

	cursor, err := customerCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}
	defer cursor.Close(ctx)

	var customers []models.Customer
	for cursor.Next(ctx) {
		var customer models.Customer
		if err := cursor.Decode(&customer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding customer data"})
			return
		}
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
}

// GetCustomerByID handles fetching a single customer by their ID
func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	customerCollection := utils.DB.Database("crm").Collection("customers")

	var customer models.Customer
	err = customerCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UpdateCustomer handles updating an existing customer's details
func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	customerCollection := utils.DB.Database("crm").Collection("customers")

	update := bson.M{"$set": customer}
	_, err = customerCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

// DeleteCustomer handles deleting a customer by their ID
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	customerCollection := utils.DB.Database("crm").Collection("customers")

	_, err = customerCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
