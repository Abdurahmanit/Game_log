package routes

import (
	"github.com/Abdurahmanit/Game_log/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(r *gin.Engine) {
	cart := r.Group("/cart")
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("/:userID", controllers.GetCart)
	}
}
