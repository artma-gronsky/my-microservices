package validation

import (
	"fmt"
	"strings"
)

// NotEmptyConstraint checks for non-empty string
func NotEmptyConstraint(value string, fieldName string) Constraint {
	return func() error {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s cannot be empty", fieldName)
		}
		return nil
	}
}
