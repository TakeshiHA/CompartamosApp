package functions

import (
	"net/mail"
	"regexp"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateDNI(dni string) bool {
	return digitCheck.MatchString(dni)
}

func ValidatePhoneNumber(phoneNumber string) bool {
	return digitCheck.MatchString(phoneNumber)
}
