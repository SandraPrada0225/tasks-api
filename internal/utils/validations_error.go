package utils

type FieldError struct {
	Field   string `json: "field"`
	Message string `json: "message"`
}

//lista de errores por campo

type ValidationError struct {
	Errors []FieldError `json: "errors"`
}

func (v *ValidationError) Error() string {
	return "validation error"
}

func (v *ValidationError) Add(field, message string) {
	v.Errors = append(v.Errors, FieldError{
		Field:   field,
		Message: message,
	})
}

func (v *ValidationError) HasErrors() bool {
	return len(v.Errors) > 0
}
