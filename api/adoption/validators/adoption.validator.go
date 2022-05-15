package validators

import (
	"github.com/cleysopnph/adote-um-pet/core/repositories"
	"github.com/go-playground/validator/v10"
)

func NewAdoptionValidator(petRepository repositories.PetRepository) AdoptionValidator {
	return &adoptionValidator{petRepository: petRepository}
}

type adoptionValidator struct {
	petRepository repositories.PetRepository
}

// PetExistsById implements AdoptionValidator
func (v *adoptionValidator) PetExistsById(fl validator.FieldLevel) bool {
	id, ok := fl.Field().Interface().(uint)
	if ok {
		return v.petRepository.ExistsById(id)
	}
	return true
}
