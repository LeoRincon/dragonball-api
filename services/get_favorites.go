package services

import (
	"dragonball-api/models"
	"dragonball-api/storage"
)

func GetFavorites() ([]models.Character, error) {
	response := storage.GetAllFavorites()
	return response, nil
}