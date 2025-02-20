package service

import (
	"context"
	"fmt"

	"github.com/Bromolima/rpg-character-manager/domain"
	"github.com/Bromolima/rpg-character-manager/repository"
	"github.com/samber/do/v2"
)

type UserService interface {
	CreateUser(ctx context.Context, payload domain.UserPayload) error
}

type userService struct {
	i              do.Injector
	userRepository repository.UserRepository
}

func NewUserService(i do.Injector) (UserService, error) {
	userRepository, err := do.Invoke[repository.UserRepository](i)
	if err != nil {
		return nil, err
	}

	return &userService{
		i:              i,
		userRepository: userRepository,
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, payload domain.UserPayload) error {
	userExist, err := s.userRepository.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		return fmt.Errorf("failed to get user")
	}

	if userExist != nil {
		return fmt.Errorf("user already exists")
	}

	if err := s.userRepository.CreateUser(ctx, *payload.ToUser()); err != nil {
		return err
	}

	return nil
}
