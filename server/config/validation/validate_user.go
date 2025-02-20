package validation

import (
	"encoding/json"
	"errors"

	apiErrors "github.com/Bromolima/rpg-character-manager/config/api_errors"
	"github.com/go-playground/validator/v10"
)

var (
	Validate = validator.New()
)

func ValidateUserError(validationErr error) *apiErrors.ApiError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return apiErrors.NewBadRequestErr("invalid field type")
	}

	if errors.As(validationErr, &jsonValidationError) {
		errorsCauses := []apiErrors.Causes{}

		for _, err := range validationErr.(validator.ValidationErrors) {
			cause := apiErrors.Causes{
				Message: err.Error(),
				Field:   err.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return apiErrors.NewBadRequestValidationErr("some fields are invalid", errorsCauses)
	}

	return apiErrors.NewBadRequestErr("error trying to convert fields")
}
