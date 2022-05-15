package config

import (
	"reflect"
	"strings"

	"github.com/cleysopnph/adote-um-pet/api/adoption/validators"
	"github.com/cleysopnph/adote-um-pet/core/repositories"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ConfigValidator(petRepository repositories.PetRepository) {
	registerTagNameFunc()

	adoptionValidator := validators.NewAdoptionValidator(petRepository)
	registerAdoptionValidator(adoptionValidator)
}

func registerTagNameFunc() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func registerAdoptionValidator(adoptionValidator validators.AdoptionValidator) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("petexists", adoptionValidator.PetExistsById)
	}
}
