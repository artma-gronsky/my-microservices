package app_domain

import (
	"context"
	"github.com/artmadar/jwt-auth-service/app-domain/entities"
)

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Store(context.Context, *entities.User) error
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Store(context.Context, *entities.User) error
	GetOneByOneOfFieldsValues(ctx context.Context, fieldsValues map[string]string) (*entities.User, error)
}
