package service

import (
	"context"

	"github.com/shinystarvn/membership/module/model"
)

type (
	UpdateUserProfileInput struct {
		FirstName string
		LastName  string
		Email     string
	}

	UpdatePasswordInput struct {
		CurrentPassword string
		NewPassword     string
	}
)
type UserService interface {
	GetUserProfile(ctx context.Context, email string) (*model.User, error)
	UpdateUserProfile(ctx context.Context, userID int64, input UpdateUserProfileInput) error
	UpdatePassword(ctx context.Context, userID int64, input UpdatePasswordInput) error
}
