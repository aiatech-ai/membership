package repository

import (
	"context"

	"github.com/shinystarvn/membership/module/model"
)

type (
	UpdatePassword struct {
		PasswordHash string `json:"password_hash"`
	}
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindOneUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindOneUserByID(ctx context.Context, id int64) (*model.User, error)
	// Update attributes with `map`
	// map[string]interface{}{"name": "hello", "age": 18, "active": false}
	// UPDATE user SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10';
	UpdateUser(ctx context.Context, userID int64, arg map[string]interface{}) error
}
