package services

import (
	"dragonball-api/models"
	"dragonball-api/storage"
)

func AddFavorite(character models.Character) (string, error) {
	storage.AddFavorites(character)
	return "Favorite added", nil
}