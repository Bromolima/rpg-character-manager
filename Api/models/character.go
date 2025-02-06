package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Character struct {
	BaseModel
	Name        string `gorm:"type:text"`
	ImageUrl    string `gorm:"type:text"`
	Description string `gorm:"type:text"`
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
		BaseModel: BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Name:        p.Name,
		ImageUrl:    p.ImageUrl,
		Description: p.Description,
	}
}

func (c *Character) CharacterToResponse() *CharacterResponse {
	return &CharacterResponse{
		Name:        c.Name,
		ImageUrl:    c.ImageUrl,
		Description: c.Description,
	}
}

func (p *CharacterPayload) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}
