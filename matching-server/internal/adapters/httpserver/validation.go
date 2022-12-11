package httpserver

import (
	"errors"
	"regexp"

	"github.com/raymondgitonga/matching-server/internal/core/dormain"
)

const phoneRegex = `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`

// ValidateCustomerRequest : Checks whether the material and phone numbers provided are valid
func ValidateCustomerRequest(request dormain.CustomerRequest) error {
	materials := []string{"tiles", "wood", "carpet"}
	isValid := false

	for i := range materials {
		if materials[i] == request.Material {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("unrecognized material type")
	}

	r := regexp.MustCompile(phoneRegex)
	validPhone := r.MatchString(request.Phone)

	if !validPhone {
		return errors.New("wrong format for phone number")
	}
	return nil
}
