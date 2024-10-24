package directives

import (
	"context"
	"errors"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()

	validate.RegisterValidation("containsNumber", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[0-9]`, fl.Field().String())
		return match
	})

	validate.RegisterValidation("containsSpecialCharacter", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[!@#$%^&*()]`, fl.Field().String())
		return match
	})
}

func ValidateRequest(ctx context.Context, obj interface{}, next graphql.Resolver, rule string) (interface{}, error) {
	val, err := next(ctx)

	if err != nil {
		return nil, err
	}

	fieldName := *graphql.GetPathContext(ctx).Field
	errorMessages := ""

	validationErr := validate.Var(val, rule)
	if validationErr != nil {
		for _, err := range validationErr.(validator.ValidationErrors) {
			errorMessages += getErrorMessage(fieldName, err)
		}
		return nil, errors.New(errorMessages)
	}

	return next(ctx)
}

func getErrorMessage(fieldName string, err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fieldName + " is required"
	case "email":
		return "the email is invalid"
	case "min":
		return "the minimum length of " + fieldName + " is equals " + err.Param()
	case "max":
		return "the maximum length of " + fieldName + " is equals " + err.Param()
	case "containsNumber":
		return "the " + fieldName + " must contains number"
	case "containsSpecialCharacter":
		return "the " + fieldName + " must contains special character"
	case "gte":
		return "the " + fieldName + " must be greater than or equal to 1"
	default:
		return "validation error in " + fieldName
	}
}
