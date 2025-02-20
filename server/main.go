package main

import (
	"log/slog"

	dependecyInjection "github.com/Bromolima/rpg-character-manager/config/dependecy_injection"
	"github.com/Bromolima/rpg-character-manager/handlers"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func main() {
	logger := slog.Default()
	logger.Info("Hello world")

	e := echo.New()
	i := do.New()

	dependecyInjection.Injections(i)
	handlers.SetupRoutes(e, i)
	e.Logger.Fatal(e.Start(":1323"))
}
