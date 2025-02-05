package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID         string      `gorm:"column:ID;type:varchar(255);not null;unique"`
	Email      string      `gorm:"column:email;type:varchar(255);not null"`
	Username   string      `gorm:"column:username;type:varchar(255);not null"`
	Password   string      `gorm:"column:password;type:varchar(255);not null"`
	Characters []Character `gorm:"column:characters;foreignKey:CharacterID"`
	CreatedAt  time.Time   `gorm:"column:CreatedAt;type:date"`
	UpdatedAt  time.Time   `gorm:"column:UpdatedAttype:date"`
}

func (u *User) TableName() string {
	return "user"
}

type UserPayload struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (p *UserPayload) PayloadToUser() *User {
	return &User{
		ID:        uuid.New().String(),
		Email:     p.Email,
		Username:  p.Username,
		Password:  p.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *UserPayload) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}
