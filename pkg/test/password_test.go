package test

import (
	"testing"

	"github.com/CRobinDev/Gemastik/pkg/validator"
)

func TestValidatePassword(t *testing.T) {
	// Test cases for valid passwords
	validPasswords := []string{
		"ValidPassword123!",
		"AnotherValidPassword!456",
		"TestPass123@",
	}

	for _, password := range validPasswords {
		if !validator.ValidatePassword(password) {
			t.Errorf("Expected password '%s' to be valid, but it was invalid", password)
		}
	}

	// Test cases for invalid passwords
	invalidPasswords := []string{
		"shortIng1@",           // Too short
		"no-uppercase-123",     // No uppercase letter
		"NoSpecialSymbol123",   // No special symbol
		"nouppercasesymbol123", // No uppercase or special symbol
	}

	for _, password := range invalidPasswords {
		if !validator.ValidatePassword(password) {
			t.Errorf("Expected password '%s' to be invalid, but it was valid", password)
		}
	}
}
