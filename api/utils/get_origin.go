package utils

import (
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/services"
)

func GetOrigin(ctx echo.Context, services services.Services) string {
	origin := ctx.Request().Header.Get("Origin")
	valid := strings.Contains(origin, "firstshipper.com") || strings.Contains(origin, "127.0.0.1")
	if !valid {
		origin = ""
	}
	if services.Conf.Env != "dev" {
		if strings.Contains(origin, "www") {
			origin = strings.Split(origin, "www.")[1]
		} else {
			origin = strings.Split(origin, "//")[1]
		}

	} else {
		origin = "127.0.0.1"

	}
	return origin
}
