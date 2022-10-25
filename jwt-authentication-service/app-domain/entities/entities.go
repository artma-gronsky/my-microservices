package entities

import "time"

type User struct {
	ID          string    `bson:"_id,omitempty"json:"id"`
	Username    string    `bson:"username"json:"username"validate:"required"`
	Password    string    `bson:"password"json:"password" validate:"required"`
	Email       string    `bson:"email"json:"email" validate:"required"`
	Description string    `bson:"description"json:"description,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at"json:"updated_at"`
	CreatedAt   time.Time `bson:"created_at"json:"created_at"`
}
