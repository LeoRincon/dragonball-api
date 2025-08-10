package handlers

import (
	"dragonball-api/models"
	"dragonball-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFavorites(c *gin.Context) {
	var request models.Character

	err := c.ShouldBindJSON(&request);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	
	message, err := services.AddFavorite(request)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add favorite"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}