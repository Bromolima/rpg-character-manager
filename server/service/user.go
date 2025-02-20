package service

import (
	"context"
	"log/slog"

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
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "CreateUser"),
	)

	log.Info("starting creating user")

	userExist, err := s.userRepository.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		log.Error("failed to get user", "error", err)
		return err
	}

	if userExist != nil {
		log.Warn("email already exists", "error", err)
		return domain.ErrEmailAlreadyExists
	}

	if err := s.userRepository.CreateUser(ctx, *payload.ToUser()); err != nil {
		log.Error("Failed to create user", "error", err)
		return err
	}

	log.Info("User created sucessfully")

	return nil
}
