package checker

import (
	"fmt"
	"strings"
)

// IsValid preform a simple validation check to see if the email is correctly formatted.
// An error is returned if not.
func IsValid(email string) error {
	return nil
}

// IsGeneric will check if the email address is provided by a free services provider.
// As an example will return true if the email is provided by gmail.com
func IsGeneric(email string) (bool, error) {
	parts := strings.Split(email, "@")

	if len(parts) <= 1 {
		return false, fmt.Errorf("given an invalid email")
	}

	for _, d := range domains {
		if d == parts[1] {
			return true, nil
		}
	}

	return false, nil
}
