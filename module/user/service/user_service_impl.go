package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fatih/structs"

	"github.com/shinystarvn/membership/internal/errs"
	"github.com/shinystarvn/membership/internal/security/crypto"
	"github.com/shinystarvn/membership/internal/util"
	"github.com/shinystarvn/membership/module/model"
	"github.com/shinystarvn/membership/module/user/repository"
)

type userService struct {
	repo            repository.UserRepository
	passwordEncoder crypto.PasswordEncoder
}

func NewUserService(repository repository.UserRepository, passwordEncoder crypto.PasswordEncoder) UserService {
	return &userService{
		repo:            repository,
		passwordEncoder: passwordEncoder,
	}
}

func (s *userService) GetUserProfile(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.FindOneUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewPublicError(errs.ErrUserNotFound)
		}
		return nil, errs.NewPrivateError("Can not get user profile", err)
	}
	return user, nil
}

func (s *userService) UpdateUserProfile(ctx context.Context, userID int64, input UpdateUserProfileInput) error {

	return s.repo.UpdateUser(ctx, userID, structs.Map(input))
}

func (s *userService) UpdatePassword(ctx context.Context, userID int64, input UpdatePasswordInput) error {

	user, err := s.repo.FindOneUserByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewPrivateError("User not found", err)
		}
		return err
	}

	if !s.passwordEncoder.Matches(input.CurrentPassword, user.PasswordHash) {
		return errors.New(errs.ErrInvalidCredential)
	}

	encodedNewPassword, err := s.passwordEncoder.Encode(input.NewPassword)
	if err != nil {
		return err
	}

	params := repository.UpdatePassword{
		PasswordHash: encodedNewPassword,
	}
	p, err := util.StructToJsonMap(params)
	if err != nil {
		return errors.New(errs.UnablePrepareQuery)
	}
	return s.repo.UpdateUser(ctx, userID, p)
}
