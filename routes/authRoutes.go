package routes

import (
	"github.com/Abdurahmanit/Game_log/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
