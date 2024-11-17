package routes

import (
	"backend-crm/controllers"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	ticketRoutes := router.Group("/tickets")
	{
		ticketRoutes.POST("/", controllers.CreateTicket)
		ticketRoutes.GET("/", controllers.GetTickets)
		ticketRoutes.GET("/:id", controllers.GetTicketByID)
		ticketRoutes.PUT("/:id", controllers.UpdateTicket)
		ticketRoutes.DELETE("/:id", controllers.DeleteTicket)
	}
}
