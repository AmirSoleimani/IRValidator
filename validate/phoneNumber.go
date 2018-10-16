package validate

import (
	"errors"
	"regexp"
)

//MobilePhone mobile phone number validation
func MobilePhone(value string) (string, error) {
	re := regexp.MustCompile(`^(?:0|\+98|0098)(\d{10})$`)
	result := re.FindStringSubmatch(value)
	if len(result) > 0 {
		return result[1], nil
	}
	err := errors.New("invalid phone number")
	return "", err
}
