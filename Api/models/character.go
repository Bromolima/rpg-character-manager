package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Character struct {
	ID          string    `gorm:"collunm:ID;type:varchar(255)"`
	Name        string    `gorm:"collunm:name;type:varchar(255)"`
	ImageUrl    string    `gorm:"colunm:imageUrl;varchar(255)"`
	Description string    `gorm:"colunm:description;type:text"`
	CreatedAt   time.Time `gorm:"column:CreatedAt;type:date"`
	UpdatedAt   time.Time `gorm:"column:UpdatedAttype:date"`
}

func (c *Character) TableName() string {
	return "character"
}

type CharacterPayload struct {
	Name        string `json:"name" validate:"required"`
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description" validate:"required"`
}

type CharacterResponse struct {
	Name        string `json:"name"`
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description"`
}

func (p *CharacterPayload) PayloadToCharacter() *Character {
	return &Character{
		ID:          uuid.New().String(),
		Name:        p.Name,
		ImageUrl:    p.ImageUrl,
		Description: p.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (p *CharacterPayload) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}
