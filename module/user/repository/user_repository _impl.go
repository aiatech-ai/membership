package repository

import (
	"context"

	"github.com/shinystarvn/membership/module/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (s *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	return s.db.Create(user).Error
}

func (s *userRepository) FindOneUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	s.db.WithContext(ctx)
	err := s.db.First(&user, "email =? ", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userRepository) FindOneUserByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update attributes with `map`
// map[string]interface{}{"name": "hello", "age": 18, "active": false}
// UPDATE user SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10';
func (s *userRepository) UpdateUser(ctx context.Context, userID int64, arg map[string]interface{}) error {
	return s.db.Model(&model.User{Id: userID}).Updates(arg).Error
}
