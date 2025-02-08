package services

import (
	"context"
	"fmt"

	"github.com/Bromolima/rpg-character-manager/internal"
	"github.com/Bromolima/rpg-character-manager/models"
	"github.com/Bromolima/rpg-character-manager/repositories"
)

type UserService interface {
	Create(ctx context.Context, payload *models.UserPayload) error
}

type userService struct {
	i              internal.Di
	userRepository repositories.UserRepository
}

func NewUserService(i internal.Di) (UserService, error) {
	userRepository, err := internal.Invoke[repositories.UserRepository](i)

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
