package routes

import (
	"backend-crm/controllers"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(router *gin.Engine) {
	notificationRoutes := router.Group("/notifications")
	{
		notificationRoutes.POST("/", controllers.CreateNotification)
		notificationRoutes.GET("/", controllers.GetNotifications)
	}
}
