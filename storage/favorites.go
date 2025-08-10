package storage

import "dragonball-api/models"

var storage = make([]models.Character, 0)

func AddFavorites(character models.Character) {
	storage = append(storage, character)
}

func GetFavorites(id int) (models.Character, bool) {
	for _, character := range storage {
		if character.ID == id {
			return character, true
		}
	}
	return models.Character{}, false
}

func RemoveFavorites(id int) {
	for i, character := range storage {
		if character.ID == id {
			storage = append(storage[:i], storage[i+1:]...)
			break
		}
	}
}

func GetAllFavorites() []models.Character {
	return storage
}

func UpdateFavorites(character models.Character) {
	for i, c := range storage {
		if c.ID == character.ID {
			storage[i] = character
			break
		}
	}
}