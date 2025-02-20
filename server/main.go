package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	dependecyInjection "github.com/Bromolima/rpg-character-manager/config/dependecy_injection"
	"github.com/Bromolima/rpg-character-manager/database"
	"github.com/Bromolima/rpg-character-manager/handlers"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func main() {
	logger := slog.Default()
	logger.Info("Hello world")

	e := echo.New()
	i := do.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.NewMysqlConnection(ctx)
	if err != nil {
		log.Fatal("error to connect to database: ", err)
	}

	do.Provide(i, func(i do.Injector) (*gorm.DB, error) {
		return db, nil
	})

	dependecyInjection.Injections(i)
	handlers.SetupRoutes(e, i)
	e.Logger.Fatal(e.Start(":1323"))
}
