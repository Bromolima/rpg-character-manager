package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID   `gorm:"column:ID"`
	Email      string      `gorm:"column:email"`
	Username   string      `gorm:"column:username"`
	Password   string      `gorm:"column:password"`
	Characters []Character `gorm:"column:characters"`
	CreatedAt  time.Time   `gorm:"column:CreatedAt"`
	UpdatedAt  time.Time   `gorm:"column:UpdatedAt"`
}

func (u *User) TableName() string {
	return "user"
}

type UserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (up *UserPayload) UserPayloadToUser() *User {
	return &User{
		ID:        uuid.New(),
		Email:     up.Email,
		Username:  up.Username,
		Password:  up.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
