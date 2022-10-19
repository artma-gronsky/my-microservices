package usecase

import (
	"context"
	domain "github.com/artmadar/jwt-auth-service/app-domain"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(a domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       a,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Store(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err = u.userRepo.Store(ctx, user)
	return
}
