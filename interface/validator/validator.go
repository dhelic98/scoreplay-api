package validator

import "github.com/go-playground/validator/v10"

var validatorSingleton *validator.Validate

func GetValidatorInstance() *validator.Validate {
	if validatorSingleton == nil {
		validatorSingleton = validator.New(validator.WithRequiredStructEnabled())
	}

	return validatorSingleton
}
