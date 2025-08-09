package services

import (
	"dragonball-api/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GetCharacter(limit int, page int) (models.HttpResponse, error){

	client := &http.Client{
    Timeout: 15 * time.Second,
	}

	resp, err := client.Get("https://dragonball-api.com/api/characters?page=" + strconv.Itoa(page) + "&limit=" + strconv.Itoa(limit)) 
	
	if err != nil {
		return models.HttpResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.HttpResponse{}, err
	}

	var response models.HttpResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.HttpResponse{}, err
	}

	return response, nil
}