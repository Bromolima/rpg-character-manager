package models

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	BaseModel
	Email      string      `gorm:"not null;unique;not null"`
	Username   string      `gorm:"not null;unique;not null"`
	Password   string      `gorm:"not null"`
	Characters []Character `gorm:"characters;foreignKey:CharacterID"`
}

type UserPayload struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required,min=6,containsany=!@$%*"`
	Username string `json:"username" validate:"required"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (p *UserPayload) PayloadToUser() *User {
	return &User{
		BaseModel: BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Email:    p.Email,
		Username: p.Username,
		Password: p.Password,
	}
}

func (u *User) UserToResponse() *UserResponse {
	return &UserResponse{
		Email:    u.Email,
		Username: u.Username,
	}
}

func (p *UserPayload) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}

func (p *UserPayload) Trim() {
	p.Email = strings.TrimSpace(p.Email)
	p.Username = strings.TrimSpace(p.Username)
	p.Password = strings.TrimSpace(p.Password)
}
