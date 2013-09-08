package utils

import (
	"errors"
	"regexp"
)

func NormalizePhoneToE164(phone string) (string, error) {
	notNumber := regexp.MustCompile("[^0-9]+")
	number := notNumber.ReplaceAllLiteralString(phone, "")
	if len(number) < 10 {
		return "", errors.New("Phone number was too short! Include area code, and if outside the US, country code")
	}

	if len(number) == 10 { // assume 10 digit entries are US phone numbers
		number = "+1" + number
	} else if len(number) <= 15 { // international numbers max out at 15 digits
		number = "+" + number
	} else {
		return "", errors.New("Provided phone number had too many digits")
	}

	return number, nil
}
