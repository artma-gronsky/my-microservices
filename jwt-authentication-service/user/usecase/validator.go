package usecase

import (
	"github.com/artmadar/jwt-auth-service/app-domain/entities"
	"github.com/artmadar/jwt-auth-service/pkg/validation"
)

func Validate(user *entities.User) (err error) {
	validator := validation.Validator{}
	validator.AddConstraint(validation.NotEmptyConstraint(user.Username, "username"))
	validator.AddConstraint(validation.NotEmptyConstraint(user.Email, "email")).AddConstraint(validation.EmailConstraint(user.Email, "email"))

	err = validator.Validate()
	return
}
