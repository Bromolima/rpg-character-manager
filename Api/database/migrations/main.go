package main

import (
	"log"

	"github.com/Bromolima/rpg-character-manager/config"
	"github.com/Bromolima/rpg-character-manager/database"
	"github.com/Bromolima/rpg-character-manager/models"
)

func main() {
	config.LoadEnvironments()

	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Character{},
	); err != nil {
		log.Fatal(err)
	}

	log.Println("Migration done")
}
