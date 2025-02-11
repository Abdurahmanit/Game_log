package controllers

import (
	"Game-Log/backend/services"
	"net/http"

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
