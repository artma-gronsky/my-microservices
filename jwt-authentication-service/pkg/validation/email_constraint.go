package validation

import (
	"fmt"
	"net/mail"
)

func EmailConstraint(value, fieldName string) Constraint {
	return func() error {
		if _, err := mail.ParseAddress(value); err != nil {
			return fmt.Errorf("%s is not email", fieldName)
		}

		return nil
	}
}
