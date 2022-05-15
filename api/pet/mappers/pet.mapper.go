package mappers

import (
	"github.com/cleysopnph/adote-um-pet/api/pet/dtos"
	"github.com/cleysopnph/adote-um-pet/core/models"
)

func NewPetMapper() PetMapper {
	return &petMapper{}
}

type petMapper struct{}

// ToModel implements PetMapper
func (m *petMapper) ToModel(request dtos.PetRequest) models.Pet {
	return models.Pet{
		Name:    request.Name,
		History: request.History,
		Picture: request.Picture,
	}
}

// ToCollectionResponse implements PetMapper
func (m *petMapper) ToCollectionResponse(models []models.Pet) []dtos.PetResponse {
	response := []dtos.PetResponse{}
	for _, model := range models {
		response = append(response, m.ToResponse(model))
	}
	return response
}

// ToResponse implements PetMapper
func (m *petMapper) ToResponse(model models.Pet) dtos.PetResponse {
	return dtos.PetResponse{
		ID:      model.ID,
		Name:    model.Name,
		History: model.History,
		Picture: model.Picture,
	}
}
