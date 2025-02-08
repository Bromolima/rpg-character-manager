package main

import (
	"log"
	"net/http"

	"github.com/Bromolima/rpg-character-manager/internal"
	"github.com/Bromolima/rpg-character-manager/repositories"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	injector := internal.NewDi()
	internal.Provide(*injector, repositories.NewUserRepository)
}
