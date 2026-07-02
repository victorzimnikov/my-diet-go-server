package models

import (
	"testing"
)

func TestMakeValidationErrors(t *testing.T) {
	position := 1

	result := MakeValidationErrors(
		new(ValidationError).New(true, "test", "Test message", nil),
		new(ValidationError).New(false, "test2", "Test message", nil),
		new(ValidationError).New(true, "test3", "Test message", &position),
	)

	if result[0].Field != "test" {
		t.Error("Error")
	}

	if result[1].Position != 1 {
		t.Error("Error")
	}
}
