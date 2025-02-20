package repository

import (
	"context"

	"github.com/Bromolima/rpg-character-manager/domain"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userRepository struct {
	i  do.Injector
	db *gorm.DB
}

func NewUserRepository(i do.Injector) (UserRepository, error) {
	db, err := do.Invoke[*gorm.DB](i)
	if err != nil {
		return nil, err
	}

	return &userRepository{
		i:  i,
		db: db,
	}, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user domain.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
