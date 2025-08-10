package handlers

import (
	"dragonball-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	response, err := services.GetFavorites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch favorites"})
		return
	}
	c.JSON(http.StatusOK, response)
}