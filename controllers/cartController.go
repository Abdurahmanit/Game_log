package controllers

import (
	"net/http"

	"github.com/Abdurahmanit/Game_log/backend/model"
	"github.com/Abdurahmanit/Game_log/backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddToCart(c *gin.Context) {
	var cart model.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddToCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game added to cart successfully"})
}

func GetCart(c *gin.Context) {
	userID := c.Param("userID")
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	cart, err := services.GetCartByUserID(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}
