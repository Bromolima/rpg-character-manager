package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmailAlreadyExists = errors.New("email is already in database")
)

type User struct {
	BaseModel
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Username string `gorm:"not null"`
	ImageUrl string `gorm:""`
}

type UserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,containsany=@#&%*"`
	Username string `json:"username" validate:"required,min=4,max=20"`
	ImageUrl string `json:"imageUrl"`
}

type UserResponse struct {
	Username string `json:"username"`
}

func (p *UserPayload) ToUser() *User {
	return &User{
		BaseModel: BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		Email:    p.Email,
		Password: p.Password,
		Username: p.Username,
		ImageUrl: p.ImageUrl,
	}
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		Username: u.Username,
	}
}
