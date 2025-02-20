package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func SetupRoutes(e *echo.Echo, i do.Injector) {
	userHanlder, err := do.Invoke[UserHanlder](i)
	if err != nil {
		panic(err)
	}

	e.POST("/", userHanlder.CreateUser)
}
