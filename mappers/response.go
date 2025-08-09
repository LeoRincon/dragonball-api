package mappers

import "dragonball-api/models"

func MapResponse(response models.HttpResponse) models.Response {
	
	return models.Response{
		Items: response.Items,
		Links: response.Links,
	}
}