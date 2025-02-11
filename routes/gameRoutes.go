package routes

import (
	"Game-Log/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupGameRoutes(r *gin.Engine) {
	games := r.Group("/games")
	{
		games.GET("/", controllers.GetAllGames)
	}
}
