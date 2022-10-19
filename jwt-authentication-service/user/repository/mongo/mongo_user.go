package mongo

import (
	"context"
	domain "github.com/artmadar/jwt-auth-service/app-domain"
)

type mongoUserRepository struct {
}

func NewMongoUserRepository() domain.UserRepository {
	return &mongoUserRepository{}
}

func (m *mongoUserRepository) Store(ctx context.Context, user *domain.User) error {
	// todo: implement repository
	return nil
}
