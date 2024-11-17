package routes

import (
	"backend-crm/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(router *gin.Engine) {
	customerRoutes := router.Group("/customers")
	{
		customerRoutes.POST("/", controllers.CreateCustomer)
		customerRoutes.GET("/", controllers.GetCustomers)
		customerRoutes.GET("/:id", controllers.GetCustomerByID)
		customerRoutes.PUT("/:id", controllers.UpdateCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomer)
	}
}
