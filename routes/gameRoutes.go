package routes

import (
	"github.com/Abdurahmanit/Game_log/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupGameRoutes(r *gin.Engine) {
	games := r.Group("/games")
	{
		games.GET("/", controllers.GetAllGames)
	}
}
