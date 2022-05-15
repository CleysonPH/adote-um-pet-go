package mappers

import (
	"github.com/cleysopnph/adote-um-pet/api/adoption/dtos"
	"github.com/cleysopnph/adote-um-pet/core/models"
)

type AdoptionMapper interface {
	ToResponse(model models.Adoption) dtos.AdoptionResponse
	ToCollectionResponse(models []models.Adoption) []dtos.AdoptionResponse
	ToModel(request dtos.AdoptionRequest) models.Adoption
}
