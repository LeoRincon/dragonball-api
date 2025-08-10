package routers

import (
	"dragonball-api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)


func SetupV1Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	
	v1.GET("/characters", func(c *gin.Context) {
		response, err := handlers.GetCharacter(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch characters"})
			return
		}
		c.JSON(http.StatusOK, response)
	}) 
	
	v1.GET("/favorites", handlers.GetFavorites)
	
	v1.POST("/favorites", func(context *gin.Context) {
    handlers.AddFavorites(context) 
  })
	
}
