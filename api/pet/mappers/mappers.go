package mappers

import (
	"github.com/cleysopnph/adote-um-pet/api/pet/dtos"
	"github.com/cleysopnph/adote-um-pet/core/models"
)

type PetMapper interface {
	ToModel(request dtos.PetRequest) models.Pet
	ToResponse(model models.Pet) dtos.PetResponse
	ToCollectionResponse(models []models.Pet) []dtos.PetResponse
}
