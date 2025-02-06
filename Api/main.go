package main

import (
	"log"
	"net/http"

	"github.com/Bromolima/rpg-character-manager/handlers"
	"github.com/Bromolima/rpg-character-manager/repositories"
	"github.com/Bromolima/rpg-character-manager/services"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func main() {
	e := echo.New()

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	injector := do.New()

	do.Provide(injector, repositories.NewUserRepository)
	do.Provide(injector, services.NewUserService)
	do.Provide(injector, handlers.NewUserHandler)
}
