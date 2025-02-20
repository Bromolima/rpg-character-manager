package repository

import (
	"context"
	"errors"
	"log/slog"

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
	log := slog.With(
		slog.String("repository", "user"),
		slog.String("func", "CreateUser"),
	)

	log.Info("Starting Creating user")

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		log.Error("Error to create user")
		return err
	}

	log.Info("User created")
	return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	log := slog.With(
		slog.String("repository", "user"),
		slog.String("func", "GetUserByEmail"),
	)

	log.Info("Starting getting user by email")

	var user *domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("User not found")
			return nil, nil
		}
		log.Error("Failed to get user")
		return nil, err
	}

	log.Info("User has been got by email")
	return user, nil
}
