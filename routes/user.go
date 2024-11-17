package routes

import (
	"backend-crm/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
	}
}
