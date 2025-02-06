package database

import (
	"github.com/Bromolima/rpg-character-manager/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Envi.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	postgresDB, err := db.DB()

	if err := postgresDB.Ping(); err != nil {
		_ = postgresDB.Close()
		return nil, err
	}

	return db, err
}
