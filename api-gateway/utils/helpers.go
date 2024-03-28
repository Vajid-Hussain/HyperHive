package helper_api_gateway

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validator(request any) []string {
	errResponse := []string{}

	var validate = validator.New()
	errs := validate.Struct(request)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		errResponse = append(errResponse, fmt.Sprintf("[%s] : '%v' | become '%s' %s", err.Field(), err.Value(), err.Tag(), err.Param()))
	}

	return errResponse
}
