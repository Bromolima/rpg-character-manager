package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"type:date"`
	UpdatedAt time.Time `gorm:"type:date"`
	DeletedAt time.Time `gorm:"type:date"`
}
