package validators

import "github.com/go-playground/validator/v10"

type AdoptionValidator interface {
	PetExistsById(fl validator.FieldLevel) bool
}
