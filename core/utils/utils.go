package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetValidationErrors(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		fmt.Println(f)
		errs[f.Field()] = f.ActualTag() + " " + f.Param()
	}

	return errs
}
