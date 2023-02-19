package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (auth Auth) EchoSignUp(ctx echo.Context) error {
	data := &v1.User{}
	if err := ctx.Bind(data); err != nil || data.Email == "" || data.Password == "" || data.Password != data.ConfirmPassword {
		return ctx.NoContent(http.StatusBadRequest)
	}
	data.UserName = strings.ToLower(data.Email)
	data.Email = strings.ToLower(data.Email)
	data.Name = strings.ToLower(data.Name)
	userId, err := auth.services.CognitoClient.CreateUser(ctx.Request().Context(), data)
	if err != nil {
		if err == errs.ErrUserAlreadyExits {
			return ctx.NoContent(http.StatusConflict)
		}
	}

}
