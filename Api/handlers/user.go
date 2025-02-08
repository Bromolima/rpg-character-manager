package handlers

import (
	"fmt"

	"github.com/Bromolima/rpg-character-manager/internal"
	"github.com/Bromolima/rpg-character-manager/models"
	"github.com/Bromolima/rpg-character-manager/services"
	"github.com/labstack/echo/v4"
)

type UserHanlder interface {
	Create(ctx echo.Context) error
}

type userHanlder struct {
	i           internal.Di
	userService services.UserService
}

func NewUserHandler(i internal.Di) (UserHanlder, error) {
	userService, err := internal.Invoke[services.UserService](i)

	if err != nil {
		return nil, err
	}

	return &userHanlder{
		userService: userService,
	}, nil
}

func (h *userHanlder) Create(ctx echo.Context) error {
	var payload *models.UserPayload

	if err := ctx.Bind(&payload); err != nil {
		return fmt.Errorf("failed to process request")
	}

	if err := payload.Validate(); err != nil {
		return fmt.Errorf("failed validating payload")
	}

	if err := h.userService.Create(ctx.Request().Context(), payload); err != nil {
		return err
	}

	return nil
}
