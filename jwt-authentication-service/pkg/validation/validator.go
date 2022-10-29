package validation

import (
	"fmt"
	"strings"
)

// Constraint is type alias for func() error
type Constraint func() error

// Validator ...
type Validator struct {
	constraints []Constraint
}

// Validate does the validation
func (validator *Validator) Validate() error {
	errors := make([]string, 0)
	for _, constraint := range validator.constraints {
		err := constraint()
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(errors, ", "))
}

func (v *Validator) AddConstraint(constraint Constraint) *Validator {
	v.constraints = append(v.constraints, constraint)
	return v
}
