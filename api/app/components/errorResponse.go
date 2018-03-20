package components

import "github.com/revel/revel"

type ValidationErrorResponse struct {
	 ErrorType string  `json:"errorType"`
	 Errors []*revel.ValidationError `json:"errors"`
}

func NewValidationErrorResponse(errors []*revel.ValidationError) (ValidationErrorResponse) {
	return ValidationErrorResponse{ "Validation",errors }
}

func (e ValidationErrorResponse) AddErrors(errors []*revel.ValidationError) {
	for _, error := range errors {
		e.Errors = append(e.Errors, error)
	}
}