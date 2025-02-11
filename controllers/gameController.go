package controllers

import (
	"net/http"

	"github.com/Abdurahmanit/Game_log/backend/services"

	"github.com/gin-gonic/gin"
)

func GetAllGames(c *gin.Context) {
	games, err := services.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}
