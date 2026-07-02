package utils

import "testing"

func TestIsJSON(t *testing.T) {
	result := IsJSON("{\"test\": true}")

	if !result {
		t.Error("Error")
	}

	result = IsJSON("test")

	if result {
		t.Error("Error")
	}
}
