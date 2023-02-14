package mid

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/configs"
)

func LimitForgotPassword(conf *configs.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			forgotPReq := struct {
				Email string `json:"email"`
			}{}
			err := ctx.Bind(&forgotPReq)
			if err != nil {
				return ctx.NoContent(http.StatusBadRequest)
			}
			ctx.Set("email", forgotPReq.Email)
			return next(ctx)
		}
	}

}
