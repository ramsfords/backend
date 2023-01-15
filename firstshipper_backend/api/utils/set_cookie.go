package utils

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

func SetCookie(ctx echo.Context, key string, value string, maxage int, services services.Services) {
	cookies := &http.Cookie{
		Name:   key,
		Value:  value,
		Path:   "/",
		Domain: GetOrigin(ctx, services),
		MaxAge: maxage,
	}
	ctx.SetCookie(cookies)
}
