package httpserver

import (
	"errors"
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
	"regexp"
)

const phoneRegex = `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`

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
