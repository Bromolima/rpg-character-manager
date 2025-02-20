package repository

import (
	"context"

	"github.com/Bromolima/rpg-character-manager/domain"
	"github.com/samber/do/v2"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
}

type userRepository struct {
	i do.Injector
}

func NewUserRepository(i do.Injector) (UserRepository, error) {
	return &userRepository{
		i: i,
	}, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user domain.User) error {
	return nil
}
