package utils

import (
	"github.com/labstack/echo/v5"
	"github.com/pkg/errors"
	"github.com/ramsfords/backend/business/core/model"
)

func GetAuthContext(ctx echo.Context) (*model.Session, error) {
	authContext := ctx.Get("authContext")
	if authContext == nil {
		return &model.Session{}, errors.New("authContext not found")
	}
	return authContext.(*model.Session), nil
}
