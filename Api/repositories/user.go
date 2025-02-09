package repositories

import (
	"context"

	"github.com/Bromolima/rpg-character-manager/internal"
	"github.com/Bromolima/rpg-character-manager/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user models.User) error
	DeleteByID(ctx context.Context, id string) error
}

type userRepository struct {
	i  internal.Di
	db *gorm.DB
}

func NewUserRepository(i internal.Di) (UserRepository, error) {
	db, err := internal.Invoke[*gorm.DB](i)

	if err != nil {
		return nil, err
	}
	return &userRepository{
		db: db,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(ctx context.Context, user models.User) error {
	if err := r.db.WithContext(ctx).Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteByID(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
