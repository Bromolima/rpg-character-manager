package services

import (
	"context"
	"fmt"

	"github.com/Bromolima/rpg-character-manager/models"
	"github.com/Bromolima/rpg-character-manager/repositories"
	"github.com/samber/do/v2"
)

type UserService interface {
	Create(ctx context.Context, payload *models.UserPayload) error
}

type userService struct {
	i              do.Injector
	userRepository repositories.UserRepository
}

func NewUserService(i do.Injector) (UserService, error) {
	userRepository, err := do.Invoke[repositories.UserRepository](i)

	if err != nil {
		return nil, err
	}
	return &userService{
		userRepository: userRepository,
	}, err
}

func (s *userService) Create(ctx context.Context, payload *models.UserPayload) error {
	userWithSameEmail, err := s.userRepository.GetByEmail(ctx, payload.Email)
	if err != nil {
		return err
	}

	if userWithSameEmail != nil {
		return fmt.Errorf("email already used")
	}

	if err := s.userRepository.Create(ctx, payload.PayloadToUser()); err != nil {
		return err
	}
	return nil
}
