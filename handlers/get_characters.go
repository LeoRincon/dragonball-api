package handlers

import (
	"dragonball-api/mappers"
	"dragonball-api/models"
	"dragonball-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCharacter(c *gin.Context) (models.Response, error) {
	limit, err := strconv.Atoi(c.Query("limit"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return models.Response{}, err
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return models.Response{}, err
	}
	response, err := services.GetCharacter(limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch characters"})
		return models.Response{}, err
	}

	resp := mappers.MapResponse(response)

	return resp, nil
}

func GetCharacterById(c *gin.Context) (models.CharacterById, error) {
	id := c.Param("id")
	response, err := services.GetCharacterById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character by id"})
		return models.CharacterById{}, err
	}
	return response, nil
}