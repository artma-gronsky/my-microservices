package validation

import "fmt"

// PositiveNonZeroConstraint checks for value > 0
func PositiveNonZeroConstraint(value int, fieldName string) Constraint {
	return func() error {
		if value <= 0 {
			return fmt.Errorf("%s cannot be zero or negative", fieldName)
		}
		return nil
	}
}
