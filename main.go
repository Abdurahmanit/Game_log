package main

import (
	"github.com/Abdurahmanit/Game_log/backend/database"
	"github.com/Abdurahmanit/Game_log/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	routes.SetupAuthRoutes(r)
	routes.SetupGameRoutes(r)
	routes.SetupCartRoutes(r)
	routes.SetupOrderRoutes(r)

	r.Run(":8080")
}
