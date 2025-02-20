package handlers

import (
	"net/http"

	apiErrors "github.com/Bromolima/rpg-character-manager/config/api_errors"
	"github.com/Bromolima/rpg-character-manager/config/validation"
	"github.com/Bromolima/rpg-character-manager/domain"
	"github.com/Bromolima/rpg-character-manager/service"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

type UserHanlder interface {
	CreateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}

type userHanlder struct {
	i           do.Injector
	userService service.UserService
}

func NewUserHandler(i do.Injector) (UserHanlder, error) {
	userService, err := do.Invoke[service.UserService](i)
	if err != nil {
		return nil, err
	}

	return &userHanlder{
		i:           i,
		userService: userService,
	}, nil
}

func (h *userHanlder) CreateUser(ctx echo.Context) error {
	var payload domain.UserPayload
	if err := ctx.Bind(&payload); err != nil {
		apiErr := apiErrors.NewBadRequestErr("failed to bind payload")
		return ctx.JSON(apiErr.Code, apiErr)
	}

	if err := validation.Validate.Struct(payload); err != nil {
		apiErr := validation.ValidateUserError(err)
		return ctx.JSON(apiErr.Code, apiErr)
	}

	return ctx.String(http.StatusOK, "user created successfuly")
}

func (h *userHanlder) DeleteUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "user deleted sucessfuly")
}
