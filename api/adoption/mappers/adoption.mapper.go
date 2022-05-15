package mappers

import (
	"github.com/cleysopnph/adote-um-pet/api/adoption/dtos"
	"github.com/cleysopnph/adote-um-pet/api/pet/mappers"
	"github.com/cleysopnph/adote-um-pet/core/models"
)

func NewAdoptionMapper(petMapper mappers.PetMapper) AdoptionMapper {
	return &adoptionMapper{petMapper: petMapper}
}

type adoptionMapper struct {
	petMapper mappers.PetMapper
}

// ToModel implements AdoptionMapper
func (*adoptionMapper) ToModel(request dtos.AdoptionRequest) models.Adoption {
	return models.Adoption{
		Email: request.Email,
		Value: request.Value,
		Pet:   models.Pet{ID: request.PetID},
	}
}

// ToCollectionResponse implements AdoptionMapper
func (m *adoptionMapper) ToCollectionResponse(models []models.Adoption) []dtos.AdoptionResponse {
	response := []dtos.AdoptionResponse{}
	for _, model := range models {
		response = append(response, m.ToResponse(model))
	}
	return response
}

// ToResponse implements AdoptionMapper
func (m *adoptionMapper) ToResponse(model models.Adoption) dtos.AdoptionResponse {
	return dtos.AdoptionResponse{
		ID:    model.ID,
		Email: model.Email,
		Value: model.Value,
		Pet:   m.petMapper.ToResponse(model.Pet),
	}
}
