package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct(obj interface{}) []ValidationError {

	errs := validate.Struct(obj)
	errors := []ValidationError{}

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ValidationError

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value

			errors = append(errors, elem)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
