package routes

import (
	"Game-Log/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(r *gin.Engine) {
	orders := r.Group("/orders")
	{
		orders.POST("/create", controllers.CreateOrder)
		orders.GET("/:userID", controllers.GetOrders)
	}
}
