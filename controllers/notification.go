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

func CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification.ID = primitive.NewObjectID()
	notification.CreatedAt = time.Now().Unix()

	notificationCollection := utils.DB.Database("crm").Collection("notifications")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := notificationCollection.InsertOne(ctx, notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	c.JSON(http.StatusOK, notification)
}

func GetNotifications(c *gin.Context) {
	notificationCollection := utils.DB.Database("crm").Collection("notifications")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := notificationCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}
	defer cursor.Close(ctx)

	var notifications []models.Notification
	for cursor.Next(ctx) {
		var notification models.Notification
		cursor.Decode(&notification)
		notifications = append(notifications, notification)
	}

	c.JSON(http.StatusOK, notifications)
}
