package mid

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/configs"
)

func LimitCheckMiddleware(conf *configs.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			letGo := false
			cookie, err := ctx.Cookie("login_guard")
			value := cookie.Value
			authorizationKey := ctx.Request().Header.Get("authorization")
			if err != nil && len(authorizationKey) < 10 {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			if authorizationKey == "" && value != "" {
				authorizationKey = value
				letGo = true
			} else {

			}
			if !letGo {
				ctx.NoContent(http.StatusUnauthorized)
			}
			return next(ctx)
		}
	}

}
