package main

import (
	"Game-Log/backend/database"
	"Game-Log/backend/routes"

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
