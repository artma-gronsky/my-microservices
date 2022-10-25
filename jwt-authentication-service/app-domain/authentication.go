package app_domain

import (
	"github.com/artmadar/jwt-auth-service/app-domain/entities"
	"golang.org/x/net/context"
)

type Authentication struct {
	Token string
}

type AuthenticationUsecase interface {
	GenerateToken(ctx context.Context, username, password string) (*Authentication, error)
}

type UserForAuthenticationRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
}
