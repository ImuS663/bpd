package net

import "testing"

const urlForValidate = "https://example.com"
const urlForValidateError = "example.com"

func TestValidateURLValid(t *testing.T) {
	result := ValidateURL(urlForValidate)

	if !result {
		t.Error("expected true, got false")
	}
}

func TestValidateURLInvalid(t *testing.T) {
	result := ValidateURL(urlForValidateError)

	if result {
		t.Error("expected false, got true")
	}
}
