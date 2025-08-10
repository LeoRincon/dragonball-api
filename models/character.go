package models

type Transformations struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	Ki          string      `json:"ki"`
	DeletedAt   interface{} `json:"deletedAt"`
}

type OriginPlanet struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	IsDestroyed bool        `json:"isDestroyed"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	DeletedAt   interface{} `json:"deletedAt"`
}

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

type CharacterById struct {
	Character `json:",inline"`
	OriginPlanet    OriginPlanet      `json:"originPlanet"`
	Transformations []Transformations `json:"transformations"`
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
	MetaData MetaData `json:"meta"`
	Links Link `json:"links"`
}

type Response struct {
	Items []Character `json:"items"`
	Links Link `json:"links"`
}