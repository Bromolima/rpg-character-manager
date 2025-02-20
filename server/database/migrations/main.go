package main

import (
	"context"
	"log"
	"time"

	"github.com/Bromolima/rpg-character-manager/database"
	"github.com/Bromolima/rpg-character-manager/domain"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.NewMysqlConnection(ctx)
	if err != nil {
		log.Fatal("error to connect to myslq: ", err)
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("error to migrate: ", err)
	}

	log.Println("Migration excecuted sucessfully")
}
