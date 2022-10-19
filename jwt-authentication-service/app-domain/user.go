package app_domain

import (
	"context"
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Owner       string    `json:"owner" validate:"required"`
	Description string    `json:"description,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Store(context.Context, *User) error
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Store(context.Context, *User) error
}
