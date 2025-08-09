package models


type Character struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Ki          string      `json:"ki"`
	MaxKi       string      `json:"maxKi"`
	Race        string      `json:"race"`
	Gender      string      `json:"gender"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Affiliation string      `json:"affiliation"`
	DeletedAt   interface{} `json:"deletedAt"`
}


type MetaData struct {
	TotalItems int `json:"totalItems"`
	ItemCount int `json:"itemCount"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalPages int `json:"totalPages"`
	CurrentPage int `json:"currentPage"`
}

type Link struct {
	First string `json:"first"`
	Previous string `json:"previous"`
	Next string `json:"next"`
	Last string `json:"last"`
}

type HttpResponse struct{
	Items []Character `json:"items"`
	MetaData MetaData `json:"metaData"`
	Links Link `json:"links"`
}

type Response struct {
	Items []Character `json:"items"`
	Links Link `json:"links"`
}