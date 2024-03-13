package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/teste/scr/configuration/rest_error"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *rest_error.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return rest_error.NewBadRequestError(jsonErr.Error())
	}
	if errors.As(validationError, &jsonValidationError) {
		var errorCauses = []rest_error.Causes{}

		for _, err := range jsonValidationError {
			cause := rest_error.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_error.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	}
	return rest_error.NewBadRequestError("Error trying convert fields")
}
