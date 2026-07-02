package models

import "slices"

type ValidationError struct {
	Position int    `json:"position,omitempty"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}

func (error *ValidationError) New(condition bool, field string, message string, position *int) *ValidationError {
	if condition {
		error.Field = field
		error.Message = message

		if position != nil {
			error.Position = *position
		}

		return error
	}

	return nil
}

func removeNullValue(slice []*ValidationError) []*ValidationError {
	var output []*ValidationError

	for _, element := range slice {
		if element != nil {
			output = append(output, element)
		}
	}

	return output
}

func MakeValidationErrors(errors ...*ValidationError) []*ValidationError {
	return removeNullValue(slices.Clone(errors))
}
