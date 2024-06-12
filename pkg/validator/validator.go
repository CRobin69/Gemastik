package validator

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func ValidatePassword(password string) bool {
	allowedSymbols := "!@#$%^&*()-_+="

	containsLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	containsUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	containsDigit := regexp.MustCompile(`\d`).MatchString(password)

	var containsSymbol bool
	for _, char := range password {
		if strings.ContainsRune(allowedSymbols, char) {
			containsSymbol = true
			break
		}
	}

	return containsLowercase && containsUppercase && containsDigit && containsSymbol && len(password) >= 8
}

func ValidatePhone(phoneNumber string) bool {
	if len(phoneNumber) > 13 || len(phoneNumber) < 12 || (phoneNumber[:1] != "0" && phoneNumber[:1] != "62") {
		return false
	}

	for _, char := range phoneNumber {
		if char < '0' || char > '9' || char == '+' {
			return false
		}
	}
	return true
}