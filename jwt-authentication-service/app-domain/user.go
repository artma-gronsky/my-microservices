package app_domain

import (
	"context"
	"time"
)

type User struct {
	ID          int64     `bson:"_id,omitempty"json:"id"`
	Username    string    `bson:"name"json:"username"validate:"required"`
	Password    string    `bson:"password"json:"password" validate:"required"`
	Email       string    `bson:"email"json:"email" validate:"required"`
	Description string    `bson:"description"json:"description,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at"json:"updated_at"`
	CreatedAt   time.Time `bson:"created_at"json:"created_at"`
}

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Store(context.Context, *User) error
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Store(context.Context, *User) error
}
