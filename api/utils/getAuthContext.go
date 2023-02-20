package utils

import (
	"github.com/labstack/echo/v5"
	"github.com/pkg/errors"
	"github.com/ramsfords/backend/api/authapi"
)

func GetAuthContext(ctx echo.Context) (authapi.LoginData, error) {
	authContext := ctx.Get("authContext")
	if authContext == nil {
		return authapi.LoginData{}, errors.New("authContext not found")
	}
	return authContext.(authapi.LoginData), nil
}
